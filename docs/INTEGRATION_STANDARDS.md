# Integration Standards: The "Authority" Protocol

## 1. Schema Enforcement
* **Contract First:** All integration partners must align with the `v1-spec.yaml` (OpenAPI).
* **Validation:** The `contract-validator` service enforces strict typing. Any field mismatch results in an immediate `422 Unprocessable Entity` response.

## 2. Reliability Patterns (The "Immune System")
To ensure system stability, all upstream clients MUST implement:
* **Circuit Breaker:** Clients must stop sending requests for 60s if the Nervous System returns five consecutive `5xx` errors.
* **Idempotency:** Every signal must include a unique `X-Request-ID`. Retries of the same ID will not create duplicate entries in the "Metabolism" (AlloyDB).

## 3. Performance Targets
* **P99 Latency:** Gateway processing must complete within <200ms.
* **Timeout:** Connection timeout is hard-coded at 3.0 seconds.