output "ai_agent_sa_email" {
  description = "The email of the AI Agent service account"
  value       = google_service_account.ai_agent.email
}