# sentinel-la-closeASCalert

Sentinel-la-closeASCalert is a Logic App with a HTTP webrequest trigger to be used together with the sentinel-la-closeincidentfromalert Logic App.

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

#### sentinelWorkspaceName

The name of the Azure Log Analytics workspace of Azure Sentinel.

#### userName

Username for Logic App connection.

## Managed Identity

A system-assigned managed identity is created at deployment. The following RBAC permissions should be assigned on the Azure subscription:
- Security Admimnistrator