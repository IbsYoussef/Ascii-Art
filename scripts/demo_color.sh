#!/bin/bash

# demo_color.sh - Showcase color features
# This script demonstrates the --color flag with various formats

clear

echo "════════════════════════════════════════════════════════════════════════════════"
echo "                     ASCII ART - COLOR FEATURES DEMO                           "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""
sleep 2

# Demo 1: Named color - Red
echo "▶ Demo 1: Named color - Red"
echo "  Command: go run ./cmd --color=red \"Hello\" standard"
echo ""
sleep 1
go run ./cmd --color=red "Hello" standard
echo ""
sleep 3

# Demo 2: Named color - Blue
echo "▶ Demo 2: Named color - Blue"
echo "  Command: go run ./cmd --color=blue \"World\" shadow"
echo ""
sleep 1
go run ./cmd --color=blue "World" shadow
echo ""
sleep 3

# Demo 3: Named color - Green
echo "▶ Demo 3: Named color - Green"
echo "  Command: go run ./cmd --color=green \"ASCII\" thinkertoy"
echo ""
sleep 1
go run ./cmd --color=green "ASCII" thinkertoy
echo ""
sleep 3

# Demo 4: Hex color - Orange
echo "▶ Demo 4: Hex color - Orange (#FF5733)"
echo "  Command: go run ./cmd --color=#FF5733 \"Hex\" standard"
echo ""
sleep 1
go run ./cmd --color=#FF5733 "Hex" standard
echo ""
sleep 3

# Demo 5: Hex color - Pink
echo "▶ Demo 5: Hex color - Pink (#FF69B4)"
echo "  Command: go run ./cmd --color=#FF69B4 \"Color\" shadow"
echo ""
sleep 1
go run ./cmd --color=#FF69B4 "Color" shadow
echo ""
sleep 3

# Demo 6: RGB color - Cyan
echo "▶ Demo 6: RGB color - Cyan"
echo "  Command: go run ./cmd --color='rgb(0,255,255)' \"RGB\" standard"
echo ""
sleep 1
go run ./cmd --color='rgb(0,255,255)' "RGB" standard
echo ""
sleep 3

# Demo 7: RGB color - Purple
echo "▶ Demo 7: RGB color - Purple"
echo "  Command: go run ./cmd --color='rgb(128,0,128)' \"Art\" thinkertoy"
echo ""
sleep 1
go run ./cmd --color='rgb(128,0,128)' "Art" thinkertoy
echo ""
sleep 3

# Demo 8: HSL color - Yellow
echo "▶ Demo 8: HSL color - Yellow"
echo "  Command: go run ./cmd --color='hsl(60,100%,50%)' \"HSL\" standard"
echo ""
sleep 1
go run ./cmd --color='hsl(60,100%,50%)' "HSL" standard
echo ""
sleep 3

# Demo 9: Substring coloring - "World" in blue
echo "▶ Demo 9: Substring coloring (only 'World' is blue)"
echo "  Command: go run ./cmd --color=blue World \"Hello World\" standard"
echo ""
sleep 1
go run ./cmd --color=blue World "Hello World" standard
echo ""
sleep 3

# Demo 10: Substring coloring - "Go" in red
echo "▶ Demo 10: Substring coloring (only 'Go' is red)"
echo "  Command: go run ./cmd --color=red Go \"Let's Go!\" shadow"
echo ""
sleep 1
go run ./cmd --color=red Go "Let's Go!" shadow
echo ""
sleep 3

# Demo 11: Multiple colors demo
echo "▶ Demo 11: Color variety showcase"
echo "  Displaying multiple colors..."
echo ""
sleep 1
go run ./cmd --color=yellow "Vibrant" standard
echo ""
sleep 2
go run ./cmd --color=magenta "Colors" shadow
echo ""
sleep 2
go run ./cmd --color=cyan "Everywhere!" thinkertoy
echo ""
sleep 2

echo "════════════════════════════════════════════════════════════════════════════════"
echo "                          ✅ COLOR DEMO COMPLETE!                              "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""