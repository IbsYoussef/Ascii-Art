package util

import "fmt"

func GenerateASCIIArt(words []string, bannerLines []string) {
	for i, word := range words {
		if word == "" {
			if i < len(words)-1 {
				fmt.Println()
			}
			continue
		}
		PrintWordBanner(word, bannerLines)
	}
}
