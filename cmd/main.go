package main

import (
	"ascii-art/internal/ascii"
	color "ascii-art/internal/ascii-color"
	justify "ascii-art/internal/ascii-justify"
	output "ascii-art/internal/ascii-output"
	reverse "ascii-art/internal/ascii-reverse"
	"fmt"
	"os"
)

func main() {
	// Obtain user arguments from command line
	args := os.Args[1:]

	// Priority 1: Check for --reverse flag first (takes precedence over everything)
	if reverse.HasReverseFlag(args) {
		reverse.HandleReverse(args)
		return
	}

	// Priority 2: Parse --align flag
	alignType, remainingArgs, err := justify.ParseAlignFlag(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Priority 3: Parse --output flag
	outputFile, remainingArgs, err := output.ParseOutputFlag(remainingArgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Temporarily replace os.Args with remaining args for color parsing
	// This allows GetUserInputWithColor to work as if --align and --output flags weren't there
	originalArgs := os.Args
	os.Args = append([]string{os.Args[0]}, remainingArgs...)

	// Priority 4: Get user input and banner choice
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

	// Handle output with alignment
	if outputFile != "" {
		// If output flag is set, write to file (alignment not applied to file output)
		err = output.HandleOutput(outputFile, renderFunc)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// No output file, apply alignment to stdout
		// Pass input and banner for justify to work properly
		err = justify.HandleJustify(renderFunc, alignType, input, result)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
