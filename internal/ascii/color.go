package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

// ANSI color codes for named colors
var namedColors = map[string]string{
	"black":   "\033[38;2;0;0;0m",       // RGB black
	"red":     "\033[38;2;255;0;0m",     // RGB pure red
	"green":   "\033[38;2;0;255;0m",     // RGB pure green
	"yellow":  "\033[38;2;255;255;0m",   // RGB yellow
	"blue":    "\033[38;2;0;0;255m",     // RGB blue
	"magenta": "\033[38;2;255;0;255m",   // RGB magenta
	"cyan":    "\033[38;2;0;255;255m",   // RGB cyan
	"white":   "\033[38;2;255;255;255m", // RGB white
	"orange":  "\033[38;2;255;165;0m",   // RGB orange
	"reset":   "\033[0m",
}

// ParseColor converts various color formats to ANSI escape code
// Supports: named colors, hex (#FF0000), rgb(255,0,0), hsl(0,100%,50%)
func ParseColor(color string) (string, error) {
	color = strings.TrimSpace(strings.ToLower(color))

	// Check for named color
	if ansiCode, ok := namedColors[color]; ok {
		return ansiCode, nil
	}

	// Check for hex color (#FF0000 or #F00)
	if strings.HasPrefix(color, "#") {
		return parseHexColor(color)
	}

	// Check for rgb(r, g, b)
	if strings.HasPrefix(color, "rgb(") && strings.HasSuffix(color, ")") {
		return parseRGBColor(color)
	}

	// Check for hsl(h, s%, l%)
	if strings.HasPrefix(color, "hsl(") && strings.HasSuffix(color, ")") {
		return parseHSLColor(color)
	}

	return "", fmt.Errorf("unsupported color format: %s", color)
}

// parseHexColor converts hex color to ANSI RGB code
// Supports #RRGGBB and #RGB formats
func parseHexColor(hex string) (string, error) {
	hex = strings.TrimPrefix(hex, "#")

	// Expand short form (#RGB to #RRGGBB)
	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}

	if len(hex) != 6 {
		return "", fmt.Errorf("invalid hex color: must be #RRGGBB or #RGB")
	}

	// Parse RGB components
	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid hex color: %v", err)
	}

	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid hex color: %v", err)
	}

	b, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid hex color: %v", err)
	}

	// Return ANSI RGB code
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil
}

// parseRGBColor converts rgb(r, g, b) to ANSI RGB code
func parseRGBColor(rgb string) (string, error) {
	// Remove "rgb(" prefix and ")" suffix
	rgb = strings.TrimPrefix(rgb, "rgb(")
	rgb = strings.TrimSuffix(rgb, ")")

	// Split by comma
	parts := strings.Split(rgb, ",")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid RGB format: expected rgb(r,g,b)")
	}

	// Parse each component
	r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || r < 0 || r > 255 {
		return "", fmt.Errorf("invalid RGB red value: must be 0-255")
	}

	g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil || g < 0 || g > 255 {
		return "", fmt.Errorf("invalid RGB green value: must be 0-255")
	}

	b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil || b < 0 || b > 255 {
		return "", fmt.Errorf("invalid RGB blue value: must be 0-255")
	}

	// Return ANSI RGB code
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil
}

// parseHSLColor converts hsl(h, s%, l%) to ANSI RGB code
// HSL: Hue (0-360), Saturation (0-100%), Lightness (0-100%)
func parseHSLColor(hsl string) (string, error) {
	// Remove "hsl(" prefix and ")" suffix
	hsl = strings.TrimPrefix(hsl, "hsl(")
	hsl = strings.TrimSuffix(hsl, ")")

	// Split by comma
	parts := strings.Split(hsl, ",")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid HSL format: expected hsl(h,s%%,l%%)")
	}

	// Parse hue (0-360)
	h, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || h < 0 || h > 360 {
		return "", fmt.Errorf("invalid HSL hue: must be 0-360")
	}

	// Parse saturation (0-100%)
	sStr := strings.TrimSpace(strings.TrimSuffix(parts[1], "%"))
	s, err := strconv.Atoi(sStr)
	if err != nil || s < 0 || s > 100 {
		return "", fmt.Errorf("invalid HSL saturation: must be 0-100%%")
	}

	// Parse lightness (0-100%)
	lStr := strings.TrimSpace(strings.TrimSuffix(parts[2], "%"))
	l, err := strconv.Atoi(lStr)
	if err != nil || l < 0 || l > 100 {
		return "", fmt.Errorf("invalid HSL lightness: must be 0-100%%")
	}

	// Convert HSL to RGB
	r, g, b := hslToRGB(h, s, l)

	// Return ANSI RGB code
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil
}

// hslToRGB converts HSL color to RGB
func hslToRGB(h, s, l int) (r, g, b int) {
	// Normalize values
	hue := float64(h) / 360.0
	sat := float64(s) / 100.0
	light := float64(l) / 100.0

	var r1, g1, b1 float64

	if sat == 0 {
		// Achromatic (gray)
		r1, g1, b1 = light, light, light
	} else {
		var q float64
		if light < 0.5 {
			q = light * (1 + sat)
		} else {
			q = light + sat - light*sat
		}
		p := 2*light - q

		r1 = hueToRGB(p, q, hue+1.0/3.0)
		g1 = hueToRGB(p, q, hue)
		b1 = hueToRGB(p, q, hue-1.0/3.0)
	}

	r = int(r1 * 255)
	g = int(g1 * 255)
	b = int(b1 * 255)

	return r, g, b
}

// hueToRGB is a helper function for HSL to RGB conversion
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

// ApplyColor wraps text with ANSI color code
func ApplyColor(text, ansiCode string) string {
	return ansiCode + text + ResetColor()
}

// ResetColor returns the ANSI reset code
func ResetColor() string {
	return "\033[0m"
}
