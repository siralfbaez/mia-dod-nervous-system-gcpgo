resource "google_alloydb_cluster" "main" {
  cluster_id = "nervous-system-cluster"
  location   = var.region

  network_config {
    network = var.vpc_id
  }

  initial_user {
    password = var.db_password
  }
}

resource "google_alloydb_instance" "primary" {
  cluster       = google_alloydb_cluster.main.name
  instance_id   = "primary-instance"
  instance_type = "PRIMARY"

  machine_config {
    cpu_count = 2 # Start lean for dev
  }
}