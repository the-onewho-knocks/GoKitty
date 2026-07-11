package model

type Issue struct {
	Summary        string    `json:"summary" yaml:"summary"`
	Findings       []Finding `json:"findings" yaml:"findings"`
	Category       Category  `json:"category" yaml:"category"`
	Score          float64   `json:"score" yaml:"score"`
	Recommendation string    `json:"recommendation,omitempty" yaml:"recommendation,omitempty"`
}
