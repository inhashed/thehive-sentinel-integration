resource "azurerm_virtual_network" "thehive-vnet" {
  name                = "${var.organization}-thehive-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.thehive-rg.location
  resource_group_name = azurerm_resource_group.thehive-rg.name
}

resource "azurerm_subnet" "thehive-subnet" {
  name                 = "${var.organization}-thehive-subnet"
  resource_group_name  = azurerm_resource_group.thehive-rg.name
  virtual_network_name = azurerm_virtual_network.thehive-vnet.name
  address_prefixes     = ["10.0.0.0/24"]
}

resource "azurerm_network_security_group" "thehive-nsg" {
  name                = "${var.organization}-thehive-nsg"
  location            = azurerm_resource_group.thehive-rg.location
  resource_group_name = azurerm_resource_group.thehive-rg.name

  security_rule {
    name                       = "AllowHTTPNVISO"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "80"
    source_address_prefix      = var.whitelist_ip
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "AllowHTTPSNVISO"
    priority                   = 110
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "443"
    source_address_prefix      = var.whitelist_ip
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "AllowSSHNVISO"
    priority                   = 120
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "22"
    source_address_prefix      = var.whitelist_ip
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "AllowTheHiveLogicApps"
    priority                   = 130
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "9000"
    source_address_prefix      = "LogicApps"
    destination_address_prefix = "*"
  }
}

resource "azurerm_subnet_network_security_group_association" "thehive-subnet-nsg-association" {
  subnet_id                 = azurerm_subnet.thehive-subnet.id
  network_security_group_id = azurerm_network_security_group.thehive-nsg.id
}