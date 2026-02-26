resource "google_alloydb_cluster" "main" {
  cluster_id = "nervous-system-db-cluster"
  location   = var.region
  network_config {
    network = var.vpc_id
  }

  initial_user {
    password = var.db_password
  }
}

# The Primary Instance (The "Writer")
resource "google_alloydb_instance" "primary" {
  cluster       = google_alloydb_cluster.main.name
  instance_id   = "primary-writer"
  instance_type = "PRIMARY"

  machine_config {
    cpu_count = 2 # N2-standard-2 equivalent
  }
}

# The Read Pool (The "Reader")
# Staff Flex: Offload GET requests to this pool to protect the writer
resource "google_alloydb_instance" "read_pool" {
  cluster       = google_alloydb_cluster.main.name
  instance_id   = "read-pool-01"
  instance_type = "READ_POOL"

  read_pool_config {
    node_count = 2 # High availability for reads
  }

  machine_config {
    cpu_count = 2
  }

  depends_on = [google_alloydb_instance.primary]
}