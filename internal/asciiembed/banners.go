package asciiembed

import "embed"

//go:embed banners/*.txt
var BannerFiles embed.FS
