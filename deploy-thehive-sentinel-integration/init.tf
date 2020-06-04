provider "azurerm" {
  subscription_id = var.az_subscription_id
  client_id       = var.az_client_id
  client_secret   = var.az_client_secret
  tenant_id       = var.az_tenant_id
  features {}
}

data "azurerm_client_config" "current" {}

resource "azurerm_resource_group" "thehive-rg" {
  name     = "${var.organization}-thehive-rg"
  location = var.az_region

}