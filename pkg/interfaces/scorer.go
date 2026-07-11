package interfaces

import (
	"context"

	"github.com/the-onewho-knocks/gokitty/pkg/model"
)

type Scorer interface {
	Assess(ctx context.Context, project model.Project, findings []model.Finding, arch model.ArchitectureReport) (*model.Score, error)
}
