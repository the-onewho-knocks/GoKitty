package interfaces

import "gokitty/pkg/model"

type LanguageDriver interface {
	Detect(root string) (model.Language, error)
	Execute(root string, language model.Language) ([]model.ToolResult, error)
	ParseOutput(result model.ToolResult) ([]model.Finding, error)
}
