package model

type Statistics struct {
	TotalFiles        int                `json:"total_files"`
	TotalLOC          int                `json:"total_loc"`
	TotalFindings     int                `json:"total_findings"`
	BySeverity         map[string]int    `json:"by_severity"`
	ByCategory       map[string]int    `json:"by_category"`
	AvgComplexity     float64           `json:"avg_complexity"`
	MaxComplexity     int               `json:"max_complexity"`
	AvgCohesion       float64           `json:"avg_cohesion"`
	AvgCoupling       float64           `json:"avg_coupling"`
	CircularDeps      int               `json:"circular_deps"`
	TestCoverage      float64           `json:"test_coverage"`
	CommentRatio      float64           `json:"comment_ratio"`
	DuplicationRatio  float64           `json:"duplication_ratio"`
}