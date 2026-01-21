package asciijustify

import (
	"strings"
)

// GetArtWidth returns the width of the widest line in ASCII art
func GetArtWidth(asciiLines []string) int {
	maxWidth := 0
	for _, line := range asciiLines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	return maxWidth
}

// GetArtHeight returns the number of lines in ASCII art
func GetArtHeight(asciiLines []string) int {
	return len(asciiLines)
}

// MeasureText measures dimensions of text as it would be rendered in ASCII art
// Returns: width, height
func MeasureText(input string, banner map[rune][]string) (int, int) {
	lines := strings.Split(input, "\n")
	maxWidth := 0
	totalHeight := 0

	for _, line := range lines {
		if line == "" {
			totalHeight += 1
			continue
		}

		// Calculate width of this line
		lineWidth := 0
		for _, ch := range line {
			if charArt, ok := banner[ch]; ok {
				lineWidth += GetArtLineWidth(charArt)
			}
		}

		if lineWidth > maxWidth {
			maxWidth = lineWidth
		}
		totalHeight += 8 // Each line of text is 8 rows in ASCII art
	}

	return maxWidth, totalHeight
}

// GetLineWidth calculates the width of a single ASCII art line
// This is a helper for measuring already-rendered ASCII art
func GetLineWidth(line string) int {
	return len(line)
}

// GetMaxLineWidth returns the width of the longest line in rendered ASCII art
func GetMaxLineWidth(asciiLines []string) int {
	return GetArtWidth(asciiLines)
}
