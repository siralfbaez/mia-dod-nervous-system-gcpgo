output "validator_uri" {
  description = "The URL of the deployed Cloud Run service"
  value       = google_cloud_run_v2_service.default.uri
}

output "validator_id" {
  description = "The unique identifier for the service"
  value       = google_cloud_run_v2_service.default.id
}