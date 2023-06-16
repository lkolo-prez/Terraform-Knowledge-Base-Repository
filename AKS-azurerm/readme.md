<h1>AKS Azurerm</h1>


W powyższym kodzie Terraform, poza tworzeniem samego klastra AKS, definiuje również grupę zasobów, sieć wirtualną i podsieć, które są wymagane do utworzenia klastra AKS. Jest to przykład bardzo rozbudowanego kodu, który definiuje wiele różnych zasobów i konfiguracji.

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
