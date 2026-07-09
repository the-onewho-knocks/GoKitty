package model

import "time"

type ToolResult struct {
	Tool     ToolInfo      `json:"tool"`
	Success  bool          `json:"success"`
	Output   string        `json:"output"`
	Duration time.Duration `json:"duration"`
	Error    string        `json:"error,omitempty"`
}

type ToolInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Command string `json:"command"`
	Args    []string `json:"args"`
}