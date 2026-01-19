package asciireverse

import (
	"strings"
)

// ParseAsciiArt parses ASCII art content into chunks of 8 lines
// Each chunk represents one line of text in the original message
// Returns a slice of chunks, where each chunk is 8 lines of ASCII art
func ParseAsciiArt(content string) ([][]string, error) {
	// Split content into lines
	lines := strings.Split(content, "\n")

	// Remove trailing empty lines
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var chunks [][]string

	// Process lines in groups of 8 (each character is 8 lines tall)
	for i := 0; i < len(lines); i += 8 {
		// Get 8 lines for this chunk
		chunk := make([]string, 8)
		for j := 0; j < 8; j++ {
			if i+j < len(lines) {
				chunk[j] = lines[i+j]
			} else {
				// If we run out of lines, pad with empty strings
				chunk[j] = ""
			}
		}

		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

// SplitChunkIntoCharacters splits a chunk (8 lines) into individual character patterns
// Returns a slice of character patterns, where each pattern is 8 lines
func SplitChunkIntoCharacters(chunk []string, charWidth int) [][]string {
	if len(chunk) != 8 {
		return nil
	}

	// Find the maximum line length
	maxLen := 0
	for _, line := range chunk {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// If maxLen is 0, return empty
	if maxLen == 0 {
		return nil
	}

	var characters [][]string

	// Extract characters by column position
	pos := 0
	for pos < maxLen {
		// Extract character pattern at current position
		pattern := make([]string, 8)
		endPos := pos + charWidth

		for i := 0; i < 8; i++ {
			if pos < len(chunk[i]) {
				if endPos <= len(chunk[i]) {
					pattern[i] = chunk[i][pos:endPos]
				} else {
					// Pad with spaces if line is too short
					pattern[i] = chunk[i][pos:]
					for len(pattern[i]) < charWidth {
						pattern[i] += " "
					}
				}
			} else {
				// Line doesn't reach this position, fill with spaces
				pattern[i] = strings.Repeat(" ", charWidth)
			}
		}

		characters = append(characters, pattern)
		pos += charWidth
	}

	return characters
}

// NormalizePattern ensures all lines in a pattern have the same length
// Pads shorter lines with spaces
func NormalizePattern(pattern []string) []string {
	if len(pattern) == 0 {
		return pattern
	}

	// Find maximum length
	maxLen := 0
	for _, line := range pattern {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Normalize all lines to maxLen
	normalized := make([]string, len(pattern))
	for i, line := range pattern {
		if len(line) < maxLen {
			normalized[i] = line + strings.Repeat(" ", maxLen-len(line))
		} else {
			normalized[i] = line
		}
	}

	return normalized
}
