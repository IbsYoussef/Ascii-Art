package asciiembed

// LoadBanner reads a banner file from embedded FS
func LoadBanner(name string) (map[rune][]string, error) {
	file, err := BannerFiles.Open("banners/" + name + ".txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return parseBannerFromReader(file)
}
