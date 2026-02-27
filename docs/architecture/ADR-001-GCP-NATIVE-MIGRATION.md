# ADR 001: Standardizing on GCP-Native Managed Services

**Status:** Proposed  
**Author:** Alf Baez  
**Deciders:** DevX Team, Staff Architect

## Context
The existing "nervous system" integration requires high availability and minimal operational overhead to support Pizza Hut's global ordering scale. Brittle manual node management is a primary source of downtime.

## Decision
We will utilize **GKE Autopilot** for compute and **AlloyDB** for state management.

## Rationale
1. **GKE Autopilot:** Moves the "responsibility line" toward Google. We no longer manage node pools, security patches, or bin-packing. This allows the DevX team to focus on *Integration Standards* rather than *Infrastructure Maintenance*.
2. **AlloyDB:** Provides 4x the throughput of standard PostgreSQL. The inclusion of **Read Pools** ensures that heavy reporting/AI analysis does not compete for resources with critical order-taking transactions.

## Consequences
- **Positive:** Faster time-to-market; built-in security hardening; automatic scaling.
- **Negative:** Higher cost floor compared to raw VMs; vendor lock-in to GCP.


## Resilience & Durability Strategy (Updated)

### 1 Pattern: Circuit Breaking (`pkg/resilience`)
We have implemented a custom Circuit Breaker in the Go transport layer. This prevents cascading failures by "tripping" the connection to downstream services when error thresholds are met, protecting both the Gateway and the Byte platform from exhaustion.

### 2 Pattern: Durable Messaging ("The Metabolism")
To ensure zero data loss during AlloyDB maintenance or regional brownouts, we utilize Pub/Sub's 7-day message retention. This allows the "Metabolism" (Worker layer) to pause during outages and "replay" signals once the state store is healthy, ensuring eventual consistency.

### 3 Pattern: Regional failover
Current architecture is Regional (us-central1), but the use of GKE Autopilot and Terraform modules allows for a "Warm Standby" deployment in us-east1 with minimal configuration changes, targeting Tier-0 availability.