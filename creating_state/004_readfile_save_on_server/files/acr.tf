resource "azurerm_container_registry" "acr" {
  count               = var.container_registry.deploy ? 1 : 0
  name                = "${var.org_name}${lookup(local.region, var.location)}${var.env}acr${var.tier}01"
  resource_group_name = azurerm_resource_group.sup_rg.name
  location            = azurerm_resource_group.sup_rg.location
  sku                 = "Premium"
  admin_enabled       = false
  network_rule_set {
    default_action = "Deny"
    ip_rule = [
      for ip in local.container_registry.allowed_ips : {
        action   = "Allow"
        ip_range = ip
      }
    ]

    virtual_network = [
      for subnet in local.container_registry.virtual_network_subnet_ids : {
        action    = "Allow"
        subnet_id = subnet
      }
    ]
  }
  tags = var.tags
}

resource "azurerm_private_endpoint" "acr" {
  count               = var.container_registry.deploy ? 1 : 0
  name                = "${var.org_name}-${lookup(local.region, var.location)}-${var.env}-pep-${var.tier}-acr-01"
  location            = azurerm_resource_group.sup_rg.location
  resource_group_name = azurerm_resource_group.sup_rg.name
  subnet_id           = local.pep_subnet
  private_dns_zone_group {
    name                 = "default"
    private_dns_zone_ids = [var.acr_dns_zone_id]
  }

  private_service_connection {
    name                           = "${var.org_name}-${lookup(local.region, var.location)}-${var.env}-psc-${var.tier}-acr-01"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_container_registry.acr[0].id
    subresource_names              = ["registry"]
  }
}