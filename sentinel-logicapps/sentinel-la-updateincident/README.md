# sentinel-la-updateincident

Sentinel-la-updateincident is a Logic App with a HTTP webrequest trigger to be used together with [thehivesentinelhooks](../../thehive-sentinel-hooks). The URL of the HTTP webrequest trigger should be set as the importedAlertURL parameter.

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

#### sentinelResourceGroupName

The name of the Azure resource group of Azure Sentinel.

#### sentinelSubscriptionId

The id of the Azure subscription of Azure Sentinel.

#### sentinelWorkspaceId

The id of the Azure Log Analytics workspace of Azure Sentinel.