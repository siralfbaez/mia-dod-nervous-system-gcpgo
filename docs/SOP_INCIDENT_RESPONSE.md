# SOP: Regional GCP Outage & System Recovery

**Severity:** Level 1 (Critical)  
**Service:** MIA-DoD Nervous System  
**Objective:** Restore order signal flow and protect AlloyDB data integrity during a regional failure.

## 1. Detection & Identification
* **Trigger:** Global monitoring alerts (Cloud Monitoring) report >15% error rate on `signal-gateway`.
* **Verification:** Check the [Google Cloud Service Health Dashboard](https://status.cloud.google.com/).
* **Validation:** Confirm if the `resilience/breaker.go` is in an `OPEN` state across all nodes.

## 2. Immediate Response Actions

### Phase A: Traffic Shedding (The Reflex)
1. **Redirect Traffic:** Update Cloud DNS or Load Balancer to point to the secondary region (e.g., `us-east1`) if GKE Autopilot is multi-regional.
2. **Circuit Breaking:** Verify that the `signal-gateway` has tripped the circuit breaker to prevent "thundering herd" attempts against the struggling API.

### Phase B: Data Persistence (The Metabolism)
1. **AlloyDB Failover:** If the Primary Writer is unresponsive, promote the Read Pool in the secondary zone to Primary.
2. **Pub/Sub Backlog:** Ensure the `worker-pubsub` is NOT acknowledging messages if the database is unreachable. Let signals sit in the 7-day Pub/Sub retention buffer.

## 3. Recovery Procedures (Self-Healing)
* **AI Diagnosis:** Invoke the `ai-agent` via the secondary region's Vertex AI endpoint to analyze the "last known good" signals to detect data corruption during the crash.
* **Database Consistency:** Run the migration check script:
  ```bash
  cd database/scripts && ./setup_db.sh --verify-integrity
  

