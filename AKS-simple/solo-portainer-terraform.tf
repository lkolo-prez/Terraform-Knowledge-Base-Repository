resource "kubernetes_namespace" "portainer" {
  metadata {
    name = "portainer"
  }
}

resource "kubernetes_service_account" "portainer_sa" {
  automount_service_account_token = true
  metadata {
    name      = "portainer-sa-clusteradmin"
    namespace = kubernetes_namespace.portainer.metadata[0].name
    labels = {
      "app.kubernetes.io/name"     = "portainer"
      "app.kubernetes.io/instance" = "portainer"
      "app.kubernetes.io/version"  = "ce-latest-ee-2.18.3"
    }
  }
}


resource "kubernetes_persistent_volume_claim" "portainer_pvc" {
  metadata {
    name        = "portainer"
    namespace   = kubernetes_namespace.portainer.metadata[0].name
    annotations = {"volume.alpha.kubernetes.io/storage-class" = "generic"}
    labels = {
      "io.portainer.kubernetes.application.stack" = "portainer"
      "app.kubernetes.io/name"                    = "portainer"
      "app.kubernetes.io/instance"                = "portainer"
      "app.kubernetes.io/version"                 = "ce-latest-ee-2.18.3"
    }
  }
  spec {
    access_modes = ["ReadWriteOnce"]
    resources {
      requests = {
        storage = "10Gi"
      }
    }
  }
}


resource "kubernetes_cluster_role_binding" "portainer_crb" {
  metadata {
    name = "portainer"
    labels = {
      "app.kubernetes.io/name"     = "portainer"
      "app.kubernetes.io/instance" = "portainer"
      "app.kubernetes.io/version"  = "ce-latest-ee-2.18.3"
    }
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
  subject {
    kind      = "ServiceAccount"
    name      = kubernetes_service_account.portainer_sa.metadata[0].name
    namespace = kubernetes_namespace.portainer.metadata[0].name
  }
}


resource "kubernetes_service" "portainer_service" {
  metadata {
    name      = "portainer"
    namespace = kubernetes_namespace.portainer.metadata[0].name
    labels = {
      "io.portainer.kubernetes.application.stack" = "portainer"
      "app.kubernetes.io/name"                    = "portainer"
      "app.kubernetes.io/instance"                = "portainer"
      "app.kubernetes.io/version"                 = "ce-latest-ee-2.18.3"
    }
    annotations = {
      "service.beta.kubernetes.io/azure-load-balancer-internal" = "true"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name"     = "portainer"
      "app.kubernetes.io/instance" = "portainer"
    }
    port {
      name        = "http"
      port        = 9000
      target_port = 9000
    }
    port {
      name        = "https"
      port        = 443
      target_port = 443
    }
    port {
      name        = "edge"
      port        = 8000
      target_port = 8000
    }
    type = "LoadBalancer"
  }
}


resource "kubernetes_deployment" "portainer_deployment" {
  metadata {
    name      = "portainer"
    namespace = kubernetes_namespace.portainer.metadata[0].name
    labels = {
      "io.portainer.kubernetes.application.stack" = "portainer"
      "app.kubernetes.io/name"                    = "portainer"
      "app.kubernetes.io/instance"                = "portainer"
      "app.kubernetes.io/version"                 = "ce-latest-ee-2.18.3"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name"     = "portainer"
        "app.kubernetes.io/instance" = "portainer"
      }
    }
    template {
      metadata {
        labels = {
          "app.kubernetes.io/name"     = "portainer"
          "app.kubernetes.io/instance" = "portainer"
        }
      }
      spec {
        service_account_name = kubernetes_service_account.portainer_sa.metadata[0].name
        volume {
          name = "data"
          persistent_volume_claim {
            claim_name = kubernetes_persistent_volume_claim.portainer_pvc.metadata[0].name
          }
        }
        container {
          image   = "portainer/portainer-ce:2.17.1"
          name    = "portainer"
          command = ["/portainer"]
          args    = ["-H", "unix:///var/run/docker.sock"]
          port {
            container_port = 9000
            name           = "http"
          }
          port {
            container_port = 443
            name           = "https"
          }
          port {
            container_port = 8000
            name           = "edge"
          }
          volume_mount {
            mount_path = "/data"
            name       = "data"
          }
          resources {}
        }
      }
    }
  }
}

