package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func printSeparator() {
	fmt.Println(colorCyan + "═══════════════════════════════════════════════════════════════════════════════" + colorReset)
}

func printTestHeader(testNum int, description string) {
	printSeparator()
	fmt.Printf("%s📋 TEST #%d: %s%s\n", colorYellow, testNum, description, colorReset)
	printSeparator()
}

func printTestResult(passed bool, expected, got string) {
	fmt.Printf("%s📤 Expected:%s   %q\n", colorBlue, colorReset, expected)
	fmt.Printf("%s📥 Got:%s        %q\n", colorWhite, colorReset, got)
	fmt.Println()

	if passed {
		fmt.Printf("%s✅ PASS - Output matches expected%s\n", colorGreen, colorReset)
	} else {
		fmt.Printf("%s❌ FAIL - Output does not match expected%s\n", colorRed, colorReset)
	}
	fmt.Println()
}

func TestReverseE2E(t *testing.T) {
	// Create temp directory for test files
	tempDir := t.TempDir()

	fmt.Println(colorPurple + "\n╔════════════════════════════════════════════════════════════════════════════╗" + colorReset)
	fmt.Println(colorPurple + "║                    ASCII ART REVERSE - E2E TEST SUITE                      ║" + colorReset)
	fmt.Println(colorPurple + "╚════════════════════════════════════════════════════════════════════════════╝" + colorReset)
	fmt.Println()

	testNum := 1

	// Test 1: Invalid flag format (missing =)
	printTestHeader(testNum, "Invalid flag format (missing =)")
	testNum++

	input := "--reverse example.txt"
	fmt.Printf("%s📝 Command:%s    go run ./cmd %s\n", colorCyan, colorReset, input)
	fmt.Printf("%s📤 Expected:%s   Usage message\n", colorBlue, colorReset)
	fmt.Println()

	cmd := exec.Command("go", "run", "./cmd", "--reverse", "example.txt")
	cmd.Dir = "../.." // Set working directory to project root
	output, _ := cmd.CombinedOutput()
	outputStr := string(output)

	// Simply check if output contains "Usage"
	passed := strings.Contains(outputStr, "Usage")

	if passed {
		fmt.Printf("%s✅ PASS - Shows usage message%s\n", colorGreen, colorReset)
		fmt.Printf("%s📥 Output:%s\n%s\n", colorWhite, colorReset, strings.TrimSpace(outputStr))
	} else {
		t.Errorf("Test #1 failed: expected usage message")
		fmt.Printf("%s❌ FAIL - No usage message found%s\n", colorRed, colorReset)
		fmt.Printf("%s📥 Output:%s\n%s\n", colorWhite, colorReset, strings.TrimSpace(outputStr))
	}
	fmt.Println()

	// Test 2: Simple text "Hello World"
	printTestHeader(testNum, "Reverse ASCII art to text")
	testNum++

	inputText := "Hello World"
	helloWorldArt := generateAsciiArt(t, inputText, "standard")
	testFile := filepath.Join(tempDir, "hello_world.txt")
	err := os.WriteFile(testFile, []byte(helloWorldArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result := runReverse(t, testFile)
	expected := "Hello World"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 3: Numbers "123"
	printTestHeader(testNum, "Reverse numbers")
	testNum++

	inputText = "123"
	numbersArt := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "numbers.txt")
	err = os.WriteFile(testFile, []byte(numbersArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "123"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 4: Special characters "#=\["
	printTestHeader(testNum, "Reverse special characters")
	testNum++

	inputText = "#=\\["
	specialArt := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "special.txt")
	err = os.WriteFile(testFile, []byte(specialArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "#=\\["
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 5: Mixed "something&234"
	printTestHeader(testNum, "Reverse mixed alphanumeric with special chars")
	testNum++

	inputText = "something&234"
	mixedArt := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "mixed.txt")
	err = os.WriteFile(testFile, []byte(mixedArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "something&234"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 6: Lowercase alphabet
	printTestHeader(testNum, "Reverse lowercase alphabet")
	testNum++

	inputText = "abcdefghijklmnopqrstuvwxyz"
	lowerAlphabetArt := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "lowercase.txt")
	err = os.WriteFile(testFile, []byte(lowerAlphabetArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "abcdefghijklmnopqrstuvwxyz"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 7: Special characters set 1
	printTestHeader(testNum, "Reverse special character set 1")
	testNum++

	inputText = "\\!\\\" #$%&'()*+,-./"
	specialSet1Art := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "special_set1.txt")
	err = os.WriteFile(testFile, []byte(specialSet1Art), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "\\!\\\" #$%&'()*+,-./"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 8: Special characters set 2
	printTestHeader(testNum, "Reverse special character set 2")
	testNum++

	inputText = ":;{=}?@"
	specialSet2Art := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "special_set2.txt")
	err = os.WriteFile(testFile, []byte(specialSet2Art), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = ":;{=}?@"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 9: Uppercase alphabet
	printTestHeader(testNum, "Reverse uppercase alphabet")
	testNum++

	inputText = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	upperAlphabetArt := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "uppercase.txt")
	err = os.WriteFile(testFile, []byte(upperAlphabetArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 10: Mixed case with spaces
	printTestHeader(testNum, "Reverse mixed case with spaces and numbers")
	testNum++

	inputText = "Hello World 123"
	mixedCaseArt := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "mixed_case.txt")
	err = os.WriteFile(testFile, []byte(mixedCaseArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "Hello World 123"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 11: Shadow banner
	printTestHeader(testNum, "Reverse with shadow banner")
	testNum++

	inputText = "Shadow"
	shadowArt := generateAsciiArt(t, inputText, "shadow")
	testFile = filepath.Join(tempDir, "shadow.txt")
	err = os.WriteFile(testFile, []byte(shadowArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file (shadow banner) containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "Shadow"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 12: Thinkertoy banner
	printTestHeader(testNum, "Reverse with thinkertoy banner")
	testNum++

	inputText = "Think"
	thinkerArt := generateAsciiArt(t, inputText, "thinkertoy")
	testFile = filepath.Join(tempDir, "thinkertoy.txt")
	err = os.WriteFile(testFile, []byte(thinkerArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file (thinkertoy banner) containing '%s'\n", colorCyan, colorReset, inputText)
	result = runReverse(t, testFile)
	expected = "Think"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Test 13: Multiline text
	printTestHeader(testNum, "Reverse multiline text")
	testNum++

	inputText = "Hello\\nWorld"
	multilineArt := generateAsciiArt(t, inputText, "standard")
	testFile = filepath.Join(tempDir, "multiline.txt")
	err = os.WriteFile(testFile, []byte(multilineArt), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fmt.Printf("%s📝 Input:%s       ASCII art file containing 'Hello' and 'World' on separate lines\n", colorCyan, colorReset)
	result = runReverse(t, testFile)
	expected = "Hello\nWorld"
	passed = result == expected
	printTestResult(passed, expected, result)
	if !passed {
		t.Errorf("Test failed")
	}

	// Final summary
	printSeparator()
	fmt.Printf("%s🎉 E2E TEST SUITE COMPLETE%s\n", colorGreen, colorReset)
	printSeparator()
	fmt.Println()
}

// generateAsciiArt generates ASCII art using the main program
func generateAsciiArt(t *testing.T, text, banner string) string {
	cmd := exec.Command("go", "run", "./cmd", text, banner)
	cmd.Dir = "../.." // Set working directory to project root
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to generate ASCII art for %q: %v\nOutput: %s", text, err, output)
	}
	return string(output)
}

// runReverse runs the reverse command and returns the output
func runReverse(t *testing.T, filename string) string {
	cmd := exec.Command("go", "run", "./cmd", fmt.Sprintf("--reverse=%s", filename))
	cmd.Dir = "../.." // Set working directory to project root
	output, err := cmd.Output()
	if err != nil {
		stderr := ""
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr = string(exitErr.Stderr)
		}
		t.Fatalf("Failed to run reverse on %s: %v\nStderr: %s", filename, err, stderr)
	}
	return strings.TrimSpace(string(output))
}
