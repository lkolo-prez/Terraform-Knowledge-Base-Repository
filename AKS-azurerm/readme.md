<h1>AKS Azurerm</h1>

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

resource "azurerm_virtual_network" "example" {
  name                = var.vnet_name
  location            = var.location
  resource_group_name = var.resource_group_name
  address_space       = [var.address_space]
}

resource "azurerm_subnet" "example" {
  name                 = var.subnet_name
  resource_group_name  = var.resource_group_name
  virtual_network_name = azurerm_virtual_network.example.name
  address_prefixes     = [var.subnet_address_prefix]
}

resource "azurerm_kubernetes_cluster" "example" {
  name                    = var.cluster_name
  location                = var.location
  resource_group_name     = var.resource_group_name
  dns_prefix              = var.dns_prefix
  kubernetes_version      = var.kubernetes_version
  private_cluster_enabled = var.private_cluster_enabled
  node_resource_group     = var.node_resource_group

  default_node_pool {
    name           = var.node_pool_name
    node_count     = var.node_count
    vm_size        = var.node_vm_size
    vnet_subnet_id = azurerm_subnet.example.id
    enable_auto_scaling = var.enable_auto_scaling
    min_count           = var.min_count
    max_count           = var.max_count
  }

  identity {
    type = "SystemAssigned"
  }

  network_profile {
    network_plugin     = var.network_plugin
    network_policy     = var.network_policy
    dns_service_ip     = var.dns_service_ip
    docker_bridge_cidr = var.docker_bridge_cidr
    service_cidr       = var.service_cidr
  }

  addon_profile {
    azure_policy {
      enabled = var.azure_policy_enabled
    }
    http_application_routing {
      enabled = var.http_application_routing_enabled
    }
    kube_dashboard {
      enabled = var.kube_dashboard_enabled
    }
    oms_agent {
      enabled                    = var.oms_agent_enabled
      log_analytics_workspace_id = var.log_analytics_workspace_id
    }
  }

  role_based_access_control {
    enabled = var.rbac_enabled
  }

  tags = var.tags
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

variable "kubernetes_version" {
  description = "Version of Kubernetes specified when creating the AKS managed cluster"
  default

```hcl
  default     = "1.25"
}

variable "private_cluster_enabled" {
  description = "Whether to enable private cluster"
  default     = false
}

variable "node_resource_group" {
  description = "The name of the resource group where AKS should create its infrastructure"
  default     = "aks-node-resources"
}

variable "node_pool_name" {
  description = "The name of the default node pool"
  default     = "default"
}

variable "node_count" {
  description = "The number of nodes in the default node pool"
  default     = 3
}

variable "node_vm_size" {
  description = "The size of the Virtual Machines to create as part of the default node pool"
  default     = "Standard_D2_v2"
}

variable "vnet_name" {
  description = "The name of the virtual network to create"
  default     = "example-vnet"
}

variable "address_space" {
  description = "The address space of the virtual network"
  default     = "10.0.0.0/16"
}

variable "subnet_name" {
  description = "The name of the subnet within the virtual network"
  default     = "example-subnet"
}

variable "subnet_address_prefix" {
  description = "The address prefix to use for the subnet"
  default     = "10.0.1.0/24"
}

variable "enable_auto_scaling" {
  description = "Whether to enable auto-scaling for nodes in the default node pool"
  default     = false
}

variable "min_count" {
  description = "The minimum number of nodes for auto-scaling"
  default     = 1
}

variable "max_count" {
  description = "The maximum number of nodes for auto-scaling"
  default     = 5
}

variable "network_plugin" {
  description = "Network plugin to use for networking. Can be either 'azure' or 'kubenet'"
  default     = "azure"
}

variable "network_policy" {
  description = "Network policy to use. Can be either 'calico', 'azure', or 'none'"
  default     = "azure"
}

variable "dns_service_ip" {
  description = "An IP address assigned to the Kubernetes DNS service"
  default     = "10.0.0.10"
}

variable "docker_bridge_cidr" {
  description = "A specific IP address and netmask for the Docker bridge, docker0"
  default     = "172.17.0.1/16"
}

variable "service_cidr" {
  description = "A CIDR notation IP range from which to assign service cluster IPs"
  default     = "10.0.0.0/16"
}

variable "azure_policy_enabled" {
  description = "Whether to enable Azure Policy"
  default     = false
}

variable "http_application_routing_enabled" {
  description = "Whether to enable HTTP application routing"
  default     = false
}

variable "kube_dashboard_enabled" {
  description = "Whether to enable Kubernetes Dashboard"
  default     = false
}

variable "oms_agent_enabled" {
  description = "Whether to enable Log Analytics"
  default     = false
}

variable "log_analytics_workspace_id" {
  description = "The ID of a Log Analytics Workspace to send logs to"
  default     = ""
}

variable "rbac_enabled" {
  description = "Whether to enable Kubernetes RBAC"
  default     = true
}

variable "tags" {
  description = "A mapping of tags to assign to the resource"
  default     = {}
}
```

W powyższym kodzie

Terraform, poza tworzeniem samego klastra AKS, definiuje również grupę zasobów, sieć wirtualną i podsieć, które są wymagane do utworzenia klastra AKS. Jest to przykład bardzo rozbudowanego kodu, który definiuje wiele różnych zasobów i konfiguracji.

Zmienne są zdefiniowane w osobnym pliku `variables.tf`, co pozwala na łatwą modyfikację i ponowne użycie kodu. Każda zmienna ma swoją domyślną wartość, ale możesz je zastąpić, dostarczając własne wartości podczas uruchamiania Terraform.

Niektóre z zmiennych, które możesz zechcieć dostosować do swoich potrzeb, to:

- `resource_group_name`, `location`, `cluster_name`, `dns_prefix` - podstawowe informacje o klastrze AKS.
- `kubernetes_version` - wersja Kubernetes dla klastra AKS. Obecnie ustawiona na najnowszą obsługiwaną wersję Kubernetes, czyli 1.25.
- `private_cluster_enabled` - określa, czy klaster powinien być prywatny.
- `node_count`, `node_vm_size` - szczegóły dotyczące domyślnej puli węzłów.
- `enable_auto_scaling`, `min_count`, `max_count` - szczegóły dotyczące automatycznego skalowania dla domyślnej puli węzłów.
- `network_plugin`, `network_policy` - szczegóły dotyczące sieci.
- `azure_policy_enabled`, `http_application_routing_enabled`, `kube_dashboard_enabled`, `oms_agent_enabled` - szczegóły dotyczące różnych dodatków.
- `rbac_enabled` - określa, czy powinno być włączone RBAC (Role-Based Access Control).

Zwróć uwagę, że to jest tylko przykładowy kod i może wymagać dalszych dostosowań do specyficznych potrzeb i zasady bezpieczeństwa twojej organizacji. Przed użyciem tego kodu upewnij się, że dobrze rozumiesz, co robi każda z tych zmiennych i jakie są konsekwencje ich zmiany.
