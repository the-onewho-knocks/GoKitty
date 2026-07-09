package interfaces

import "gokitty/pkg/model"

type Scorer interface {
	Assess(project model.Project, findings []model.Finding, arch model.ArchitectureReport, config map[string]any) (*model.Score, error)
}
