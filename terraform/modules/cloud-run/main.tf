resource "google_cloud_run_v2_service" "default" {
  name     = var.service_name
  location = var.region
  ingress  = "INGRESS_TRAFFIC_INTERNAL_ONLY" # Staff move: Keep it off the public internet

  template {
    containers {
      image = var.container_image
      env {
        name  = "GCP_PROJECT_ID"
        value = var.project_id
      }
      resources {
        limits = {
          cpu    = "1"
          memory = "512Mi"
        }
      }
    }

    # Connect to the VPC so it can talk to AlloyDB
    vpc_access {
      connector = var.vpc_connector_id
      egress    = "ALL_TRAFFIC"
    }
  }
}

output "service_url" {
  value = google_cloud_run_v2_service.default.uri
}