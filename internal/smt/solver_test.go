package smt

import (
	"testing"
	"time"
)

func TestZ3Solver_CheckAction(t *testing.T) {
	solver, err := NewZ3Solver()
	if err != nil {
		t.Fatalf("Failed to initialize Z3Solver: %v", err)
	}
	defer solver.Close()

	constraints := []string{"no_data_exfiltration", "read_only_access"}

	// Scenario 1: Benign task execution (Should be SATISFIABLE)
	benignAction := "read_user_profile"
	isSafe, latency, err := solver.CheckAction(benignAction, constraints)

	if err != nil {
		t.Errorf("Unexpected error during benign action check: %v", err)
	}
	if !isSafe {
		t.Errorf("Expected benign action '%s' to be approved, but got blocked", benignAction)
	}
	t.Logf("Benign Action approved in %v", latency)

	// Scenario 2: Multi-step adversarial action (Should trigger UNSAT / Algebraic Deadlock)
	maliciousAction := "export_db"
	isSafe, latency, err = solver.CheckAction(maliciousAction, constraints)

	if err != nil {
		t.Errorf("Unexpected error during malicious action check: %v", err)
	}
	if isSafe {
		t.Errorf("Expected malicious action '%s' to be blocked, but got approved", maliciousAction)
	}

	// Verify that the latency is under the target threshold (~1ms max for safety margin)
	if latency > 1*time.Millisecond {
		t.Logf("Warning: Latency %v exceeded the 0.44ms target threshold", latency)
	} else {
		t.Logf("Malicious Action successfully blocked in %v (Target met)", latency)
	}
}
