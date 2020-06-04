# thehive-sentinel-hooks

thehive-sentinel-hooks is an application written in Golang created for integrating TheHive with Azure Sentinel. It's goal is to create a 2-way synchronization between incidents and playbooks in Azure Sentinel and alerts, cases and tasks in TheHive. It is intented to be used together with the [thehive-sentinel-incidents](../thehive-sentinel-incidents/README.md) script and the Azure Logic Apps in [sentinel-logicapps](../sentinel-logicapps/README.md) but can also be used independently.

Thehive-sentinel-hooks should set as the HTTP endpoint for TheHive webhooks and based on the properties of the actions taken in TheHive, it performs a HTTP POST request to a URL specified in the config.yml configuration file. 

The following use cases have been implemented based on actions performed in TheHive:

- Call a HTTP POST request when a new TheHive alert is created
- Call a HTTP POST request when a TheHive alert is set to "Mark as read"
- Call a HTTP POST request when a TheHive alert is imported
- Call a HTTP POST request when a new TheHive case is created
- Call a HTTP Post request when a TheHiveccase is resolved

The HTTP endpoint of thehive-sentinel-hooks listens on port 9002.

## config.yml

Thehive-sentinel-hooks is configured with a config.yml file and can be set by using the -config parameter with a default value of the current directory.

```sh
./main --config "/etc/thehivesentinelhooks/config.yml"
```

Config.yml:

```yaml
newAlertURL: 
ignoredAlertURL: 
importedAlertURL:
newCaseURL: 
resolvedCaseURL:
organization: 
logLevel: 
```

#### newAlertURL

URL that is called in a HTTP POST request when a new TheHive alert is created.

#### ignoredAlertURL

URL that is called in a HTTP POST request when a TheHive alert is set to "Mark as read".

#### importedAlertURL

URL that is called in a HTTP POST request when a TheHive alert is imported to a case.

#### newCaseURL

URL that is called in a HTTP POST request when a new TheHive case is created.

#### resolvedCaseURL

URL that is called in a HTTP POST request when a TheHive case is set to resolved.

#### organization

The organization name that will be added to the JSON body of the HTTP POST request.

#### logLevel

The logging level of thehive-sentinel-hooks can be set to the following values:
- info: displays all info, error and warning messages
- debug: displays all info, error, warning and debug messages
- trace: displays all info, error, warning, debug and trace messages i.e. JSON body from TheHive Webhooks and HTTP POST request calls

## HTTP POST Request

A HTTP Post request is called when one of these actions is performed in TheHive:
- A new TheHive alert is created
- A TheHive alert is set to "Mark as read"
- A new TheHive case is created
- A TheHiveccase is resolved

### Request Body

The following JSON structure is created based on the JSON body of TheHive webhooks and passed as the body of the HTTP POST request:

```json
{
    "operation": "",
    "objectType": "",
    "objectId": "",
    "details": {
        "status": "",
        "caseId": 0
    },
    "object": {
        "title": "",
        "description": "",
        "owner": "",
        "resolutionStatus": "",
        "resolutionSummary": "",
        "severity": 0,
        "customFields": {
            "numberStruct": {
                "incidentNumber": 0
            },
            "alertStruct": {
                "alertIds": ""
            },
            "URLStruct": {
                "incidentURL": ""
            }
        },
        "tlp": 0,
        "source": ""
    },
    "organization": ""
}
```

## TheHive Configuration

### Webhooks

TheHive webhooks are configured using the `webhook` key in the configuration file (`/etc/thehive/application.conf`). Set thehive-sentinel-hooks as the HTTP Endpoint.

```
webhooks {
  myLocalWebHook {
    url = "http://thehivesentinelhooks:9002"
  }
}
```

### Custom Fields

When the thehive-sentinel-incidents script is used to pull Sentinel incidents into TheHive as alerts, custom fields are populated. These custom fields need to be created in TheHive and added to the case templates.

#### SentinelIncidentNumber

The incident number of the Azure Sentinel incident.

#### AlertIds

All Azure Sentinel Alert ID's from the Azure Sentinel incident.

#### IncidentURL

The URL of the Azure Sentinel incident.

## Build 

To build the Golang application use the following command line:

```sh
CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-w -s" -installsuffix cgo -o main .
```

## Docker

Using the [Dockerfile](Dockerfile), you can create a Docker container of thehive-sentinel-hooks.

A Docker container of thehive-sentinel-hooks is available on Docker Hub.

```sh
docker run -v /opt/thehivesentinelhooks/etc:/etc/thehivesentinelhooks -d -p 9002:9002 --name thehivesentinelhooks wstinkens/thehivesentinelhooks:latest
```

## Ansible

To deploy a TheHive/Cortex instance with thehive-sentinel-hooks preconfigured, you can use the [ansible-thehive](https://github.com/NVISO-BE/ansible-thehive) Ansible role available on the NVISO GitHub.