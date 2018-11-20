provider "kubernetes" {
  config_context_auth_info = "ops"
  config_context_cluster   = "mycluster"
}

resource "kubernetes_namespace" "example" {
  metadata {
    name = "my-first-namespace"
  }
}

resource "kubernetes_replication_controller" "example" {
  metadata {
    name = "terraform-example"
    labels {
      test = "MyExampleApp"
    }
  }

  spec {
    selector {
      test = "MyExampleApp"
    }
    template {
      container {
        image = "nginx:1.7.8"
        name  = "example"

        resources{
          limits{
            cpu    = "0.5"
            memory = "512Mi"
          }
          requests{
            cpu    = "250m"
            memory = "50Mi"
          }
        }
      }
    }
  }
}