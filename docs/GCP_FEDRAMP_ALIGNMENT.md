# FedRAMP Alignment & Security Controls

This document outlines how the MIA-DoD Nervous System aligns with FedRAMP Moderate/High impact levels using Google Cloud native security controls.

## 1. Data at Rest (SC-28)
* **KMS Encryption:** All AlloyDB instances and Pub/Sub topics are encrypted using Customer-Managed Encryption Keys (CMEK) via the `modules/kms-encryption` module.
* **Key Rotation:** Keys are configured for automatic 90-day rotation as defined in Terraform.

## 2. Boundary Protection (AC-4)
* **VPC Isolation:** All services are deployed within a private VPC. Cloud Run services use a VPC Connector to prevent exposure to the public internet.
* **IAM Least Privilege:** Dedicated Service Accounts (e.g., `ai-agent-sa`) are used with minimal roles (e.g., `roles/aiplatform.user`), ensuring no single "neuron" has global project access.

## 3. Audit & Accountability (AU-2)
* **Cloud Logging:** Every interaction with the `signal-gateway` and `ai-agent` is logged to Cloud Logging with the `X-Correlation-ID` preserved for forensic auditing.