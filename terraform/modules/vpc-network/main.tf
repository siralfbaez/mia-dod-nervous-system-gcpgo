resource "google_compute_network" "main" {
  name                    = "nervous-system-vpc"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "main" {
  name          = "nervous-system-subnet"
  ip_cidr_range = "10.0.0.0/16"
  network       = google_compute_network.main.id
  region        = var.region
}

output "vpc_id"    { value = google_compute_network.main.id }
output "subnet_id" { value = google_compute_subnetwork.main.id }