package epca

import (
	"errors"
	"strings"
)

// Parser is responsible for translating natural language intents from AI agents
// into First-Order Logic constraints suitable for the SMT solver.
type Parser struct {
	// In a production environment, this would hold lightweight deterministic syntax trees
	// rather than empirical LLM adjudicators to prevent reasoning-loop DoS attacks.
}

// NewParser initializes a new ePCA constraint parser.
func NewParser() *Parser {
	return &Parser{}
}

// ParseIntent extracts mathematically verifiable constraints from an agent's raw action string.
func (p *Parser) ParseIntent(rawAction string) (string, []string, error) {
	if strings.TrimSpace(rawAction) == "" {
		return "", nil, errors.New("empty action payload")
	}

	var constraints []string
	normalizedAction := strings.ToLower(rawAction)

	// Simulate the translation of semantic intent into first-order logic axioms.
	// This mapping ensures that high-privilege autonomous operations are bounded.
	if strings.Contains(normalizedAction, "database") || strings.Contains(normalizedAction, "export") {
		constraints = append(constraints, "no_data_exfiltration")
	}

	if strings.Contains(normalizedAction, "delete") || strings.Contains(normalizedAction, "remove") {
		constraints = append(constraints, "read_only_access")
	}

	// Fallback safe constraint for unknown or general actions
	if len(constraints) == 0 {
		constraints = append(constraints, "default_safe_execution")
	}

	actionTarget := extractActionTarget(normalizedAction)

	return actionTarget, constraints, nil
}

// extractActionTarget maps raw text to a definitive action signature for the solver.
func extractActionTarget(action string) string {
	if strings.Contains(action, "export") {
		return "export_db"
	}
	if strings.Contains(action, "read") || strings.Contains(action, "profile") {
		return "read_user_profile"
	}
	return "unknown_action"
}
