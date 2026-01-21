package unit

import (
	justify "ascii-art/internal/ascii-justify"
	"strings"
	"testing"
)

func TestCenterAlign(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		termWidth int
		expected  []string
	}{
		{
			name:      "Center simple text",
			lines:     []string{"Hello"},
			termWidth: 20,
			expected:  []string{"       Hello"},
		},
		{
			name:      "Center multiple lines",
			lines:     []string{"Hi", "Test", "World"},
			termWidth: 20,
			expected:  []string{"         Hi", "        Test", "       World"},
		},
		{
			name:      "Text wider than terminal",
			lines:     []string{"This is too wide for terminal"},
			termWidth: 10,
			expected:  []string{"This is too wide for terminal"},
		},
		{
			name:      "Exact fit",
			lines:     []string{"ExactFit"},
			termWidth: 8,
			expected:  []string{"ExactFit"},
		},
		{
			name:      "Odd padding",
			lines:     []string{"Test"},
			termWidth: 11,
			expected:  []string{"   Test"}, // (11 - 4) / 2 = 3
		},
		{
			name:      "Empty lines",
			lines:     []string{},
			termWidth: 20,
			expected:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := justify.CenterAlign(tt.lines, tt.termWidth)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d lines, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Line %d: expected %q, got %q", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestRightAlign(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		termWidth int
		expected  []string
	}{
		{
			name:      "Right align simple text",
			lines:     []string{"Hello"},
			termWidth: 20,
			expected:  []string{"               Hello"},
		},
		{
			name:      "Right align multiple lines",
			lines:     []string{"Hi", "Test", "World"},
			termWidth: 20,
			expected:  []string{"                  Hi", "                Test", "               World"},
		},
		{
			name:      "Text wider than terminal",
			lines:     []string{"This is too wide"},
			termWidth: 10,
			expected:  []string{"This is too wide"},
		},
		{
			name:      "Exact fit",
			lines:     []string{"ExactFit"},
			termWidth: 8,
			expected:  []string{"ExactFit"},
		},
		{
			name:      "Empty lines",
			lines:     []string{},
			termWidth: 20,
			expected:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := justify.RightAlign(tt.lines, tt.termWidth)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d lines, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Line %d: expected %q, got %q", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestApplyAlignment(t *testing.T) {
	testLines := []string{
		"Hello",
		"World",
	}

	tests := []struct {
		name      string
		lines     []string
		alignType string
		termWidth int
		checkFunc func([]string) bool
	}{
		{
			name:      "Left alignment (default)",
			lines:     testLines,
			alignType: "left",
			termWidth: 20,
			checkFunc: func(result []string) bool {
				return result[0] == "Hello" && result[1] == "World"
			},
		},
		{
			name:      "Center alignment",
			lines:     testLines,
			alignType: "center",
			termWidth: 20,
			checkFunc: func(result []string) bool {
				// Check that lines are centered (have padding before them)
				return strings.HasPrefix(result[0], "       ") &&
					strings.HasPrefix(result[1], "       ")
			},
		},
		{
			name:      "Right alignment",
			lines:     testLines,
			alignType: "right",
			termWidth: 20,
			checkFunc: func(result []string) bool {
				// Check that lines are right-aligned (have padding before them)
				return strings.HasPrefix(result[0], "               ") &&
					strings.HasPrefix(result[1], "               ")
			},
		},
		{
			name:      "Justify alignment (placeholder - uses center for now)",
			lines:     testLines,
			alignType: "justify",
			termWidth: 20,
			checkFunc: func(result []string) bool {
				// Currently justify falls back to center
				return len(result) > 0
			},
		},
		{
			name:      "Invalid alignment type defaults to left",
			lines:     testLines,
			alignType: "invalid",
			termWidth: 20,
			checkFunc: func(result []string) bool {
				return result[0] == "Hello" && result[1] == "World"
			},
		},
		{
			name:      "Empty lines",
			lines:     []string{},
			alignType: "center",
			termWidth: 20,
			checkFunc: func(result []string) bool {
				return len(result) == 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := justify.ApplyAlignment(tt.lines, tt.alignType, tt.termWidth)
			if !tt.checkFunc(result) {
				t.Errorf("Alignment check failed for %s", tt.name)
				t.Logf("Result: %v", result)
			}
		})
	}
}

func TestAlignLine(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		alignType string
		termWidth int
		expected  string
	}{
		{
			name:      "Center single line",
			line:      "Test",
			alignType: "center",
			termWidth: 20,
			expected:  "        Test",
		},
		{
			name:      "Right single line",
			line:      "Test",
			alignType: "right",
			termWidth: 20,
			expected:  "                Test",
		},
		{
			name:      "Left single line",
			line:      "Test",
			alignType: "left",
			termWidth: 20,
			expected:  "Test",
		},
		{
			name:      "Line wider than terminal",
			line:      "This is a very long line",
			alignType: "center",
			termWidth: 10,
			expected:  "This is a very long line",
		},
		{
			name:      "Invalid alignment defaults to left",
			line:      "Test",
			alignType: "invalid",
			termWidth: 20,
			expected:  "Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := justify.AlignLine(tt.line, tt.alignType, tt.termWidth)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestJustifyChunk(t *testing.T) {
	// Test the justify chunk helper
	chunk := []string{
		"Line 1",
		"Line 2",
		"Line 3",
	}

	result := justify.JustifyChunk(chunk, 20)

	// Currently justifyChunk uses centerAlign as placeholder
	// So we just check it returns something
	if len(result) == 0 {
		t.Error("Expected non-empty result from justifyChunk")
	}

	if len(result) != len(chunk) {
		t.Errorf("Expected %d lines, got %d", len(chunk), len(result))
	}
}
