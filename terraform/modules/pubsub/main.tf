resource "google_pubsub_topic" "signal_topic" {
  name = "nervous-system-signals"

  labels = {
    environment = "dev"
    managed_by  = "terraform"
  }
}

# The Dead Letter Topic for "brittle" signals that fail processing
resource "google_pubsub_topic" "signal_dead_letter" {
  name = "nervous-system-signals-dlq"
}

resource "google_pubsub_subscription" "ai_observation_sub" {
  name  = "nervous-system-ai-sub"
  topic = google_pubsub_topic.signal_topic.name

  # AI might take longer to think, so we increase the ack deadline
  ack_deadline_seconds = 60
}

resource "google_pubsub_subscription" "signal_subscription" {
  name  = "nervous-system-sub"
  topic = google_pubsub_topic.signal_topic.name

  # Retain messages for 7 days
  message_retention_duration = "604800s"

  # Dead Letter Policy
  dead_letter_policy {
    dead_letter_topic     = google_pubsub_topic.signal_dead_letter.id
    max_delivery_attempts = 5
  }

  retry_policy {
    minimum_backoff = "10s"
    maximum_backoff = "600s"
  }
}