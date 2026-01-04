package main

import (
	"ascii-art/internal/ascii"
	"fmt"
)

func main() {
	// Get user input and banner choice
	input, banner, colorConfig, err := ascii.GetUserInputWithColor()
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
		ascii.RenderAsciiWithColor(input, result, colorConfig)
	} else {
		ascii.RenderAscii(input, result)
	}
}
