# sentinel-la-closeincidentfromcase

Sentinel-la-closeincidentfromcase is a Logic App with a HTTP webrequest trigger to be used together with [thehivesentinelhooks](../../thehive-sentinel-hooks). The URL of the HTTP webrequest trigger should be set as the resolvedCaseURL parameter.

## Request body

The HTTP webrequest trigger of the Logic App expects a request body with the following json structure:

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

#### sentinel-la-closeincidentfromalertURL

The URL of the HTTP webrequest trigger of sentinel-la-closeincidentfromalert Logic Apps

#### thehiveurl

The URL of a TheHive instance.
