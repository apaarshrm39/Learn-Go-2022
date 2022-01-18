resource "azurerm_log_analytics_workspace" "fmz-law-aks" {
  count               = var.enable_log_analytics_workspace == true && var.use_existing_log_analytics_workspace == false ? 1 : 0
  name                = "${var.org_name}-${lookup(local.region, var.location)}-${var.env}-log-${var.tier}-aks-01"
  location            = azurerm_resource_group.sup_rg.location
  resource_group_name = azurerm_resource_group.sup_rg.name
  sku                 = "PerGB2018"
  retention_in_days   = 180
  tags                = var.tags
}