resource "random_password" "psswrd-adoagnt-ais" {
  count            = var.build_agents.enabled == true ? 1 : 0
  length           = 12
  special          = true
  override_special = "_%@"
  min_upper        = 1
  min_lower        = 1
  min_special      = 1
}



resource "azurerm_key_vault_secret" "scrt-adoagnt-ais" {
  count = var.build_agents.enabled == true ? 1 : 0
  depends_on = [
    azurerm_role_assignment.tf-akv,
  ]
  name         = "ado-agents-password"
  value        = random_password.psswrd-adoagnt-ais[0].result
  key_vault_id = azurerm_key_vault.aks-kvt.id
}

resource "azurerm_linux_virtual_machine_scale_set" "vmss-adoagnt-ais" {
  count                           = var.build_agents.enabled == true ? 1 : 0
  name                            = "${var.org_name}-${lookup(local.region, var.location)}-${var.env}-${var.tier}-vmss-ado-01"
  resource_group_name             = azurerm_resource_group.rgp-adoagnt-ais[0].name
  location                        = azurerm_resource_group.rgp-adoagnt-ais[0].location
  sku                             = local.build_agents.sku
  instances                       = 1
  admin_username                  = local.build_agents.admin_username
  disable_password_authentication = false
  admin_password                  = random_password.psswrd-adoagnt-ais[0].result
  upgrade_mode                    = "Manual"
  single_placement_group          = false
  overprovision                   = false

  tags = var.tags

  source_image_id = local.build_agents.use_custom_image ? local.build_agents.source_image_id : null

  dynamic "source_image_reference" {
    for_each = local.build_agents.use_custom_image == false ? [
      {
        publisher = local.build_agents.source_image_reference.publisher
        offer     = local.build_agents.source_image_reference.offer
        sku       = local.build_agents.source_image_reference.sku
        version   = local.build_agents.source_image_reference.version
      }
    ] : []
    content {
      publisher = source_image_reference.value["publisher"]
      offer     = source_image_reference.value["offer"]
      sku       = source_image_reference.value["sku"]
      version   = source_image_reference.value["version"]
    }
  }

  os_disk {
    storage_account_type = "StandardSSD_LRS"
    caching              = "ReadWrite"
  }

  network_interface {
    name    = "${var.org_name}-${lookup(local.region, var.location)}-${var.env}-${var.tier}-nic-agents-01"
    primary = true

    ip_configuration {
      name      = "${var.org_name}-${lookup(local.region, var.location)}-${var.env}-${var.tier}-ipcnfg-01"
      primary   = true
      subnet_id = local.ado_subnet
    }

  }

  lifecycle {
    ignore_changes = [
      instances,
      tags,
    ]
  }
}