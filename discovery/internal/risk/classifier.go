// Package risk provides initial risk classification for discovered assets.
package risk

// Level represents a risk severity.
type Level int

const (
	LevelLow      Level = 1
	LevelMedium   Level = 2
	LevelHigh     Level = 3
	LevelCritical Level = 4
)

func (l Level) String() string {
	m := map[Level]string{LevelLow: "low", LevelMedium: "medium", LevelHigh: "high", LevelCritical: "critical"}
	if s, ok := m[l]; ok {
		return s
	}
	return "unknown"
}

// Factor is a single risk factor.
type Factor struct {
	Name   string  `json:"name"`
	Desc   string  `json:"description"`
	Weight float64 `json:"weight"`
	Score  float64 `json:"score"`
}

// Classification holds a risk assessment.
type Classification struct {
	Level   Level    `json:"level"`
	Score   float64  `json:"score"`
	Factors []Factor `json:"factors"`
}

type Classifier struct{}

func NewClassifier() *Classifier { return &Classifier{} }

// Classify returns an initial risk classification.
func (c *Classifier) Classify(tools []string, hasAuth bool, isPublic bool) *Classification {
	score := 0.0
	var factors []Factor

	if !hasAuth {
		factors = append(factors, Factor{"no_auth", "No authentication configured", 0.4, 100})
		score += 40
	}
	if isPublic {
		factors = append(factors, Factor{"public", "Publicly accessible", 0.3, 100})
		score += 30
	}
	sensitive := countSensitive(tools)
	if sensitive > 0 && len(tools) > 0 {
		ts := float64(sensitive) / float64(len(tools)) * 100
		factors = append(factors, Factor{"sensitive_tools", "Access to sensitive MCP tools", 0.3, ts})
		score += 0.3 * ts
	}

	level := LevelLow
	switch {
	case score >= 75:
		level = LevelCritical
	case score >= 50:
		level = LevelHigh
	case score >= 25:
		level = LevelMedium
	}
	return &Classification{Level: level, Score: score, Factors: factors}
}

func countSensitive(tools []string) int {
	s := map[string]bool{"execute_command": true, "write_file": true, "database_query": true, "send_email": true, "delete_resource": true, "deploy": true}
	n := 0
	for _, t := range tools {
		if s[t] { n++ }
	}
	return n
}
