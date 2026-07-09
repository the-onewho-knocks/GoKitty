package model

type Severity int

const (
	SeverityCritical Severity = iota + 1
	SeverityHigh
	SeverityMedium
	SeverityLow
	SeverityInfo
)

var severityNames = map[Severity]string{
	SeverityCritical: "critical",
	SeverityHigh:     "high",
	SeverityMedium:   "medium",
	SeverityLow:      "low",
	SeverityInfo:     "info",
}

func (s Severity) String() string {
	if name, ok := severityNames[s]; ok {
		return name
	}
	return "unknown"
}

func (s Severity) Int() int {
	return int(s)
}