package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo/pkg/resilience"
	"github.com/sony/gobreaker"
)

type Gateway struct {
	breaker *gobreaker.CircuitBreaker
}

func main() {
	// Initialize the "Immune System" (Circuit Breaker)
	cb := resilience.NewCircuitBreaker("byte-platform-api")

	gw := &Gateway{breaker: cb}

	http.HandleFunc("/v1/order", gw.HandleOrder)

	log.Println("⚡ Signal Gateway firing on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func (g *Gateway) HandleOrder(w http.ResponseWriter, r *http.Request) {
	// Staff Pattern: Use the breaker to protect the system
	result, err := g.breaker.Execute(func() (interface{}, error) {
		// This simulates a call to the "Byte" platform
		return CallByteAPI(r.Context())
	})

	if err != nil {
		log.Printf("🚨 Breaker tripped or Byte API failed: %v", err)
		http.Error(w, "Service Temporarily Unavailable (Resilience Mode)", http.StatusServiceUnavailable)
		return
	}

	fmt.Fprintf(w, "Signal Processed: %v", result)
}

func CallByteAPI(ctx context.Context) (string, error) {
	// Simulate integration latency
	time.Sleep(100 * time.Millisecond)
	return "Order_Success_200", nil
}