package model

import "time"

type Project struct {
	Name         string         `json:"name" yaml:"name"`
	Path         string         `json:"path" yaml:"path"`
	Languages    []Language     `json:"languages" yaml:"languages"`
	Framework    Framework      `json:"framework" yaml:"framework"`
	Dependencies []string       `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	FileCount    int            `json:"file_count" yaml:"file_count"`
	LOC          int            `json:"loc" yaml:"loc"`
	Config       map[string]any `json:"config,omitempty" yaml:"config,omitempty"`
	AnalyzedAt   time.Time      `json:"analyzed_at" yaml:"analyzed_at"`
}
