module "alloydb" {
  source      = "./modules/alloydb"
  project_id  = var.project_id
  region      = var.region
  vpc_id      = module.vpc.vpc_id
  db_password = var.db_password
}

module "cloud_run" {
  source          = "./modules/cloud-run"
  project_id      = var.project_id
  region          = var.region
  service_name    = "signal-validator"
  container_image = "gcr.io/${var.project_id}/validator:latest"
  vpc_connector_id = module.vpc.connector_id # You can add this to VPC module later
}

module "gke" {
  source     = "./modules/gke-cluster"
  project_id = var.project_id
  region     = var.region
  vpc_id     = module.vpc.vpc_id
  subnet_id  = module.vpc.subnet_id
}

module "iam" {
  source     = "./modules/iam-roles"
  project_id = var.project_id
}

module "kms" {
  source     = "./modules/kms-encryption"
  project_id = var.project_id
  region     = var.region
}

module "pubsub" {
  source     = "./modules/pubsub"
  project_id = var.project_id
}
module "vpc" {
  source = "./modules/vpc-network"
}

