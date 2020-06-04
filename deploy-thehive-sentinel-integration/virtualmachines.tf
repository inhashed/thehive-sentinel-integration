resource "azurerm_public_ip" "thehive-pip" {
  name                = "${var.organization}-thehive-pip"
  location            = azurerm_resource_group.thehive-rg.location
  resource_group_name = azurerm_resource_group.thehive-rg.name
  allocation_method   = "Dynamic"
  domain_name_label   = "${var.organization}-thehive"
}

resource "azurerm_network_interface" "thehive-nic" {
  name                = "${var.organization}-thehive-nic"
  location            = azurerm_resource_group.thehive-rg.location
  resource_group_name = azurerm_resource_group.thehive-rg.name

  ip_configuration {
    name                          = "${var.organization}-thehive-nic-ipconf"
    subnet_id                     = azurerm_subnet.thehive-subnet.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.thehive-pip.id
  }
}

resource "azurerm_linux_virtual_machine" "thehive-vm" {
  name                = "${var.organization}-thehive-vm"
  resource_group_name = azurerm_resource_group.thehive-rg.name
  location            = azurerm_resource_group.thehive-rg.location
  size                = "Standard_DS2_v2"
  admin_username      = var.thehive_admin_user
  network_interface_ids = [
    azurerm_network_interface.thehive-nic.id,
  ]

  admin_ssh_key {
    username   = var.thehive_admin_user
    public_key = var.ssh_key_public
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Premium_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }
}

resource "null_resource" "run_ansible" {
  depends_on = [azurerm_linux_virtual_machine.thehive-vm]
  triggers = {
    always_run = "${timestamp()}"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt install python3 -y",
    ]

    connection {
      host        = azurerm_public_ip.thehive-pip.ip_address
      type        = "ssh"
      user        = var.thehive_admin_user
      private_key = file(var.ssh_key_private)
      timeout     = "2m"

    }
  }

  provisioner "local-exec" {
    working_dir = "ansible"
    command     = "ansible-playbook -u ${var.thehive_admin_user} -i '${azurerm_public_ip.thehive-pip.ip_address},' --private-key ${var.ssh_key_private} deploy_thehive.yml --extra-vars \"organization=${var.organization} sslcert_commonname=${var.sslcert_commonname} cortex_api_key=${var.cortex_api_key} thehivesentinelhooks=${var.thehivesentinelhooks} thehivesentinelhooks_logLevel=${var.thehivesentinelhooks_logLevel} resolvedCaseURL=${var.resolvedCaseURL} ignoredAlertURL=${var.ignoredAlertURL} importedAlertURL=${var.importedAlertURL} newCaseURL=${var.newCaseURL} newAlertURL=${var.newAlertURL} thehivesentinelincidents=${var.thehivesentinelincidents} thehive_api_key=${var.thehive_api_key} sentinel_tenantId=${var.sentinel_tenantId} sentinel_subscriptionId=${var.sentinel_subscriptionId} sentinel_clientId=${var.sentinel_clientId} sentinel_clientSecret=${var.sentinel_clientSecret} sentinel_resourceGroupName=${var.sentinel_resourceGroupName} sentinel_workspaceName=${var.sentinel_workspaceName} sentinel_workspaceId=${var.sentinel_workspaceId} ansible_python_interpreter=/usr/bin/python3\""
  }
}