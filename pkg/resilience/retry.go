package resilience

import (
	"time"
	"github.com/avast/retry-go"
)

func ExecuteWithRetry(fn func() error) error {
	return retry.Do(
		fn,
		retry.Attempts(3),
		retry.DelayType(retry.BackOffDelay),
		retry.Delay(1 * time.Second),
	)
}