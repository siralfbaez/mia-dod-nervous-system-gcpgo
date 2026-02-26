package processor

import (
	"context"
	"testing"
	"cloud.google.com/go/pubsub"
)

func TestHandleMessage_Structural(t *testing.T) {
	// Verifies the signature of the handler
	ctx := context.Background()
	msg := &pubsub.Message{Data: []byte(`{"test": "data"}`)}

	// We just ensure it doesn't panic
	HandleMessage(ctx, msg)
}