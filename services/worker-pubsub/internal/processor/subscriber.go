package processor

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

// HandleMessage encapsulates the processing logic for a single signal
func HandleMessage(ctx context.Context, msg *pubsub.Message) {
	log.Printf("📥 Received signal for processing: %s", msg.ID)

	// In a full implementation, this is where we would:
	// 1. Unmarshal JSON
	// 2. Call the Validator
	// 3. Persist to AlloyDB

	fmt.Printf("⚙️ Processing payload: %s\n", string(msg.Data))

	// Staff Pattern: Acknowledge only on successful handoff
	msg.Ack()
}