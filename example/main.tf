provider azurermext {
  subscription_id = ""
  client_id       = ""
  client_secret   = ""
  tenant_id       = ""
}

provider azurerm {
  subscription_id = ""
  client_id       = ""
  client_secret   = ""
  tenant_id       = ""
}

resource "azurerm_mysql_firewall_rule" "fr" {
  name                = "whitelisted-ip"
  resource_group_name = "rg"
  server_name         = "mysqlserver"
  start_ip_address    = "0.0.0.0"
  end_ip_address      = "0.0.0.1"
}

# Some operation on DB happens here

resource "azurermext_mysql_firewall_rule_rm" "b" {
  depends_on          = ["azurerm_mysql_firewall_rule.fr"]
  name                = "whitelisted-ip"
  resource_group_name = "rg"
  server_name         = "mysqlserver"
}