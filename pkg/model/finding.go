package model

type ToolSource struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

type Location struct {
	File   string `json:"file" yaml:"file"`
	Line   int    `json:"line" yaml:"line"`
	Column int    `json:"column,omitempty" yaml:"column,omitempty"`
}

type Finding struct {
	ID        string            `json:"id" yaml:"id"`
	Tool      ToolSource        `json:"tool" yaml:"tool"`
	RuleID    string            `json:"rule_id" yaml:"rule_id"`
	Message   string            `json:"message" yaml:"message"`
	Severity  Severity          `json:"severity" yaml:"severity"`
	Category  Category          `json:"category" yaml:"category"`
	Location  Location          `json:"location" yaml:"location"`
	Snippet   string            `json:"snippet,omitempty" yaml:"snippet,omitempty"`
	Context   map[string]string `json:"context,omitempty" yaml:"context,omitempty"`
	Language  Language          `json:"language" yaml:"language"`
	RawOutput string            `json:"-" yaml:"-"`
}
