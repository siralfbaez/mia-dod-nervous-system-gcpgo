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