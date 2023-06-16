
1. **Utwórz grupę zasobów**: Stwórz grupę zasobów w określonej lokalizacji.
   ```hcl
   resource "azurerm_resource_group" "example" {
     name     = "example-resources"
     location = "West Europe"
   }
   ```

2. **Utwórz wirtualną sieć**: Stwórz wirtualną sieć z określonym zakresem adresów IP.
   ```hcl
   resource "azurerm_virtual_network" "example" {
     name                = "example-network"
     resource_group_name = azurerm_resource_group.example.name
     location            = azurerm_resource_group.example.location
     address_space       = ["10.0.0.0/16"]
   }
   ```

3. **Utwórz podsieć**: Stwórz podsieć w ramach wirtualnej sieci.
   ```hcl
   resource "azurerm_subnet" "example" {
     name                 = "example-subnet"
     resource_group_name  = azurerm_resource_group.example.name
     virtual_network_name = azurerm_virtual_network.example.name
     address_prefixes     = ["10.0.1.0/24"]
   }
   ```

4. **Utwórz maszynę wirtualną**: Stwórz maszynę wirtualną z określonym obrazem i rozmiarem.
   ```hcl
   resource "azurerm_virtual_machine" "example" {
     name                  = "example-vm"
     location              = azurerm_resource_group.example.location
     resource_group_name  = azurerm_resource_group.example.name
     network_interface_id = azurerm_network_interface.example.id
     vm_size               = "Standard_DS1_v2"
     delete_os_disk_on_termination = true

     storage_image_reference {
       publisher = "Canonical"
       offer     = "UbuntuServer"
       sku       = "16.04-LTS"
       version   = "latest"
     }

     storage_os_disk {
       name              = "os-disk"
       caching           = "ReadWrite"
       create_option     = "FromImage"
     }

     os_profile {
       computer_name  = "hostname"
       admin_username = "adminuser"
       admin_password = "adminpassword"
     }

     os_profile_linux_config {
       disable_password_authentication = false
     }

     network_interface_id = azurerm_network_interface.example.id
   }
   ```

5. **Utwórz interfejs sieciowy**: Stwórz interfejs sieciowy z publicznym adresem IP.
   ```hcl
   resource "azurerm_network_interface" "example" {
     name                = "example-nic"
     location            = azurerm_resource_group.example.location
     resource_group_name = azurerm_resource_group.example.name

     ip_configuration {
       name                          = "internal"
       subnet_id                     = azurerm_subnet.example.id
       private_ip_address_allocation = "Dynamic"
     }
   }
   ```

6. **Utwórz magazyn bloba**: Stwórz magazyn bloba w określonej grupie zasobów.
   ```hcl
  

resource "azurerm_storage_account" "example" {
  name                     = "examplestorageaccount"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
}
```

7. **Utwórz kontener bloba**: Utwórz kontener bloba w określonym magazynie bloba.
```hcl
resource "azurerm_storage_container" "example" {
  name                  = "examplecontainer"
  storage_account_name  = azurerm_storage_account.example.name
  container_access_type = "private"
}
```

8. **Utwórz zasób App Service**: Utwórz zasób App Service, który hostuje aplikację internetową.
```hcl
resource "azurerm_app_service" "example" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
  }

  app_settings = {
    "SOME_KEY" = "some-value"
  }
}
```

9. **Utwórz plan App Service**: Utwórz plan App Service, który definiuje poziom wydajności aplikacji.
```hcl
resource "azurerm_app_service_plan" "example" {
  name                = "example-app-service-plan"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  sku {
    tier = "Standard"
    size = "S1"
  }
}
```

10. **Utwórz grupę skalowania**: Utwórz grupę skalowania, która automatycznie skaluje liczbę maszyn wirtualnych.
```hcl
resource "azurerm_autoscale_setting" "example" {
  name                = "example-autoscale-setting"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  target_resource_id  = azurerm_virtual_machine_scale_set.example.id

  profile {
    name = "default"

    capacity {
      minimum = "1"
      maximum = "5"
      default = "1"
    }

    rule {
      metric_trigger {
        metric_name        = "Percentage CPU"
        metric_resource_id = azurerm_virtual_machine_scale_set.example.id
        time_grain         = "PT1M"
        statistic          = "Average"
        time_window        = "PT5M"
        time_aggregation   = "Average"
        operator           = "GreaterThan"
        threshold          = 75
      }

      scale_action {
        direction = "Increase"
        type      = "ChangeCount"
        value     = "1"
        cooldown  = "PT1M"
      }
    }
  }
}
```

11. **Utwórz zasób bazy danych SQL**: Utwórz zasób bazy danych SQL w określonej grupie zasobów.
```hcl
resource "azurerm_sql_database" "example" {
  name                = "example-sql-database"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  server_name         = az

urerm_sql_server.example.name
  edition             = "Standard"
  collation           = "SQL_Latin1_General_CP1_CI_AS"
  max_size_bytes      = "1073741824"
  create_mode         = "Default"
}
```

12. **Utwórz serwer SQL**: Utwórz serwer SQL, który hostuje bazę danych SQL.
```hcl
resource "azurerm_sql_server" "example" {
  name                         = "examplesqlserver"
  resource_group_name          = azurerm_resource_group.example.name
  location                     = azurerm_resource_group.example.location
  version                      = "12.0"
  administrator_login          = "4dm1n157r470r"
  administrator_login_password = "4-v3ry-53cr37-p455w0rd"
}
```

13. **Utwórz usługę Kubernetes**: Utwórz usługę Kubernetes z określoną liczbą węzłów.
```hcl
resource "azurerm_kubernetes_cluster" "example" {
  name                = "example-aks"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }
}
```

14. **Utwórz zasób sieci prywatnej wirtualnej (VPN)**: Utwórz zasób sieci VPN dla bezpiecznego połączenia.
```hcl
resource "azurerm_virtual_network_gateway" "example" {
  name                = "examplevpngateway"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  type     = "Vpn"
  vpn_type = "RouteBased"

  active_active = false
  enable_bgp    = false

  sku = "VpnGw1"

  ip_configuration {
    name                          = "vnetGatewayConfig"
    public_ip_address_id          = azurerm_public_ip.example.id
    private_ip_address_allocation = "Dynamic"
    subnet_id                     = azurerm_subnet.example.id
  }

  vpn_client_configuration {
    address_space        = ["10.2.0.0/24"]
    vpn_client_protocols = ["SSTP"]
  }
}
```

15. **Utwórz dysk Managed Disk**: Utwórz dysk Managed Disk do przechowywania danych maszyny wirtualnej.
```hcl
resource "azurerm_managed_disk" "example" {
  name                 = "example-disk"
  location             = azurerm_resource_group.example.location
  resource_group_name  = azurerm_resource_group.example.name
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "100"
}
```

16. **Utwórz grupę bezpieczeństwa sieci**: Utwórz grupę bezpieczeństwa sieci do zarządzania zasadami bezpieczeństwa.
```hcl
resource "azurerm_network_security_group" "example" {
  name                = "example-nsg"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
}
```

17. **Utwó

rz zasób Azure Functions**: Utwórz zasób Azure Functions do uruchamiania bezserwerowych funkcji.
```hcl
resource "azurerm_function_app" "example" {
  name                       = "example-function-app"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_connection_string  = azurerm_storage_account.example.primary_connection_string
}
```

18. **Utwórz Azure Logic App**: Utwórz Logic App do orkiestracji procesów biznesowych i integracji systemów.
```hcl
resource "azurerm_logic_app_workflow" "example" {
  name                = "workflow1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
}
```

19. **Utwórz Azure Service Bus Namespace**: Utwórz Service Bus Namespace dla komunikacji między usługami.
```hcl
resource "azurerm_servicebus_namespace" "example" {
  name                = "example-namespace"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  sku                 = "Standard"
}
```

20. **Utwórz Azure Event Hub**: Utwórz Event Hub do przesyłania dużych ilości danych w czasie rzeczywistym.
```hcl
resource "azurerm_eventhub" "example" {
  name                = "example-eventhub"
  namespace_name      = azurerm_eventhub_namespace.example.name
  resource_group_name = azurerm_resource_group.example.name
  partition_count     = 2
  message_retention   = 1
}
```

Proszę pamiętać, że niektóre z tych przykładów mogą wymagać dodatkowych zasobów lub konfiguracji, które nie są pokazane. Na przykład, tworzenie maszyny wirtualnej może wymagać utworzenia sieci wirtualnej i interfejsu sieciowego, co nie jest pokazane w przykładzie.
