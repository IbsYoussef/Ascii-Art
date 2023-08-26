package util

import "fmt"

func printBanner(words []string, bannerLines []string) {
	for i, word := range words {
		if word == "" {
			if i < len(words)-1 {
				fmt.Println()
			}
			continue
		}
		printWordBanner(word, bannerLines)
	}
}
