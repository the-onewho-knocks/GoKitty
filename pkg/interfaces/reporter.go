package interfaces

import "github.com/the-onewho-knocks/gokitty/pkg/model"

type Reporter interface {
	Generate(report *model.Report) ([]byte, error)
	Format() string
}
