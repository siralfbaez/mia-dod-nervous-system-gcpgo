package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID")
	subID := "nervous-system-sub"

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	// Context for graceful shutdown
	cctx, cancel := context.WithCancel(ctx)

	// Listen for OS signals (SIGINT, SIGTERM)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		fmt.Println("\n🛑 Shutdown signal received. Finishing active tasks...")
		cancel()
	}()

	fmt.Printf("📮 Worker listening on subscription: %s\n", subID)

	// Receive messages until the context is cancelled
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		processOrder(msg)
	})

	if err != nil && err != context.Canceled {
		log.Fatalf("Receive error: %v", err)
	}
}

func processOrder(msg *pubsub.Message) {
	fmt.Printf("🍕 Processing Signal ID: %s | Data: %s\n", msg.ID, string(msg.Data))

	// Staff Pattern: Always acknowledge only after successful logic
	// In a real scenario, this is where AlloyDB persistence happens
	msg.Ack()
}