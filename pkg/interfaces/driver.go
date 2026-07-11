package interfaces

import (
	"context"

	"github.com/the-onewho-knocks/gokitty/pkg/model"
)

type LanguageDriver interface {
	Detect(ctx context.Context, root string) (model.Language, error)
	Execute(ctx context.Context, root string, lang model.Language) ([]model.ToolResult, error)
	ParseOutput(result model.ToolResult) ([]model.Finding, error)
}
