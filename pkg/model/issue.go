package model

type Issue struct {
	Summary        string
	Findings       []Finding
	Category       Category
	Score          float64
	Recommendation string
}