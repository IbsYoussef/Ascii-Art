#!/bin/bash

# demo_output.sh - Showcase output to file features
# This script demonstrates the --output flag

clear

echo "════════════════════════════════════════════════════════════════════════════════"
echo "                     ASCII ART - OUTPUT FEATURES DEMO                          "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""
sleep 2

# Create temp directory for demo files
DEMO_DIR="demo_outputs"
mkdir -p "$DEMO_DIR"
echo "📁 Created demo directory: $DEMO_DIR"
echo ""
sleep 2

# Demo 1: Basic output to file
echo "▶ Demo 1: Save ASCII art to file"
echo "  Command: go run ./cmd --output=$DEMO_DIR/hello.txt \"Hello\" standard"
echo ""
sleep 1
go run ./cmd --output=$DEMO_DIR/hello.txt "Hello" standard
echo "✅ File created: $DEMO_DIR/hello.txt"
echo ""
echo "📄 File contents:"
cat $DEMO_DIR/hello.txt
echo ""
sleep 3

# Demo 2: Output with shadow banner
echo "▶ Demo 2: Save with shadow banner"
echo "  Command: go run ./cmd --output=$DEMO_DIR/shadow.txt \"World\" shadow"
echo ""
sleep 1
go run ./cmd --output=$DEMO_DIR/shadow.txt "World" shadow
echo "✅ File created: $DEMO_DIR/shadow.txt"
echo ""
echo "📄 Preview (first 8 lines):"
head -n 8 $DEMO_DIR/shadow.txt
echo ""
sleep 3

# Demo 3: Output with thinkertoy banner
echo "▶ Demo 3: Save with thinkertoy banner"
echo "  Command: go run ./cmd --output=$DEMO_DIR/think.txt \"ASCII\" thinkertoy"
echo ""
sleep 1
go run ./cmd --output=$DEMO_DIR/think.txt "ASCII" thinkertoy
echo "✅ File created: $DEMO_DIR/think.txt"
echo ""
echo "📄 Preview (first 8 lines):"
head -n 8 $DEMO_DIR/think.txt
echo ""
sleep 3

# Demo 4: Output with color (ANSI codes preserved)
echo "▶ Demo 4: Save colored ASCII art"
echo "  Command: go run ./cmd --output=$DEMO_DIR/colored.txt --color=red \"Color\" standard"
echo ""
sleep 1
go run ./cmd --output=$DEMO_DIR/colored.txt --color=red "Color" standard
echo "✅ File created: $DEMO_DIR/colored.txt (with ANSI color codes)"
echo ""
echo "📄 Viewing file with color rendering:"
cat $DEMO_DIR/colored.txt
echo ""
sleep 3

# Demo 5: Output with color substring
echo "▶ Demo 5: Save with substring coloring"
echo "  Command: go run ./cmd --output=$DEMO_DIR/substring.txt --color=blue Art \"ASCII Art\" shadow"
echo ""
sleep 1
go run ./cmd --output=$DEMO_DIR/substring.txt --color=blue Art "ASCII Art" shadow
echo "✅ File created: $DEMO_DIR/substring.txt"
echo ""
echo "📄 Viewing file with color rendering:"
cat $DEMO_DIR/substring.txt
echo ""
sleep 3

# Demo 6: Multiline output
echo "▶ Demo 6: Save multiline text"
echo "  Command: go run ./cmd --output=$DEMO_DIR/multiline.txt \"Line1\\nLine2\" standard"
echo ""
sleep 1
go run ./cmd --output=$DEMO_DIR/multiline.txt "Line1\nLine2" standard
echo "✅ File created: $DEMO_DIR/multiline.txt"
echo ""
echo "📄 File contents:"
cat $DEMO_DIR/multiline.txt
echo ""
sleep 3

# Demo 7: Special characters
echo "▶ Demo 7: Save special characters"
echo "  Command: go run ./cmd --output=$DEMO_DIR/special.txt \"123!@#\" shadow"
echo ""
sleep 1
go run ./cmd --output=$DEMO_DIR/special.txt "123!@#" shadow
echo "✅ File created: $DEMO_DIR/special.txt"
echo ""
echo "📄 Preview (first 8 lines):"
head -n 8 $DEMO_DIR/special.txt
echo ""
sleep 3

# Demo 8: Multiple color formats in output
echo "▶ Demo 8: Save with hex color"
echo "  Command: go run ./cmd --output=$DEMO_DIR/hex.txt --color=#FF5733 \"Hex\" standard"
echo ""
sleep 1
go run ./cmd --output=$DEMO_DIR/hex.txt --color=#FF5733 "Hex" standard
echo "✅ File created: $DEMO_DIR/hex.txt"
echo ""
echo "📄 Viewing file with color rendering:"
cat $DEMO_DIR/hex.txt
echo ""
sleep 3

# Show all created files
echo "════════════════════════════════════════════════════════════════════════════════"
echo "                     📂 ALL GENERATED FILES                                     "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""
ls -lh $DEMO_DIR/
echo ""
sleep 2

# Demonstrate viewing files
echo "════════════════════════════════════════════════════════════════════════════════"
echo "                     🎨 VIEWING COLORED FILES                                   "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""
echo "When you 'cat' a colored file, the colors appear:"
echo ""
sleep 2

echo "▶ Viewing colored.txt:"
cat $DEMO_DIR/colored.txt
echo ""
sleep 2

echo "▶ Viewing hex.txt:"
cat $DEMO_DIR/hex.txt
echo ""
sleep 2

echo "════════════════════════════════════════════════════════════════════════════════"
echo "                          ✅ OUTPUT DEMO COMPLETE!                             "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""
echo "💡 Tip: Files are saved in '$DEMO_DIR/' directory"
echo "    You can view any file with: cat $DEMO_DIR/<filename>"
echo ""

# Optional: Clean up demo files
read -p "🗑️  Delete demo files? (y/n): " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
    rm -rf "$DEMO_DIR"
    echo "✅ Demo files cleaned up!"
else
    echo "📁 Demo files preserved in $DEMO_DIR/"
fi
echo ""