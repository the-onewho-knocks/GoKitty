package interfaces

import "context"

type AIService interface {
	GenerateSummary(ctx context.Context, data map[string]any) (string, error)
	GenerateRoast(ctx context.Context, data map[string]any) (string, error)
}
