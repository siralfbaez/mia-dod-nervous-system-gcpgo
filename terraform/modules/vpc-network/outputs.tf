output "vpc_id" {
  description = "The ID of the VPC"
  value       = google_compute_network.main.id
}

output "subnet_id" {
  description = "The ID of the subnetwork"
  value       = google_compute_subnetwork.main.id
}