package main

import (
	"ascii-art/internal/ascii"
	color "ascii-art/internal/ascii-color"
	output "ascii-art/internal/ascii-output"
	reverse "ascii-art/internal/ascii-reverse"
	"fmt"
	"os"
)

func main() {
	// Obtain user arguements from command line
	args := os.Args[1:]

	// Priority 1: Check for --reverse flag first (takes precendence over everything)
	if reverse.HasReverseFlag(args) {
		reverse.HandleReverse(args)
		return
	}

	// Priority 2: Parse --output flag
	outputFile, remainingArgs, err := output.ParseOutputFlag(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Temporarily replace os.Args with remaining args for color parsing
	// This allows GetUserInputWithColor to work as if --output flag wasn't there
	originalArgs := os.Args
	os.Args = append([]string{os.Args[0]}, remainingArgs...)

	// Priority 3: Get user input and banner choice
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

	// Create the render function that will be executed
	renderFunc := func() {
		if colorConfig.Enabled {
			color.RenderAsciiWithColor(input, result, colorConfig)
		} else {
			ascii.RenderAscii(input, result)
		}
	}

	err = output.HandleOutput(outputFile, renderFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
}
