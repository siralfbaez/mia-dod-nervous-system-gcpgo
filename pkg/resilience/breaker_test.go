package resilience

import (
	"errors"
	"testing"
)

func TestCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker("test-breaker")

	_, err := cb.Execute(func() (interface{}, error) {
		return nil, errors.New("simulated failure")
	})

	if err == nil {
		t.Error("Expected error from breaker, got nil")
	}
}