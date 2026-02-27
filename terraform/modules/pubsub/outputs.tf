output "topic_name" {
  description = "The name of the signal topic"
  value       = google_pubsub_topic.signal_topic.name
}

output "subscription_name" {
  description = "The name of the worker subscription"
  value       = google_pubsub_subscription.signal_subscription.name
}