# Terraform Knowledge Base Repository

This repository serves as my knowledge base for Terraform technology. It contains various Terraform configuration files (`.tf`) based on my experiences, primarily focused on Azure infrastructure provisioning. Here, you will find helpful resources, configurations, and examples to aid in understanding and working with Terraform.

## What is Terraform?

Terraform is an open-source infrastructure as code (IaC) tool that allows you to define and provision infrastructure resources declaratively. It provides a consistent and reliable way to manage infrastructure across various cloud providers and services.

With Terraform, you can define your desired infrastructure configuration in code using a simple and readable language. It automates the deployment and management of resources, ensuring consistency and reducing the chance of manual errors. Terraform tracks the state of your infrastructure, enabling you to make changes safely and efficiently.

## Examples

Here are a few simple examples of Terraform code commonly used to provision infrastructure resources:

1. **Azure Resource Group**:

```hcl
resource "azurerm_resource_group" "example" {
  name     = "my-resource-group"
  location = "West Europe"
}
```

2. **Azure Virtual Network**:

```hcl
resource "azurerm_virtual_network" "example" {
  name                = "my-virtual-network"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
}
```

3. **Azure Virtual Machine**:

```hcl
resource "azurerm_virtual_machine" "example" {
  name                  = "my-virtual-machine"
  location              = azurerm_resource_group.example.location
  resource_group_name   = azurerm_resource_group.example.name
  vm_size               = "Standard_DS2_v2"
  network_interface_ids = [azurerm_network_interface.example.id]
  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
  os_disk {
    name              = "osdisk"
    caching           = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }
}
```

Feel free to explore more Terraform examples within this repository to learn and experiment with different infrastructure provisioning scenarios.

## Contributing

If you have any suggestions, improvements, or additional examples that you would like to contribute to this knowledge base, please feel free to submit a pull request. Your contributions are greatly appreciated!

## License

This repository is licensed under the [MIT License](LICENSE).
