# thehive-sentinel-incident

thehive-sentinel-incident is a Python script created to pull Azure Sentinel incidents into TheHive as alerts. It queries the Azure Sentinel API for new incidents and gets the alert entities from the Azure Sentinel Log Analytics workspace to be created as TheHive observables.

## Azure App Registration

An Azure app registration (service principle) is required to query the Azure Sentinel API and the Azure Log analytics workspace. 

See the Azure documentation to create a service principle:
https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal

The app registration (service principle) should have the following permissions on the Azure Log Analytics workspace with Azure Sentinel enabled:
- Sentinel Contributor
- Log Analytics Contributor

## sentinel.ini

thehive-sentinel-incidents is configured in the file sentinel.ini.

```
[TheHive]
apiKey = 
server  = http://localhost:9000

[Azure]
tenantId = 
subscriptionId = 
clientId = 
clientSecret = 
resourceGroupName = 
workspaceName = 
workspaceId = 
```
### TheHive

#### apiKey

The API key of a TheHive user with the following permissions:
- Read/Write role
- Allow alerts creation 

#### server

The URL and port for connecting to TheHive.

### AZure

#### tenantId

The id of the Azure tenant containing the Azure Log Analytics workspace with Azure Sentinel enabled.

#### subscriptionId

The id of the Azure subscription containing the Azure Log Analytics workspace with Azure Sentinel enabled.

#### clientId

The client id of the Azure service principle.

#### clientSecret

The client secret of the Azure service principle.

#### resourceGroupName

The name of the resource group containing the Azure Log Analytics workspace with Azure Sentinel enabled.

#### workspaceName

The name of the Azure Log Analytics workspace with Azure Sentinel enabled.

#### workspaceId

The ID of the Azure Log Analytics workspace with Azure Sentinel enabled.

## Cron

To periodically run thehive-sentinel-incidents, use a cron job:

```sh
*/5 * * * * python3 /opt/thehive-sentinel-incident/thehivesentinelincidents.py >> /var/log/thehivesentinelincidents.log 2>&1
```

## Ansible

To deploy a TheHive/Cortex instance with thehive-sentinel-incidents preconfigured, you can use the [ansible-thehive](https://github.com/NVISO-BE/ansible-thehive) Ansible role available on the NVISO GitHub.