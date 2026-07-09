package interfaces

import (
	"context"
	"gokitty/pkg/model"
	"time"
)

type ToolExecutor interface {
	Run(ctx context.Context, name string, args ...string) (*model.ToolResult, error)
	RunWithTimeout(ctx context.Context, timeout time.Duration, name string, args ...string) (*model.ToolResult, error)
}
