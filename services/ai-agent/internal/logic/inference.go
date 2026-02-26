package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/genai"
)

type AnomalyResult struct {
	IsAnomaly bool   `json:"is_anomaly"`
	Reason    string `json:"reason"`
	Severity  string `json:"severity"`
}

func AnalyzeOrder(ctx context.Context, orderJSON string) (*AnomalyResult, error) {
	projectID := os.Getenv("GCP_PROJECT_ID")
	location := os.Getenv("GCP_REGION")

	// 1. Initialize the Vertex AI Client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  projectID,
		Location: location,
		Backend:  genai.BackendVertexAI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create vertex client: %w", err)
	}

	// 2. Load System Instructions
	instruction, _ := os.ReadFile("../../../agent/prompts/integration_healing_system.txt")

	// 3. Call Gemini 2.0 Flash (Fast & Cost-Effective)
	resp, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash",
		genai.Text(string(instruction)),
		genai.Text("Analyze this order: "+orderJSON),
		nil)

	if err != nil {
		return nil, fmt.Errorf("gemini inference failed: %w", err)
	}

	// 4. Parse JSON Output
	var result AnomalyResult
	if err := json.Unmarshal([]byte(resp.Candidates[0].Content.Parts[0].Text), &result); err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	return &result, nil
}