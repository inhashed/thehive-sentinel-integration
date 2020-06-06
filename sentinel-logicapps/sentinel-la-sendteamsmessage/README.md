# sentinel-la-sendteamsmessage

Sentinel-la-sendteamsmessage is a Logic App with a webrequest trigger to be used together with [thehivesentinelhooks](../../thehive-sentinel-hooks). The URL of the webrequest trigger should be set as the newAlertURL parameter.

## Request body

The Webrequest trigger in the Logic App expects a request body with the following json structure:

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

## Parameters

#### lowseverityteamswebhook

URL of the incomming Teams webhook for low severity incidents.

#### mediumseverityteamswebhook

URL of the incomming Teams webhook for medium severity incidents.

#### highseverityteamswebhook

URL of the incomming Teams webhook for high severity indicents.

#### thehivebaseurl

The base url of a TheHive instance. In the teams message, a link to a TheHive instance is provided based on the organziation name: \<orangization\>-thehive-\<thehivebaseurl\>

## Teams webhook

To send a message in a Teams channel, a incomming webhook needs to be created for each severity. See the [Microsoft Teams documentation](https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook).
