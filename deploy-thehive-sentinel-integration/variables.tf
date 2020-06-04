variable "az_subscription_id" {
  description = "Azure subscription id"
}

variable "az_client_id" {
  description = "Azure client id"
}

variable "az_client_secret" {
  description = "Azure client secret"
}

variable "az_tenant_id" {
  description = "Azure tenant id"
}

variable "az_region" {
  default     = "westeurope"
  description = "Azure region"
}

variable "organization" {
  description = "Organization name"
}

variable "whitelist_ip" {
  description = "Public IP to whitelist on NSG"
}

variable "ssh_key_public" {
  description = "Public SSH key"
}

variable "ssh_key_private" {
  description = "Private SSH key"
}

variable "thehive_api_key" {
  description = "API key of TheHive user used by Azure Logic Apps"
}

variable "thehive_admin_user" {
  description = "admin user for ThehiveVM"
  default     = "nvisoadmin"
}

variable "sslcert_commonname" {
  description = ""
  default     = "test.lab"
}

variable "cortex_api_key" {
  description = ""
}

variable "thehivesentinelhooks" {
  default = "true"
}

variable "thehivesentinelhooks_logLevel" {
  default = "trace"
}

variable "resolvedCaseURL" {
  description = ""
}

variable "ignoredAlertURL" {
  description = ""
}

variable "importedAlertURL" {
  description = ""
}

variable "newCaseURL" {
  description = ""
}

variable "newAlertURL" {
  description = ""
}

variable "thehivesentinelincidents" {
  default = "true"
}

variable "sentinel_tenantId" {
  description = ""
}

variable "sentinel_subscriptionId" {
  description = ""
}

variable "sentinel_clientId" {
  description = ""
}

variable "sentinel_clientSecret" {
  description = ""
}

variable "sentinel_resourceGroupName" {
  description = ""
}

variable "sentinel_workspaceName" {
  description = ""
}

variable "sentinel_workspaceId" {
  description = ""
}