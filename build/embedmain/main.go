package main

import (
	"ascii-art/internal/ascii"
	"ascii-art/internal/asciiembed"
	"fmt"
)

func main() {
	input, banner, err := ascii.GetUserInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	bannerMap, err := asciiembed.LoadBanner(banner)
	if err != nil {
		fmt.Println("Failed to load banner:", err)
		return
	}

	ascii.RenderAscii(input, bannerMap)
}
