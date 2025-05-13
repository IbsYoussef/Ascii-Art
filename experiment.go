package main

import (
	"ascii-art/util"
	"fmt"
	"os"
)

func main() {
	input, err := util.ParseInput(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output for testing
	fmt.Printf("Text to convert: '%s'\n", input.Text)
	fmt.Printf("Using Banner: '%s'\n", input.Banner)
}
