package model

type Severity int

const (
	SeverityUnknown Severity = iota
	SeverityInfo
	SeverityLow
	SeverityMedium
	SeverityHigh
	SeverityCritical
)

var severityNames = map[Severity]string{
	SeverityUnknown: "unknown",
	SeverityInfo:    "info",
	SeverityLow:     "low",
	SeverityMedium:  "medium",
	SeverityHigh:    "high",
	SeverityCritical: "critical",
}

func (s Severity) String() string {
	if n, ok := severityNames[s]; ok {
		return n
	}
	return "unknown"
}

func ParseSeverity(s string) Severity {
	for sev, name := range severityNames {
		if name == s {
			return sev
		}
	}
	return SeverityUnknown
}
