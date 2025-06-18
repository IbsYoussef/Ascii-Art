package ascii

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRenderAscii(t *testing.T) {
	// Step 1: Create a mock banner map
	mockBanner := map[rune][]string{
		'A': {
			"   A   ",
			"  A A  ",
			" A   A ",
			" AAAAA ",
			" A   A ",
			" A   A ",
			" A   A ",
			"       ",
		},
	}

	// Step 2: Capture stdout
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Step 3: Call RenderAscii with test input
	RenderAscii("A", mockBanner)

	// Step 4: Read and restore stdout
	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)

	// Step 5: Check output
	expected := strings.Join(mockBanner['A'], "\n") + "\n" // fmt.Println adds newlines
	if buf.String() != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, buf.String())
	}
}
