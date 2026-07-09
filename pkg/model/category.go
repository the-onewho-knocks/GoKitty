package model

type Category int

const (
	CategoryBug Category = iota + 1
	CategorySecurity
	CategoryPerformance
	CategoryStyle
	CategoryComplexity
	CategoryDuplication
	CategoryArchitecture
	CategoryTesting
	CategoryDocumentation
	CategoryBestPractice
)

var categoryNames = map[Category]string{
	CategoryBug:            "bug",
	CategorySecurity:        "security",
	CategoryPerformance:    "performance",
	CategoryStyle:          "style",
	CategoryComplexity:     "complexity",
	CategoryDuplication:    "duplication",
	CategoryArchitecture:   "architecture",
	CategoryTesting:        "testing",
	CategoryDocumentation:  "documentation",
	CategoryBestPractice:   "best_practice",
}

func (c Category) String() string {
	if name, ok := categoryNames[c]; ok {
		return name
	}
	return "unknown"
}