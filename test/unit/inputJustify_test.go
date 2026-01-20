package unit

import (
	justify "ascii-art/internal/ascii-justify"
	"strings"
	"testing"
)

func TestParseAlignFlag(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedAlign string
		expectedArgs  []string
		expectError   bool
		errorContains string
	}{
		{
			name:          "Valid center alignment",
			args:          []string{"--align=center", "Hello", "standard"},
			expectedAlign: "center",
			expectedArgs:  []string{"Hello", "standard"},
			expectError:   false,
		},
		{
			name:          "Valid left alignment",
			args:          []string{"--align=left", "World", "shadow"},
			expectedAlign: "left",
			expectedArgs:  []string{"World", "shadow"},
			expectError:   false,
		},
		{
			name:          "Valid right alignment",
			args:          []string{"--align=right", "Test", "thinkertoy"},
			expectedAlign: "right",
			expectedArgs:  []string{"Test", "thinkertoy"},
			expectError:   false,
		},
		{
			name:          "Valid justify alignment",
			args:          []string{"--align=justify", "Hello World", "standard"},
			expectedAlign: "justify",
			expectedArgs:  []string{"Hello World", "standard"},
			expectError:   false,
		},
		{
			name:          "No align flag - defaults to left",
			args:          []string{"Hello", "standard"},
			expectedAlign: "left",
			expectedArgs:  []string{"Hello", "standard"},
			expectError:   false,
		},
		{
			name:          "Align flag with other flags",
			args:          []string{"--color=red", "--align=center", "Text", "standard"},
			expectedAlign: "center",
			expectedArgs:  []string{"--color=red", "Text", "standard"},
			expectError:   false,
		},
		{
			name:          "Invalid alignment type",
			args:          []string{"--align=invalid", "Hello"},
			expectedAlign: "",
			expectedArgs:  nil,
			expectError:   true,
			errorContains: "invalid alignment type",
		},
		{
			name:          "Malformed flag - missing =",
			args:          []string{"--align", "center", "Hello"},
			expectedAlign: "",
			expectedArgs:  nil,
			expectError:   true,
			errorContains: "invalid --align flag format",
		},
		{
			name:          "Empty alignment type",
			args:          []string{"--align=", "Hello", "standard"},
			expectedAlign: "",
			expectedArgs:  nil,
			expectError:   true,
			errorContains: "invalid --align flag format",
		},
		{
			name:          "Multiple flags with align at end",
			args:          []string{"--color=blue", "--output=file.txt", "--align=right", "Test"},
			expectedAlign: "right",
			expectedArgs:  []string{"--color=blue", "--output=file.txt", "Test"},
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alignType, remainingArgs, err := justify.ParseAlignFlag(tt.args)

			// Check error expectation
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error containing %q, got %q", tt.errorContains, err.Error())
				}
				return
			}

			// Check no error when not expected
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check alignment type
			if alignType != tt.expectedAlign {
				t.Errorf("Expected alignType %q, got %q", tt.expectedAlign, alignType)
			}

			// Check remaining args
			if len(remainingArgs) != len(tt.expectedArgs) {
				t.Errorf("Expected %d remaining args, got %d", len(tt.expectedArgs), len(remainingArgs))
				return
			}

			for i, arg := range remainingArgs {
				if arg != tt.expectedArgs[i] {
					t.Errorf("Arg %d: expected %q, got %q", i, tt.expectedArgs[i], arg)
				}
			}
		})
	}
}

func TestHasAlignFlag(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected bool
	}{
		{
			name:     "Has align flag with center",
			args:     []string{"--align=center", "Hello"},
			expected: true,
		},
		{
			name:     "Has align flag with right",
			args:     []string{"--align=right", "Test"},
			expected: true,
		},
		{
			name:     "Has malformed align flag",
			args:     []string{"--align", "center"},
			expected: true,
		},
		{
			name:     "No align flag",
			args:     []string{"Hello", "standard"},
			expected: false,
		},
		{
			name:     "Has other flags but not align",
			args:     []string{"--color=red", "Hello"},
			expected: false,
		},
		{
			name:     "Empty args",
			args:     []string{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := justify.HasAlignFlag(tt.args)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestValidateAlignType(t *testing.T) {
	tests := []struct {
		name      string
		alignType string
		expected  bool
	}{
		{"Valid left", "left", true},
		{"Valid center", "center", true},
		{"Valid right", "right", true},
		{"Valid justify", "justify", true},
		{"Invalid type", "invalid", false},
		{"Empty string", "", false},
		{"Mixed case", "Center", false}, // Case-sensitive
		{"With spaces", " center ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := justify.ValidateAlignType(tt.alignType)
			if result != tt.expected {
				t.Errorf("Expected %v for %q, got %v", tt.expected, tt.alignType, result)
			}
		})
	}
}
