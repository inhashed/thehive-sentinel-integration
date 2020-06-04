#!/usr/bin/python3
# encoding: utf-8
import json
from datetime import datetime
import urllib.request
from urllib.parse import urlencode
from datetime import datetime
import time
import ipaddress
import adal
import logging
import requests
import sys
import json
import time
import uuid
from thehive4py.api import TheHiveApi
from thehive4py.models import Alert, AlertArtifact, CustomFieldHelper
import configparser
import os
import dateutil.parser


def getAADToken(clientId, clientSecret, tenantId):
    resourceUri = 'https://westeurope.api.loganalytics.io'
    body = {
        'client_id': clientId,
        'client_secret': clientSecret,
        'grant_type': 'client_credentials',
        'resource': resourceUri
    }

    AAD_TOKEN_URL = "https://login.microsoftonline.com/%s/oauth2/token" % (tenantId)
    r = requests.post(AAD_TOKEN_URL, data=body)
    jsonResponse = json.loads(r.text)
    aadToken = jsonResponse["access_token"]

    return aadToken


def queryAzureLogAnaytics(aadToken, workspaceId, alertId):
    headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + aadToken,
        'Prefer': 'include-permissions=true'
    }

    # use the token to query loganalytics
    LA_URL = "https://api.loganalytics.io/v1/workspaces/" + workspaceId + "/query"
    data = {
        "query": "SecurityAlert| where SystemAlertId in(\"" + alertId + "\") "
    }
    req = requests.post(LA_URL, headers=headers, json=data)
    data = json.loads(req.text)

    return data


def getSentinelAlertArtifacts(data):
    artifacts = []
    #every elemnt defines a column element
    columns_list = data["tables"][0]["columns"]
    #every element is a list that contains all elements of one specific alerts
    row_list = data["tables"][0]["rows"]

    #dictionary to map columnames to indices in a rowlist
    column_index = {}
    index = 0
    for c in columns_list:
        column_index[c["name"]] = index
        index += 1

    for row in row_list:
        entities_json = json.loads(row[column_index["Entities"]])
        for artifact in entities_json:
            if artifact['Type'] == "ip":
                address = artifact.get('Address')
                if address:
                    artifacts.append(AlertArtifact(dataType='ip', data=artifact['Address']))
                else:
                    logging.info("Unknown ip entity: " + json.dumps(entities_json, indent=4, sort_keys=True))
            elif artifact['Type'] == "host":
                hostname = artifact.get('HostName')
                azureID = artifact.get('AzureID')
                if hostname:
                    artifacts.append(AlertArtifact(dataType='host', data=artifact['HostName']))
                elif azureID:
                    artifacts.append(AlertArtifact(dataType='host', data=artifact['AzureID']))
                else:
                    logging.info("Unknown host entity: " + json.dumps(entities_json, indent=4, sort_keys=True))
            elif artifact['Type'] == "account":
                name = artifact.get('Name')
                aadUserId = artifact.get('AadUserId')
                if name:
                    artifacts.append(AlertArtifact(dataType='account', data=artifact['Name']))
                elif aadUserId:
                    artifacts.append(AlertArtifact(dataType='account', data=artifact['AadUserId']))
                else:
                    logging.info("Unknown account entity: " + json.dumps(entities_json, indent=4, sort_keys=True))
            else:
                logging.info("Unknown entity: " + json.dumps(entities_json, indent=4, sort_keys=True))
    return artifacts


def getSentinelAlertLinks(data):
    links = []
    #every elemnt defines a column element
    columns_list = data["tables"][0]["columns"]
    #every element is a list that contains all elements of one specific alerts
    row_list = data["tables"][0]["rows"]

    #dictionary to map columnames to indices in a rowlist
    column_index = {}
    index = 0
    for c in columns_list:
        column_index[c["name"]] = index
        index += 1

    for row in row_list:
        if row[column_index["ExtendedLinks"]] != "":
            links_json = json.loads(row[column_index["ExtendedLinks"]])
            for link in links_json:
                link_to_dashboard = str(link["Href"])
                links.append(link_to_dashboard)
    return links


def CreateHiveAlertFromSentinel(api, title, description, incidentnumber, severity, source, artifacts, alertIds, incidentURL):
    tags = []

    if severity == "Low":
        theHiveSeverity = 1
    elif severity == "Informational":
        theHiveSeverity = 1
    elif severity == "Medium":
        theHiveSeverity = 2
    elif severity == "High":
        theHiveSeverity = 3

    alertIdsStr = ' '.join(map(str, alertIds))

    customFields = CustomFieldHelper()
    customFields.add_number('sentinelIncidentNumber', incidentnumber)
    customFields.add_string('alertIds', alertIdsStr)
    customFields.add_string('incidentURL', incidentURL)

    customFields = customFields.build()
    alert = Alert(title=title,
              tlp=2,
              tags=tags,
              description=description,
              type='Sentinel',
              severity=theHiveSeverity,
              source='Sentinel:'+ source,
              customFields=customFields,
              sourceRef="Sentinel"+str(incidentnumber),
              artifacts=artifacts)

    # Create the Alert
    response = api.create_alert(alert)
    if response.status_code == 201:
        logging.info('Alert created: ' + 'Sentinel' + str(incidentnumber) + ': '+ source + ': ' + title + '. StatusCode: {}/{}'.format(response.status_code, response.text))
    elif (response.status_code == 400 and response.json()['type'] == "ConflictError"):
        logging.info('Duplicate alert: ' +  'Sentinel'+str(incidentnumber) + ': '+ source + ': ' + title + '. StatusCode: {}/{}'.format(response.status_code, response.text))
    else:
        logging.error('failed to create alert: ' + source + ' ' + title + ' Sentinel' +str(incidentnumber) + '. StatusCode: {}/{}'.format(response.status_code, response.text))
        sys.exit(0)


def main():
    inifilename = "sentinel.ini"
    path = os.path.dirname(os.path.abspath(__file__))
    if not os.path.isfile(path + "/" + inifilename):
        sys.exit("sentinel.ini config file not found !")
    config = configparser.ConfigParser()
    config.read(path + "/" + inifilename)

    try:
        THE_HIVE_API_KEY = config["TheHive"]["apiKey"]
        THE_HIVE_SERVER = config["TheHive"]["server"]
        tenantId = config["Azure"]["tenantId"]
        subscriptionId = config["Azure"]["subscriptionId"]
        clientId = config["Azure"]["clientId"]
        clientSecret = config["Azure"]["clientSecret"]
        workspace = config["Azure"]["workspaceName"]
        workspaceId = config["Azure"]["workspaceId"]
        resourceGroup = config["Azure"]["resourceGroupName"]
    except NameError:
        sys.exit("Missing parameters in sentinel.ini")

    #GET the TOKEN
    the_hive_api = TheHiveApi(THE_HIVE_SERVER, THE_HIVE_API_KEY)
    resource = 'https://management.azure.com/'
    authority_url = "https://login.microsoftonline.com/%s" % (tenantId)

    context = adal.AuthenticationContext(authority_url)
    token = context.acquire_token_with_client_credentials(resource, clientId, clientSecret)

    headers = {
        'Content-Type': 'application/json',
        'Authorization': "Bearer " + token['accessToken']
    }

    SE_URL = "https://management.azure.com/subscriptions/" + subscriptionId + "/resourceGroups/" + resourceGroup + "/providers/Microsoft.OperationalInsights/workspaces/" + workspace + "/providers/Microsoft.SecurityInsights/cases?api-version=2019-01-01-preview"
    req = requests.get(SE_URL, headers=headers)
    data = json.loads(req.text)
    logging.debug("Sentinel Incident JSON:" + json.dumps(data, indent=4, sort_keys=True))

    incident_BaseURL = 'https://portal.azure.com/#asset/Microsoft_Azure_Security_Insights/Incident/subscriptions/' + subscriptionId + "/resourceGroups/" + resourceGroup + "/providers/Microsoft.OperationalInsights/workspaces/" + workspace + "/providers/Microsoft.SecurityInsights/Incidents/"

    aadtoken = getAADToken(clientId, clientSecret, tenantId)

    for incident in data['value']:
        if incident['properties']['status'] == "New":
            artifacts = []
            alertIds = []
            incidentURL = str(incident_BaseURL + incident['name'])
            for alertId in incident['properties']['relatedAlertIds']:
                querydata = queryAzureLogAnaytics(aadtoken, workspaceId, alertId)
                artifacts = getSentinelAlertArtifacts(querydata)
                extendedLinks = getSentinelAlertLinks(querydata)
                alertIds.append(alertId)
            description = str(incident['properties']['description']) + \
            "\n\r\n\rLink: " + incidentURL
            if extendedLinks != "":
                for link in extendedLinks:
                    description = description + "\n\rAlertLink: " + str(link)

            CreateHiveAlertFromSentinel(the_hive_api, incident['properties']['title'], description, incident['properties']['caseNumber'], incident['properties']['severity'] , incident['properties']['relatedAlertProductNames'][0], artifacts, alertIds, incidentURL)

if __name__ == "__main__":
    logging.basicConfig(
        level=logging.WARNING,
        format = "%(asctime)s [%(levelname)s] %(message)s"
    )

    main()
