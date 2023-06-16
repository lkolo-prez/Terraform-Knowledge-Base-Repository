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
