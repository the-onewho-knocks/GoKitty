package model

type Statistics struct {
	TotalFiles       int              `json:"total_files" yaml:"total_files"`
	TotalLOC         int              `json:"total_loc" yaml:"total_loc"`
	TotalFindings    int              `json:"total_findings" yaml:"total_findings"`
	BySeverity       map[string]int   `json:"by_severity" yaml:"by_severity"`
	ByCategory       map[string]int   `json:"by_category" yaml:"by_category"`
	ByLanguage       map[string]int   `json:"by_language" yaml:"by_language"`
	AvgComplexity    float64          `json:"avg_complexity" yaml:"avg_complexity"`
	MaxComplexity    int              `json:"max_complexity" yaml:"max_complexity"`
	AvgCohesion      float64          `json:"avg_cohesion" yaml:"avg_cohesion"`
	AvgCoupling      float64          `json:"avg_coupling" yaml:"avg_coupling"`
	CircularDeps     int              `json:"circular_deps" yaml:"circular_deps"`
	TestCoverage     float64          `json:"test_coverage" yaml:"test_coverage"`
	CommentRatio     float64          `json:"comment_ratio" yaml:"comment_ratio"`
	DuplicationRatio float64          `json:"duplication_ratio" yaml:"duplication_ratio"`
}
