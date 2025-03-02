package main

import (
	utils "ascii-art/util"
	"os"
	"strings"
)

func main() {
	str := strings.Join(os.Args[1:], " ")
	words := strings.Split(str, `\n`)
	filename := "standard"
	if len(os.Args[1:]) == 2 {
		filename = os.Args[2]
	}

	bannerLines, err := utils.ReadBannerFile("Banners/" + filename + ".txt")
	if err != nil {
		panic(err)
	}

	utils.GenerateASCIIArt(words, bannerLines)
}
