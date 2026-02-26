module "vpc" {
  source = "./modules/vpc-network"
}

module "alloydb" {
  source      = "./modules/alloydb"
  project_id  = var.project_id
  region      = var.region
  vpc_id      = module.vpc.vpc_id
  db_password = var.db_password
}

module "gke" {
  source     = "./modules/gke-cluster"
  project_id = var.project_id
  region     = var.region
  vpc_id     = module.vpc.vpc_id
  subnet_id  = module.vpc.subnet_id
}