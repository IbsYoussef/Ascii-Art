package main

import (
	"ascii-art/internal/ascii"
	color "ascii-art/internal/ascii-color"
	output "ascii-art/internal/ascii-output"
	"fmt"
	"os"
)

func main() {
	// Parse --output flag first
	outputFile, remainingArgs, err := output.ParseOutputFlag(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Temporarily replace os.Args with remaining args for color parsing
	// This allows GetUserInputWithColor to work as if --output flag wasn't there
	originalArgs := os.Args
	os.Args = append([]string{os.Args[0]}, remainingArgs...)

	// Get user input and banner choice
	input, banner, colorConfig, err := color.GetUserInputWithColor()

	// Restore original args
	os.Args = originalArgs

	if err != nil {
		fmt.Println(err)
		return
	}

	// Load the ASCII banner from the specified file
	result, err := ascii.LoadBannerFile(banner)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create render function that will be executed
	renderFunc := func() {
		if colorConfig.Enabled {
			color.RenderAsciiWithColor(input, result, colorConfig)
		} else {
			ascii.RenderAscii(input, result)
		}
	}

	// Handle output routing (to file or stdout)
	err = output.HandleOutput(outputFile, renderFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
}
