package main

import (
	"ascii-art/internal/ascii"
	color "ascii-art/internal/ascii-color"
	"fmt"
)

func main() {
	// Get user input and banner choice
	input, banner, colorConfig, err := color.GetUserInputWithColor()
	// Check if error occured when getting user input
	if err != nil {
		fmt.Println(err)
		return
	}

	// Load the ASCII banner from the specified file
	result, err := ascii.LoadBannerFile(banner)
	// Check if error occured with loading specific banner file
	if err != nil {
		fmt.Println(err)
		return
	}

	// Render the input in banner style as ASCII art, use colour rendering if color flag is present
	if colorConfig.Enabled {
		color.RenderAsciiWithColor(input, result, colorConfig)
	} else {
		ascii.RenderAscii(input, result)
	}
}
