

resource "azurerm_resource_group" "rg" {
  name     = "your_resource_group"
  location = "West Europe"
}

resource "azurerm_storage_account" "sa" {
  name                     = "your_storage_account_name"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Premium"
  account_replication_type = "LRS"
}

resource "azurerm_storage_share" "ss" {
  name                 = "your_file_share_name"
  storage_account_name = azurerm_storage_account.sa.name
  quota                = 50
}

resource "kubernetes_secret" "example" {
  metadata {
    name = "your_kubernetes_secret_name"
  }

  data = {
    accountname = azurerm_storage_account.sa.name
    accountkey  = azurerm_storage_account.sa.primary_access_key
  }

  type = "kubernetes.io/azure-file" // azure file share
}



resource "azurerm_storage_account" "sa_tfstate" {
  name                     = "your_tfstate_account_name"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Premium"
  account_replication_type = "LRS"
}

resource "azurerm_storage_container" "sc_tfstate" {
  name                  = "tfstate"
  storage_account_name  = azurerm_storage_account.sa_tfstate.name
  container_access_type = "private"
}
