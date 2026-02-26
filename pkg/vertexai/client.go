package vertexai

import (
	"context"
	"fmt"
	"google.golang.org/genai"
)

type Client struct {
	genaiClient *genai.Client
}

// NewClient creates a reusable Vertex AI wrapper
func NewClient(ctx context.Context, projectID, location string) (*Client, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  projectID,
		Location: location,
		Backend:  genai.BackendVertexAI,
	})
	if err != nil {
		return nil, fmt.Errorf("vertexai client init failed: %w", err)
	}
	return &Client{genaiClient: client}, nil
}