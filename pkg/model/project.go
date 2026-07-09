package model

import "time"

type Project struct {
	Name         string    `json:"name" yaml:"name"`
	Path         string    `json:"path" yaml:"path"`
	Language     Language  `json:"language" yaml:"language"`
	Framework    Framework `json:"framework" yaml:"framework"`
	Dependencies []string  `json:"dependencies" yaml:"dependencies"`
	FileCount    int       `json:"file_count" yaml:"file_count"`
	LOC          int       `json:"loc" yaml:"loc"`
	Config       map[string]any `json:"config" yaml:"config"`
	AnalyzedAt   time.Time `json:"analyzed_at" yaml:"analyzed_at"`
}