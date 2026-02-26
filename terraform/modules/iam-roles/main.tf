# Staff Pattern: Dedicated service account for the AI Agent
resource "google_service_account" "ai_agent" {
  account_id   = "ai-agent-sa"
  display_name = "AI Agent Service Account"
}

resource "google_project_iam_member" "vertex_user" {
  project = var.project_id
  role    = "roles/aiplatform.user"
  member  = "serviceAccount:${google_service_account.ai_agent.email}"
}