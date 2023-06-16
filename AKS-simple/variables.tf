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
