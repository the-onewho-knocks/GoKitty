package model

import (
	"time"
)

type Report struct {
	Project              Project            `json:"project" yaml:"project"`
	Scores               Score              `json:"scores" yaml:"scores"`
	Findings             []Finding          `json:"findings" yaml:"findings"`
	Issues               []Issue            `json:"issues" yaml:"issues"`
	Statistics           Statistics         `json:"statistics" yaml:"statistics"`
	PerLanguageFindings  map[string][]Finding `json:"per_language_findings,omitempty" yaml:"per_language_findings,omitempty"`
	ArchitectureAnalysis ArchitectureReport `json:"architecture_analysis" yaml:"architecture_analysis"`
	AISummary            string             `json:"ai_summary,omitempty" yaml:"ai_summary,omitempty"`
	AIRoast              string             `json:"ai_roast,omitempty" yaml:"ai_roast,omitempty"`
	GeneratedAt          time.Time          `json:"generated_at" yaml:"generated_at"`
	Duration             time.Duration      `json:"duration" yaml:"duration"`
	Metadata             map[string]string  `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

type ArchitectureReport struct {
	DependencyGraph *DependencyGraphResult `json:"dependency_graph,omitempty"`
	Coupling        *CouplingResult        `json:"coupling,omitempty"`
	Cohesion        *CohesionResult        `json:"cohesion,omitempty"`
	CircularDeps    []CircularDependency   `json:"circular_deps,omitempty"`
	Complexity      *ComplexityResult      `json:"complexity,omitempty"`
	Layering        *LayeringResult        `json:"layering,omitempty"`
	Modularity      *ModularityResult      `json:"modularity,omitempty"`
	Structure       *StructureResult       `json:"structure,omitempty"`
	CodeStats       *CodeStatistics        `json:"code_stats,omitempty"`
	TechDebt        *TechDebtResult        `json:"tech_debt,omitempty"`
}

type DependencyGraphResult struct {
	Nodes []string         `json:"nodes"`
	Edges []DependencyEdge `json:"edges"`
}

type DependencyEdge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type CouplingResult struct {
	Packages        []PackageCoupling `json:"packages"`
	AverageEfferent float64           `json:"average_efferent"`
	AverageAfferent float64           `json:"average_afferent"`
}

type PackageCoupling struct {
	Name        string  `json:"name"`
	Efferent    int     `json:"efferent"`
	Afferent    int     `json:"afferent"`
	Instability float64 `json:"instability"`
}

type CohesionResult struct {
	Packages    []PackageCohesion `json:"packages"`
	AverageLCOM float64           `json:"average_lcom"`
}

type PackageCohesion struct {
	Name string  `json:"name"`
	LCOM float64 `json:"lcom"`
}

type CircularDependency struct {
	Path     []string `json:"path"`
	Packages []string `json:"packages"`
}

type ComplexityResult struct {
	AverageCyclomatic float64           `json:"average_cyclomatic"`
	MaxCyclomatic     int               `json:"max_cyclomatic"`
	Files             []FileComplexity  `json:"files"`
}

type FileComplexity struct {
	File       string `json:"file"`
	Complexity int    `json:"complexity"`
}

type LayeringResult struct {
	Violations      []LayerViolation `json:"violations"`
	TotalViolations int              `json:"total_violations"`
}

type LayerViolation struct {
	Source      string `json:"source"`
	Target      string `json:"target"`
	Description string `json:"description"`
}

type ModularityResult struct {
	QScore               float64 `json:"q_score"`
	InterPackageCoupling float64 `json:"inter_package_coupling"`
	IntraPackageCoupling float64 `json:"intra_package_coupling"`
}

type StructureResult struct {
	DirectoryDepth  int            `json:"directory_depth"`
	PackageCount    int            `json:"package_count"`
	FileTypes       map[string]int `json:"file_types"`
}

type CodeStatistics struct {
	TotalFiles    int     `json:"total_files"`
	TotalLOC      int     `json:"total_loc"`
	TotalComments int     `json:"total_comments"`
	CommentRatio  float64 `json:"comment_ratio"`
	AvgFileSize   float64 `json:"avg_file_size"`
	MaxFileLength int     `json:"max_file_length"`
	AvgNestingDepth float64 `json:"avg_nesting_depth"`
}

type TechDebtResult struct {
	TotalDebt       float64  `json:"total_debt"`
	DebtRatio       float64  `json:"debt_ratio"`
	RemediationCost string   `json:"remediation_cost"`
	MainIssues      []string `json:"main_issues"`
}
