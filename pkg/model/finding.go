package model

type ToolSource struct {
	Name    string
	Version string
}

type Location struct {
	File   string
	Line   int
	Column int
}

type Finding struct {
	ID       string
	Tool     ToolSource
	RuleID   string
	Message  string
	Severity Severity
	Category Category
	Location Location
	Snippet  string
	Context  map[string]string
	Language Language
	RawOutput string
}