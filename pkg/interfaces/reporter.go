package interfaces

import "gokitty/pkg/model"

type Reporter interface {
	Generate(report *model.Report) ([]byte, error)
}
