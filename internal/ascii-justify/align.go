package asciijustify

import "strings"

// ApplyAlignment applies the specified alignment to ASCII art lines
// Returns a new slice with aligned lines
func ApplyAlignment(asciiLines []string, alignType string, terminalWidth int) []string {
	if len(asciiLines) == 0 {
		return asciiLines
	}

	switch alignType {
	case "left":
		return asciiLines // Already left aligned by default
	case "center":
		return CenterAlign(asciiLines, terminalWidth)
	case "right":
		return RightAlign(asciiLines, terminalWidth)
	case "justify":
		return JustifyAlign(asciiLines, terminalWidth)
	default:
		return asciiLines
	}
}

// CenterAlign centers each line of ASCII art
func CenterAlign(lines []string, termWidth int) []string {
	result := make([]string, len(lines))
	for i, line := range lines {
		lineWidth := len(line)
		if lineWidth >= termWidth {
			result[i] = line
			continue
		}
		padding := (termWidth - lineWidth) / 2
		result[i] = strings.Repeat(" ", padding) + line
	}
	return result
}

// RightAlign right-aligns each line of ASCII art
func RightAlign(lines []string, termWidth int) []string {
	result := make([]string, len(lines))
	for i, line := range lines {
		lineWidth := len(line)
		if lineWidth >= termWidth {
			result[i] = line
			continue
		}
		padding := termWidth - lineWidth
		result[i] = strings.Repeat(" ", padding) + line
	}
	return result
}

// JustifyAlign distributes words evenly across terminal width
// For ASCII art, each group of 8 lines represents one line of text
func JustifyAlign(lines []string, termWidth int) []string {
	// Justify works on groups of 8 lines (one text line = 8 ASCII art lines)
	// We need to process each group separately

	result := make([]string, 0, len(lines))

	// Process in chunks of 8 lines
	for i := 0; i < len(lines); i += 8 {
		end := i + 8
		if end > len(lines) {
			end = len(lines)
		}

		chunk := lines[i:end]
		alignedChunk := JustifyChunk(chunk, termWidth)
		result = append(result, alignedChunk...)
	}
	return result
}

// JustifyChunk justifies a single 8-line chunk (one line of text)
// This extracts individual words and distributes them evenly across terminal width
// JustifyChunk justifies a single 8-line chunk (one line of text)
// This extracts individual words and distributes them evenly across terminal width
func JustifyChunk(chunk []string, termWidth int) []string {
	// ========================================
	// TUNING VARIABLES - Adjust these to fine-tune justify spacing
	// ========================================
	const (
		marginPercent    = 8  // Margin on each side (%) - higher = words closer to center
		maxSpacingPerGap = 45 // Maximum spaces between words - lower = tighter, higher = more spread
	)
	// ========================================

	if len(chunk) != 8 {
		return CenterAlign(chunk, termWidth)
	}

	// Step 1: Extract individual words from the ASCII art chunk
	words := extractWords(chunk)

	// If only one word or no words, center it instead
	if len(words) <= 1 {
		return CenterAlign(chunk, termWidth)
	}

	// Step 2: Calculate total width of all words
	totalWordWidth := 0
	for _, word := range words {
		wordWidth := getWordWidth(word)
		totalWordWidth += wordWidth
	}

	// Step 3: Apply margin
	margin := (termWidth * marginPercent) / 100
	usableWidth := termWidth - (2 * margin)

	// Check if content fits in usable width
	if totalWordWidth >= usableWidth {
		return CenterAlign(chunk, termWidth)
	}

	// Step 4: Calculate available space for gaps
	availableSpace := usableWidth - totalWordWidth
	if availableSpace <= 0 {
		return CenterAlign(chunk, termWidth)
	}

	// Step 5: Distribute space between words
	gaps := len(words) - 1
	if gaps <= 0 {
		return CenterAlign(chunk, termWidth)
	}

	spacePerGap := availableSpace / gaps

	// Step 5a: Check if spacing exceeds maximum allowed
	if spacePerGap > maxSpacingPerGap {
		// Spacing too wide, center instead
		return CenterAlign(chunk, termWidth)
	}

	extraSpace := availableSpace % gaps

	// Step 6: Build justified output by combining words with calculated spacing
	result := make([]string, 8)
	for row := 0; row < 8; row++ {
		// Start with left margin
		line := strings.Repeat(" ", margin)

		for wordIdx, word := range words {
			// Add the word's row
			if row < len(word) {
				line += word[row]
			}

			// Add spacing after word (except for last word)
			if wordIdx < len(words)-1 {
				spacing := spacePerGap
				// Distribute extra space to first few gaps
				if wordIdx < extraSpace {
					spacing++
				}
				line += strings.Repeat(" ", spacing)
			}
		}
		result[row] = line
	}

	return result
}

// extractWords splits an 8-line ASCII art chunk into individual words
// Returns a slice of words, where each word is represented as 8 lines
func extractWords(chunk []string) [][]string {
	if len(chunk) != 8 {
		return nil
	}

	// Find the maximum width
	maxWidth := 0
	for _, line := range chunk {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	if maxWidth == 0 {
		return nil
	}

	// Pad all lines to same width with spaces
	paddedChunk := make([]string, 8)
	for i, line := range chunk {
		if len(line) < maxWidth {
			paddedChunk[i] = line + strings.Repeat(" ", maxWidth-len(line))
		} else {
			paddedChunk[i] = line
		}
	}

	// Detect word boundaries by finding columns that are all spaces
	words := make([][]string, 0)
	currentWord := make([]string, 8)
	inWord := false
	consecutiveSpaces := 0

	for col := 0; col < maxWidth; col++ {
		// Check if this column is all spaces across all 8 rows
		allSpaces := true
		for row := 0; row < 8; row++ {
			if col < len(paddedChunk[row]) && paddedChunk[row][col] != ' ' {
				allSpaces = false
				break
			}
		}

		if allSpaces {
			consecutiveSpaces++
			// If we have 2+ consecutive space columns, it's a word boundary
			if consecutiveSpaces >= 2 && inWord {
				// Trim trailing spaces from current word
				trimmedWord := make([]string, 8)
				for i := 0; i < 8; i++ {
					trimmedWord[i] = strings.TrimRight(currentWord[i], " ")
				}
				words = append(words, trimmedWord)

				// Reset for next word
				currentWord = make([]string, 8)
				inWord = false
			}
		} else {
			consecutiveSpaces = 0
			inWord = true
			// Add this column to current word
			for row := 0; row < 8; row++ {
				if col < len(paddedChunk[row]) {
					currentWord[row] += string(paddedChunk[row][col])
				}
			}
		}
	}

	// Don't forget the last word
	if inWord {
		trimmedWord := make([]string, 8)
		for i := 0; i < 8; i++ {
			trimmedWord[i] = strings.TrimRight(currentWord[i], " ")
		}
		words = append(words, trimmedWord)
	}

	return words
}

// getWordWidth returns the width of a word (length of its first line)
func getWordWidth(word []string) int {
	if len(word) == 0 {
		return 0
	}
	return len(word[0])
}

// AlignLine applies alignment to a single line of text
// This is a helper function for simpler alignment scenarios
func AlignLine(line string, alignType string, termWidth int) string {
	lineWidth := len(line)
	if lineWidth >= termWidth {
		return line
	}

	switch alignType {
	case "center":
		padding := (termWidth - lineWidth) / 2
		return strings.Repeat(" ", padding) + line
	case "right":
		padding := termWidth - lineWidth
		return strings.Repeat(" ", padding) + line
	case "left":
		return line
	default:
		return line
	}
}
