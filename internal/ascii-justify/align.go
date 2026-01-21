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
		lindeWidth := len(line)
		if lindeWidth >= termWidth {
			result[i] = line
			continue
		}
		padding := (termWidth - lindeWidth) / 2
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
	// We need to process each group seperately

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

// justifyChunk justifies a single 8-line chunk (one line of text)
// For now, we'll implement basic justify that distributes space
func JustifyChunk(chunk []string, termWidth int) []string {
	if len(chunk) == 0 {
		return chunk
	}

	// For single-word lines or lines that already fit, center them instead
	// Full justify implementation will come in a later refinement
	// For now, use center alignment as a placeholder
	return CenterAlign(chunk, termWidth)
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
