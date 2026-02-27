output "cluster_name" {
  description = "GKE Cluster Name"
  value       = google_container_cluster.primary.name
}

output "endpoint" {
  description = "GKE Control Plane Endpoint"
  value       = google_container_cluster.primary.endpoint
}