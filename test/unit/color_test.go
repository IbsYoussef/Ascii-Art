package unit_test

import (
	color "ascii-art/internal/ascii-color"
	"testing"
)

func TestParseColor(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		// Named colors
		{"Named red", "red", false},
		{"Named blue", "blue", false},
		{"Named orange", "orange", false},
		{"Named uppercase", "RED", false},

		// Hex colors
		{"Hex full", "#FF0000", false},
		{"Hex short", "#F00", false},
		{"Hex lowercase", "#ff0000", false},

		// RGB colors
		{"RGB valid", "rgb(255, 0, 0)", false},
		{"RGB no spaces", "rgb(255,0,0)", false},

		// HSL colors
		{"HSL valid", "hsl(0, 100%, 50%)", false},
		{"HSL green", "hsl(120, 100%, 50%)", false},

		// Invalid
		{"Invalid name", "purple", true},
		{"Invalid hex", "#GGGGGG", true},
		{"Invalid RGB", "rgb(300, 0, 0)", true},
		{"Invalid HSL", "hsl(400, 100%, 50%)", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := color.ParseColor(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("ParseColor(%q) error = %v, wantError %v", tt.input, err, tt.wantError)
			}
		})
	}
}
