package model

type Score struct {
	Overall         float64 `json:"overall" yaml:"overall"`
	Architecture    float64 `json:"architecture" yaml:"architecture"`
	Maintainability float64 `json:"maintainability" yaml:"maintainability"`
	Security        float64 `json:"security" yaml:"security"`
	Testing         float64 `json:"testing" yaml:"testing"`
	Performance     float64 `json:"performance" yaml:"performance"`
	Documentation   float64 `json:"documentation" yaml:"documentation"`
	Dependency      float64 `json:"dependency" yaml:"dependency"`
}