
Plik `main.tf`:

```hcl
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_kubernetes_cluster" "example" {
  name                = "example-aks"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "example-aks"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }
}

provider "kubernetes" {
  host                   = azurerm_kubernetes_cluster.example.kube_config.0.host
  client_certificate     = base64decode(azurerm_kubernetes_cluster.example.kube_config.0.client_certificate)
  client_key             = base64decode(azurerm_kubernetes_cluster.example.kube_config.0.client_key)
  cluster_ca_certificate = base64decode(azurerm_kubernetes_cluster.example.kube_config.0.cluster_ca_certificate)
}

resource "kubernetes_deployment" "example" {
  metadata {
    name = "nginx-deployment"
  }

  spec {
    replicas = 3

    selector {
      match_labels = {
        app = "nginx"
      }
    }

    template {
      metadata {
        labels = {
          app = "nginx"
        }
      }

      spec {
        container {
          image = "nginx:1.15.8"

          name = "nginx"

          port {
            container_port = 80
          }
        }
      }
    }
  }
}
```

Plik `variables.tf`:

```hcl
variable "resource_group_name" {
  description = "The name of the resource group in which to create the resources"
  default     = "example-resources"
}

variable "location" {
  description = "The location/region in which to create the resources"
  default     = "West Europe"
}

variable "cluster_name" {
  description = "The name of the AKS cluster"
  default     = "example-aks"
}

variable "dns_prefix" {
  description = "The DNS prefix for the AKS cluster"
  default     = "example-aks"
}
```

Pamiętaj, że ten kod jest tylko przykładem i może wymagać dalszych modyfikacji w zależności od specyficznych wymagań i zasad bezpieczeństwa Twojej organizacji.
