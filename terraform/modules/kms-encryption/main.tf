resource "google_kms_key_ring" "main" {
  name     = "nervous-system-keyring"
  location = var.region
}

resource "google_kms_crypto_key" "main" {
  name            = "integration-key"
  key_ring        = google_kms_key_ring.main.id
  rotation_period = "7776000s" # 90 days rotation for compliance

  lifecycle {
    prevent_destroy = true
  }
}

output "key_id" {
  value = google_kms_crypto_key.main.id
}