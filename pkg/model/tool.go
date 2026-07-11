package model

import "time"

type ToolResult struct {
	Tool     ToolInfo      `json:"tool" yaml:"tool"`
	Command  string        `json:"command" yaml:"command"`
	Success  bool          `json:"success" yaml:"success"`
	Output   string        `json:"output" yaml:"output"`
	Duration time.Duration `json:"duration" yaml:"duration"`
	Error    string        `json:"error,omitempty" yaml:"error,omitempty"`
}

type ToolInfo struct {
	Name    string   `json:"name" yaml:"name"`
	Version string   `json:"version,omitempty" yaml:"version,omitempty"`
	Command string   `json:"command" yaml:"command"`
	Args    []string `json:"args,omitempty" yaml:"args,omitempty"`
}
