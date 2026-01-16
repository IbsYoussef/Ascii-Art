package ascii

import (
	"fmt"
	"strings"
)

// RenderAsciiWithColor renders ASCII art with color support
func RenderAsciiWithColor(input string, banner map[rune][]string, colorConfig ColorConfig) {
	// Parse the color to get ANSI code
	ansiCode, err := ParseColor(colorConfig.Color)
	if err != nil {
		// This shouldn't happen as we validated earlier, but just in case
		fmt.Println("Error:", err)
		return
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// If line is empty, just print newline
		if line == "" {
			fmt.Println()
			continue
		}

		// Determine which characters to color
		colorMap := BuildColorMap(line, colorConfig.Substring)

		// Build and print each of the 8 ASCII lines
		for row := 0; row < 8; row++ {
			var outputLine string

			for charIndex, ch := range line {
				if asciiChar, ok := banner[ch]; ok {
					charArt := asciiChar[row]

					// Apply color if this character should be colored
					if colorMap[charIndex] {
						charArt = ansiCode + charArt + ResetColor()
					}

					outputLine += charArt
				} else {
					// Character not in banner, use placeholder
					outputLine += "        " // 8 spaces
				}
			}

			fmt.Println(outputLine)
		}
	}
}

// buildColorMap determines which character indices should be colored
// Returns a map of character index -> should color (true/false)
func BuildColorMap(line string, substring string) map[int]bool {
	colorMap := make(map[int]bool)

	// If no substring specified, color everything
	if substring == "" {
		for i := range line {
			colorMap[i] = true
		}
		return colorMap
	}

	// Find all occurrences of substring (case-sensitive)
	for i := 0; i <= len(line)-len(substring); i++ {
		// Check if substring matches at position i
		if line[i:i+len(substring)] == substring {
			// Mark all characters in this occurrence as "to be colored"
			for j := i; j < i+len(substring); j++ {
				colorMap[j] = true
			}
		}
	}

	return colorMap
}
