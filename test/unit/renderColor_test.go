package unit_test

import (
	color "ascii-art/internal/ascii-color"
	"testing"
)

func TestBuildColorMap(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		substring string
		expected  map[int]bool
	}{
		{
			name:      "No substring - color all",
			line:      "hello",
			substring: "",
			expected: map[int]bool{
				0: true, 1: true, 2: true, 3: true, 4: true,
			},
		},
		{
			name:      "Single occurrence",
			line:      "hello world",
			substring: "world",
			expected: map[int]bool{
				6: true, 7: true, 8: true, 9: true, 10: true,
			},
		},
		{
			name:      "Multiple occurrences",
			line:      "a king kitten have kit",
			substring: "kit",
			expected: map[int]bool{
				7: true, 8: true, 9: true, // "kitten"
				19: true, 20: true, 21: true, // "kit"
			},
		},
		{
			name:      "Case sensitive - no match",
			line:      "HELLO",
			substring: "hello",
			expected:  map[int]bool{},
		},
		{
			name:      "Partial overlap",
			line:      "aaa",
			substring: "aa",
			expected: map[int]bool{
				0: true, 1: true, // First "aa"
				2: true, // Second "aa" (overlaps)
			},
		},
		{
			name:      "No match",
			line:      "hello",
			substring: "xyz",
			expected:  map[int]bool{},
		},
		{
			name:      "Single character",
			line:      "RGB()",
			substring: "B",
			expected: map[int]bool{
				2: true, // 'B' at index 2
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := color.BuildColorMap(tt.line, tt.substring)

			// Check length
			if len(result) != len(tt.expected) {
				t.Errorf("ColorMap length = %d, want %d", len(result), len(tt.expected))
			}

			// Check each expected index
			for idx, shouldColor := range tt.expected {
				if result[idx] != shouldColor {
					t.Errorf("ColorMap[%d] = %v, want %v", idx, result[idx], shouldColor)
				}
			}

			// Check no unexpected indices
			for idx := range result {
				if _, exists := tt.expected[idx]; !exists {
					t.Errorf("Unexpected index %d in colorMap", idx)
				}
			}
		})
	}
}
