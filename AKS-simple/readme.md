Poniżej znajduje się przykład jak można zdefiniować cluster AKS (Azure Kubernetes Service) w Terraformie. Kod jest podzielony na dwa pliki: jeden definiuje zasób AKS (main.tf), a drugi zawiera wszystkie zmienne i ich wartości domyślne (variables.tf).

Plik main.tf:

```hcl
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>2.46"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = var.resource_group_name
  location = var.location
}

resource "azurerm_kubernetes_cluster" "example" {
  name                = var.cluster_name
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = var.dns_prefix

  default_node_pool {
    name       = var.node_pool_name
    node_count = var.node_count
    vm_size    = var.node_vm_size
  }

  identity {
    type = "SystemAssigned"
  }
}
```

Plik variables.tf:

```hcl
variable "resource_group_name" {
  description = "The name of the resource group in which to create the Kubernetes Cluster"
  default     = "example-resources"
}

variable "location" {
  description = "The location/region in which to create the resource group"
  default     = "West Europe"
}

variable "cluster_name" {
  description = "The name of the Kubernetes Cluster"
  default     = "example-aks"
}

variable "dns_prefix" {
  description = "DNS prefix specified when creating the managed cluster"
  default     = "exampleaks"
}

variable "node_pool_name" {
  description = "The name of the default node pool"
  default     = "default"
}

variable "node_count" {
  description = "The number of nodes in the default node pool"
  default     = 1
}

variable "node_vm_size" {
  description = "The size of the Virtual Machines to create as part of the default node pool"
  default     = "Standard_D2_v2"
}
```

Ten kod tworzy grupę zasobów oraz klaster AKS z domyślnym pulą węzłów. Wszystkie możliwe parametry są zdefiniowane w pliku variables.tf z domyślnymi wartościami, które można dostosować do swoich potrzeb.
