resource "azurerm_key_vault" "thehive-keyvault" {
  name                     = "${var.organization}-thehive-kv"
  location                 = azurerm_resource_group.thehive-rg.location
  resource_group_name      = azurerm_resource_group.thehive-rg.name
  tenant_id                = data.azurerm_client_config.current.tenant_id
  soft_delete_enabled      = true
  purge_protection_enabled = false

  sku_name = "standard"

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    secret_permissions = [
      "set",
      "get",
      "delete",
    ]
  }

  network_acls {
    default_action = "Allow"
    bypass         = "AzureServices"
  }
}

resource "azurerm_key_vault_secret" "thehive-apikey" {
  name         = "thehive-apikey"
  value        = var.thehive_api_key
  key_vault_id = azurerm_key_vault.thehive-keyvault.id
}