package unit_test

import (
	"ascii-art/internal/ascii"
	"os"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	testCases := []struct {
		name           string
		args           []string
		expectedInput  string
		expectedBanner string
		expectedError  bool
		errorContains  string
	}{
		{
			name:          "No input provided",
			args:          []string{"cmd"},
			expectedError: true,
			errorContains: "missing input",
		},
		{
			name:          "Empty input string",
			args:          []string{"cmd", ""},
			expectedError: true,
			errorContains: "cannot be empty",
		},
		{
			name:           "Valid input with default banner",
			args:           []string{"cmd", "Hello"},
			expectedInput:  "Hello",
			expectedBanner: "standard",
		},
		{
			name:           "Valid input with shadow banner",
			args:           []string{"cmd", "Hi", "shadow"},
			expectedInput:  "Hi",
			expectedBanner: "shadow",
		},
		{
			name:          "Invalid banner style",
			args:          []string{"cmd", "Hey", "unknown"},
			expectedError: true,
			errorContains: "invalid banner",
		},
		{
			name:           "Escaped newline in input",
			args:           []string{"cmd", "Hello\\nWorld"},
			expectedInput:  "Hello\nWorld",
			expectedBanner: "standard",
		},
		{
			name:           "Input with leading and trailing spaces",
			args:           []string{"cmd", "   Hello   "},
			expectedInput:  "Hello",
			expectedBanner: "standard",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			input, banner, err := ascii.GetUserInput()

			if tc.expectedError {
				if err == nil {
					t.Errorf("Expected error containing %q, got nil", tc.errorContains)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if input != tc.expectedInput {
				t.Errorf("Expected input %q, got %q", tc.expectedInput, input)
			}
			if banner != tc.expectedBanner {
				t.Errorf("Expected banner %q, got %q", tc.expectedBanner, banner)
			}
		})
	}
}
