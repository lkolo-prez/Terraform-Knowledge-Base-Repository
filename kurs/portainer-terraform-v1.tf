resource "kubernetes_deployment" "portainer"{
    metadata {
        name = "portainer-deployment"
    }

    spec {
        replicas = 1

        selector {
            match_labels = {
                app = "portainer"
            }
        }
    }

    template {
        metadata {
            labels = {
                app = "portainer"
            }
        }

        spec {
            container {
                image = "ce-latest-ee-2.17.1"
                name = "portainer"
                port {
                    container_port = 443
                }
            }
        }
    }
}
