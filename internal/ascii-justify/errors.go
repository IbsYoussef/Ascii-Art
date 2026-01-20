package asciijustify

import "fmt"

// Usage message for the align feature
const UsageAlign = `Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run ./cmd --align=right "text" standard`

// Error messages
var (
	// ErrInvalidAlignFormat is returned when --align flag format is incorrect
	ErrInvalidAlignFormat = fmt.Errorf("invalid --align flag format\n%s", UsageAlign)

	
	// ErrInvalidAlignType is returned when alignment type is not valid
	ErrInvalidAlignType = fmt.Errorf("invalid alignment type\nValid types: left, center, right, justify")

	// ErrContentTooWide is returned when content doesn't fit terminal
	ErrContentTooWide = fmt.Errorf("content too wide for terminal")
)

// WrapAlignTypeError wraps an invalid alignment type error
func WrapAlignTypeError(alignType string) error {
	return fmt.Errorf("invalid alignment type: %s\nValid types: left, center, right, justify", alignType)
}