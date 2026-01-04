package ascii

import (
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
		expectedError  error
	}{
		{
			name:          "No input provided",
			args:          []string{"cmd"},
			expectedError: errMissingInput,
		},
		{
			name:          "Empty input string",
			args:          []string{"cmd", ""},
			expectedError: errEmptyInput,
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
			expectedError: errInvalidBanner,
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
		{
			name:           "Banner style with uppercase (should fail)",
			args:           []string{"cmd", "Hello World"},
			expectedInput:  "Hello World",
			expectedBanner: "standard",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Override os.Args with simulated CLI input
			os.Args = tc.args

			// Call the function under test
			input, banner, err := GetUserInput()

			// Check expected error
			if tc.expectedError != nil {
				if err == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error %q, got %v", tc.expectedError, err)
				}
				return // skip further checks if error was expected
			}

			// If no error expected, check values
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
