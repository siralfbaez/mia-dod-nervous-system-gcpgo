package handler

import (
	"net/http"
	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo/pkg/resilience"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	// Example of using the resilience package we built
	err := resilience.ExecuteWithRetry(func() error {
		// Logic to forward to Pub/Sub would go here
		return nil
	})
	if err != nil {
		http.Error(w, "Failed to process signal", 500)
	}
}