# deploy-thehive-sentinel-integration

Deploy a TheHive/Cortex instance with the Azure Sentinel ingetration with Terraform and Ansible. It installs and configures [thehive-sentinel-hooks](https://github.com/NVISO-BE/thehive-sentinel-integration) to integrate TheHive with Azure Sentinel and [thehive-sentinel-incidents](https://github.com/NVISO-BE/thehive-sentinel-integration) to pull Azure Sentinel incidents to TheHive.

## Azure App Registration

An Azure app registration (service principle) is required to deploy the Azure resources. 

See the Azure documentation to create a service principle:
https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal

The app registration (service principle) should have the following permissions on the Azure subscription:
- Contributor

## terraform.tfvars

Before deployment, a terraform.tfvars file should be created with the following parameters:

```
az_subscription_id         = ""
az_client_id               = ""
az_client_secret           = ""
az_tenant_id               = "
organization               = "NVISO"
whitelist_ip               = ""
ssh_key_public             = ""
thehive_api_key            = ""
ssh_key_private            = "~/.ssh/id_rsa"
cortex_api_key             = ""
resolvedCaseURL            = ""
ignoredAlertURL            = ""
importedAlertURL           = ""
newCaseURL                 = ""
newAlertURL                = ""
sentinel_tenantId          = ""
sentinel_subscriptionId    = ""
sentinel_clientId          = ""
sentinel_clientSecret      = ""
sentinel_resourceGroupName = ""
sentinel_workspaceName     = "
sentinel_workspaceId       = ""
```

## Deploy

Initialize terraform to get the required providers:

```
terraform init
```

Deploy the TheHive instance to Azure:

```
terraform apply
```