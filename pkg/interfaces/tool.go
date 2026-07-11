package interfaces

import (
	"context"
	"time"

	"github.com/the-onewho-knocks/gokitty/pkg/model"
)

type ToolExecutor interface {
	Run(ctx context.Context, name string, args ...string) (*model.ToolResult, error)
	RunWithTimeout(ctx context.Context, timeout time.Duration, name string, args ...string) (*model.ToolResult, error)
}
