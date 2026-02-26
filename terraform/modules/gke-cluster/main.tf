resource "google_container_cluster" "primary" {
  name     = "nervous-system-cluster"
  location = var.region

  # Staff Choice: Autopilot manages nodes, scaling, and security hardening
  enable_autopilot = true

  # Networking: Enable VPC-native traffic for better performance
  networking_mode = "VPC_NATIVE"
  network         = var.vpc_id
  subnetwork      = var.subnet_id

  ip_allocation_policy {
    cluster_secondary_range_name  = "pods"
    services_secondary_range_name = "services"
  }

  # Security: Enable Workload Identity for least-privilege service accounts
  workload_identity_config {
    workload_pool = "${var.project_id}.svc.id.goog"
  }

  release_channel {
    channel = "REGULAR"
  }
}