package asciijustify

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

// winsize is the struct for terminal size
type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// GetTerminalWidth returns the current terminal width
// Falls back to 80 columns if detection fails
func GetTerminalWidth() int {
	ws := &winsize{}

	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) != -1 && errno == 0 && ws.Col > 0 {
		return int(ws.Col)
	}

	// Fallback to tput command
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err == nil {
		width, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err == nil && width > 0 {
			return width
		}
	}

	// Final fallback to default width
	return 80
}

// FitsInTerminal checks if the given width fits in terminal
func FitsInTerminal(contentWidth int) bool {
	termWidth := GetTerminalWidth()
	return contentWidth <= termWidth
}

// GetArtLineWidth returns the width of a single line of ASCII art
// Returns the length of the first line (all lines should be same width)
func GetArtLineWidth(charArt []string) int {
	if len(charArt) == 0 {
		return 0
	}
	return len(charArt[0])
}

// CalculatePadding calculates the padding needed for alignment
// Returns the number of spaces to add before the content
func CalculatePadding(contentWidth, terminalWidth int, alignType string) int {
	if contentWidth >= terminalWidth {
		return 0
	}

	switch alignType {
	case "center":
		return (terminalWidth - contentWidth) / 2
	case "right":
		return terminalWidth - contentWidth
	case "left":
		return 0
	default:
		return 0
	}
}
