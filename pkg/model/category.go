package model

type Category int

const (
	CategoryUnknown Category = iota
	CategoryBug
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
	CategoryBug:           "bug",
	CategorySecurity:       "security",
	CategoryPerformance:   "performance",
	CategoryStyle:         "style",
	CategoryComplexity:    "complexity",
	CategoryDuplication:   "duplication",
	CategoryArchitecture:  "architecture",
	CategoryTesting:       "testing",
	CategoryDocumentation: "documentation",
	CategoryBestPractice:  "best_practice",
}

func (c Category) String() string {
	if n, ok := categoryNames[c]; ok {
		return n
	}
	return "unknown"
}

func ParseCategory(s string) Category {
	for cat, name := range categoryNames {
		if name == s {
			return cat
		}
	}
	return CategoryUnknown
}
