# GKE Outputs
output "kubernetes_cluster_name" {
  description = "The name of the GKE Autopilot cluster"
  value       = module.gke.cluster_name
}

output "kubernetes_cluster_endpoint" {
  description = "The IP address of the cluster master"
  value       = module.gke.endpoint
  sensitive   = true
}

# AlloyDB Outputs
output "alloydb_cluster_id" {
  description = "The ID of the AlloyDB cluster"
  value       = module.alloydb.cluster_id
}

output "alloydb_primary_instance_ip" {
  description = "The private IP of the primary writer instance"
  value       = module.alloydb.primary_ip
}

# Pub/Sub Outputs
output "signal_topic_name" {
  description = "The name of the main Pub/Sub topic for signals"
  value       = module.pubsub.topic_name
}

# KMS Outputs
output "kms_key_id" {
  description = "The encryption key for the nervous system"
  value       = module.kms.crypto_key_id
}

# Cloud Run Outputs
output "validator_endpoint" {
  description = "The final endpoint for the Signal Validator"
  value       = module.cloud_run.validator_uri
}
