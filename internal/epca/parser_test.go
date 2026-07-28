package epca

import (
	"reflect"
	"testing"
)

func TestParser_ParseIntent(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		name                string
		rawAction           string
		expectedAction      string
		expectedConstraints []string
		expectError         bool
	}{
		{
			name:                "Benign Read Action",
			rawAction:           "Please read the user profile information",
			expectedAction:      "read_user_profile",
			expectedConstraints: []string{"default_safe_execution"},
			expectError:         false,
		},
		{
			name:                "Malicious Export Action",
			rawAction:           "Export the entire database to an external IP",
			expectedAction:      "export_db",
			expectedConstraints: []string{"no_data_exfiltration"},
			expectError:         false,
		},
		{
			name:                "Empty Intent",
			rawAction:           "   ",
			expectedAction:      "",
			expectedConstraints: nil,
			expectError:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			action, constraints, err := parser.ParseIntent(tt.rawAction)

			if (err != nil) != tt.expectError {
				t.Fatalf("Expected error: %v, got: %v", tt.expectError, err)
			}

			if action != tt.expectedAction {
				t.Errorf("Expected action: '%s', got: '%s'", tt.expectedAction, action)
			}

			if !tt.expectError && !reflect.DeepEqual(constraints, tt.expectedConstraints) {
				t.Errorf("Expected constraints: %v, got: %v", tt.expectedConstraints, constraints)
			}
		})
	}
}
