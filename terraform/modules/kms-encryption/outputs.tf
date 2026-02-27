output "key_ring_id" {
  description = "The self-link of the created KMS keyring"
  value       = google_kms_key_ring.main.id
}

output "crypto_key_id" {
  description = "The ID of the KMS crypto key for the nervous system"
  value       = google_kms_crypto_key.main.id
}