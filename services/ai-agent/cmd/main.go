package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/siralfbaez/mia-dod-nervous-system-gcpgo/services/ai-agent/internal/logic"
	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID")
	subID := "nervous-system-ai-sub" // Dedicated AI observation stream

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	// Graceful Shutdown Setup
	cctx, cancel := context.WithCancel(ctx)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		fmt.Println("\n🤖 AI Agent shutting down gracefully...")
		cancel()
	}()

	fmt.Printf("🧠 AI Agent watching signals on: %s\n", subID)

	// The Observation Loop
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("🔍 Analyzing Signal: %s\n", msg.ID)

		// 1. Pass the signal to Gemini for Anomaly Detection
		result, err := logic.AnalyzeOrder(ctx, string(msg.Data))
		if err != nil {
			log.Printf("⚠️ AI Inference Error: %v", err)
			msg.Nack() // Try again later if AI is down
			return
		}

		// 2. Act on the Result (The "Staff Architect" Flex)
		if result.IsAnomaly {
			log.Printf("🚨 ANOMALY DETECTED [%s]: %s", result.Severity, result.Reason)
			// In a real system, here we would trigger an 'Incident' or 'Mute' the vendor
		} else {
			log.Printf("✅ Signal Healthy: No anomaly detected.")
		}

		msg.Ack()
	})

	if err != nil && err != context.Canceled {
		log.Fatalf("AI Subscription error: %v", err)
	}
}