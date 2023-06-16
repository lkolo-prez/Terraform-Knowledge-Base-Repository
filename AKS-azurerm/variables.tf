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

