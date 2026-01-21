package unit

import (
	justify "ascii-art/internal/ascii-justify"
	"testing"
)

func TestGetArtWidth(t *testing.T) {
	tests := []struct {
		name       string
		asciiLines []string
		expected   int
	}{
		{
			name: "Single line with consistent width",
			asciiLines: []string{
				"     _    ",
				"    / \\   ",
				"   / _ \\  ",
				"  / ___ \\ ",
				" /_/   \\_\\",
			},
			expected: 10,
		},
		{
			name: "Multiple lines with varying widths",
			asciiLines: []string{
				"Hello",
				"World!!",
				"Hi",
			},
			expected: 7,
		},
		{
			name:       "Empty lines",
			asciiLines: []string{},
			expected:   0,
		},
		{
			name: "Lines with spaces",
			asciiLines: []string{
				"   Test   ",
				" Short ",
				"    ",
			},
			expected: 10,
		},
		{
			name: "All same width",
			asciiLines: []string{
				"AAAA",
				"BBBB",
				"CCCC",
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			width := justify.GetArtWidth(tt.asciiLines)
			if width != tt.expected {
				t.Errorf("Expected width %d, got %d", tt.expected, width)
			}
		})
	}
}

func TestGetArtHeight(t *testing.T) {
	tests := []struct {
		name       string
		asciiLines []string
		expected   int
	}{
		{
			name:       "8 lines (standard ASCII art character)",
			asciiLines: make([]string, 8),
			expected:   8,
		},
		{
			name:       "16 lines (two characters)",
			asciiLines: make([]string, 16),
			expected:   16,
		},
		{
			name:       "Empty",
			asciiLines: []string{},
			expected:   0,
		},
		{
			name:       "Single line",
			asciiLines: []string{"test"},
			expected:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			height := justify.GetArtHeight(tt.asciiLines)
			if height != tt.expected {
				t.Errorf("Expected height %d, got %d", tt.expected, height)
			}
		})
	}
}

func TestGetLineWidth(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{
			name:     "Regular line",
			line:     "Hello World",
			expected: 11,
		},
		{
			name:     "Line with spaces",
			line:     "   Test   ",
			expected: 10,
		},
		{
			name:     "Empty line",
			line:     "",
			expected: 0,
		},
		{
			name:     "ASCII art line",
			line:     " _   _      _ _       ",
			expected: 22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			width := justify.GetLineWidth(tt.line)
			if width != tt.expected {
				t.Errorf("Expected width %d, got %d", tt.expected, width)
			}
		})
	}
}

func TestGetMaxLineWidth(t *testing.T) {
	tests := []struct {
		name       string
		asciiLines []string
		expected   int
	}{
		{
			name: "Varying widths",
			asciiLines: []string{
				"Short",
				"This is longer",
				"Mid",
			},
			expected: 14,
		},
		{
			name: "All same width",
			asciiLines: []string{
				"Same",
				"Same",
				"Same",
			},
			expected: 4,
		},
		{
			name:       "Empty",
			asciiLines: []string{},
			expected:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			width := justify.GetMaxLineWidth(tt.asciiLines)
			if width != tt.expected {
				t.Errorf("Expected width %d, got %d", tt.expected, width)
			}
		})
	}
}
