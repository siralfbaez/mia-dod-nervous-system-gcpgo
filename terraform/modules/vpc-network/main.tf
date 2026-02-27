resource "google_vpc_access_connector" "connector" {
  name          = "run-connector"
  ip_cidr_range = "10.8.0.0/28"
  network       = google_compute_network.main.name
  region        = var.region
}

output "connector_id" {
  value = google_vpc_access_connector.connector.id
}