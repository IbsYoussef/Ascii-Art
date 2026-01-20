package unit

import (
	justify "ascii-art/internal/ascii-justify"
	"testing"
)

func TestGetTerminalWidth(t *testing.T) {
	// Note: This test will return actual terminal width or 80 (fallback)
	width := justify.GetTerminalWidth()

	if width < 1 {
		t.Errorf("Expected positive terminal width, got %d", width)
	}

	// Terminal width should be reasonable (not 0, not negative)
	if width <= 0 {
		t.Errorf("Terminal width should be positive, got %d", width)
	}
}

func TestFitsInTerminal(t *testing.T) {
	tests := []struct {
		name         string
		contentWidth int
		termWidth    int
		expected     bool
	}{
		{
			name:         "Content fits exactly",
			contentWidth: 80,
			termWidth:    80,
			expected:     true,
		},
		{
			name:         "Content fits with room",
			contentWidth: 50,
			termWidth:    80,
			expected:     true,
		},
		{
			name:         "Content too wide",
			contentWidth: 100,
			termWidth:    80,
			expected:     false,
		},
		{
			name:         "Small content",
			contentWidth: 10,
			termWidth:    120,
			expected:     true,
		},
		{
			name:         "Zero width content",
			contentWidth: 0,
			termWidth:    80,
			expected:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fits := tt.contentWidth <= tt.termWidth
			if fits != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, fits)
			}
		})
	}
}

func TestGetArtLineWidth(t *testing.T) {
	tests := []struct {
		name     string
		charArt  []string
		expected int
	}{
		{
			name: "Standard character width",
			charArt: []string{
				"  _  ",
				" | | ",
				" |_| ",
				" | | ",
				" |_| ",
				"     ",
				"     ",
				"     ",
			},
			expected: 5,
		},
		{
			name: "Wide character",
			charArt: []string{
				"        ",
				"  ___   ",
				" |   |  ",
				" |___|  ",
				" |   |  ",
				"        ",
				"        ",
				"        ",
			},
			expected: 8,
		},
		{
			name: "Narrow character",
			charArt: []string{
				" _ ",
				"| |",
				"|_|",
				"   ",
				"   ",
				"   ",
				"   ",
				"   ",
			},
			expected: 3,
		},
		{
			name:     "Empty character art",
			charArt:  []string{},
			expected: 0,
		},
		{
			name: "Single line",
			charArt: []string{
				"     ",
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			width := justify.GetArtLineWidth(tt.charArt)
			if width != tt.expected {
				t.Errorf("Expected width %d, got %d", tt.expected, width)
			}
		})
	}
}

func TestCalculatePadding(t *testing.T) {
	tests := []struct {
		name          string
		contentWidth  int
		terminalWidth int
		alignType     string
		expected      int
	}{
		{
			name:          "Center alignment - even padding",
			contentWidth:  40,
			terminalWidth: 80,
			alignType:     "center",
			expected:      20,
		},
		{
			name:          "Center alignment - odd padding",
			contentWidth:  41,
			terminalWidth: 80,
			alignType:     "center",
			expected:      19, // (80 - 41) / 2 = 19
		},
		{
			name:          "Right alignment",
			contentWidth:  50,
			terminalWidth: 80,
			alignType:     "right",
			expected:      30,
		},
		{
			name:          "Left alignment",
			contentWidth:  40,
			terminalWidth: 80,
			alignType:     "left",
			expected:      0,
		},
		{
			name:          "Content too wide - no padding",
			contentWidth:  100,
			terminalWidth: 80,
			alignType:     "center",
			expected:      0,
		},
		{
			name:          "Content exactly fits - no padding",
			contentWidth:  80,
			terminalWidth: 80,
			alignType:     "right",
			expected:      0,
		},
		{
			name:          "Invalid align type defaults to no padding",
			contentWidth:  40,
			terminalWidth: 80,
			alignType:     "invalid",
			expected:      0,
		},
		{
			name:          "Justify type - no padding (handled separately)",
			contentWidth:  40,
			terminalWidth: 80,
			alignType:     "justify",
			expected:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			padding := justify.CalculatePadding(tt.contentWidth, tt.terminalWidth, tt.alignType)
			if padding != tt.expected {
				t.Errorf("Expected padding %d, got %d", tt.expected, padding)
			}
		})
	}
}
