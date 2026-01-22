#!/bin/bash

# demo_reverse.sh - Showcase reverse feature
# This script demonstrates the --reverse flag

clear

echo "════════════════════════════════════════════════════════════════════════════════"
echo "                     ASCII ART - REVERSE FEATURES DEMO                         "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""
sleep 2

# Create temp directory for demo files
DEMO_DIR="demo_reverse"
mkdir -p "$DEMO_DIR"
echo "📁 Created demo directory: $DEMO_DIR"
echo ""
sleep 2

# Demo 1: Basic reverse - standard banner
echo "▶ Demo 1: Reverse standard banner ASCII art"
echo "  Step 1: Generate ASCII art"
echo "  Command: go run ./cmd \"Hello\" standard > $DEMO_DIR/hello.txt"
echo ""
sleep 1
go run ./cmd "Hello" standard > $DEMO_DIR/hello.txt
echo "📄 Generated ASCII art:"
cat $DEMO_DIR/hello.txt
echo ""
sleep 3

echo "  Step 2: Reverse it back to text"
echo "  Command: go run ./cmd --reverse=$DEMO_DIR/hello.txt"
echo ""
sleep 1
RESULT=$(go run ./cmd --reverse=$DEMO_DIR/hello.txt)
echo "✅ Result: $RESULT"
echo ""
sleep 3

# Demo 2: Shadow banner reverse
echo "▶ Demo 2: Reverse shadow banner ASCII art"
echo "  Step 1: Generate ASCII art"
echo "  Command: go run ./cmd \"World\" shadow > $DEMO_DIR/shadow.txt"
echo ""
sleep 1
go run ./cmd "World" shadow > $DEMO_DIR/shadow.txt
echo "📄 Generated ASCII art:"
cat $DEMO_DIR/shadow.txt
echo ""
sleep 3

echo "  Step 2: Reverse it back to text"
echo "  Command: go run ./cmd --reverse=$DEMO_DIR/shadow.txt"
echo ""
sleep 1
RESULT=$(go run ./cmd --reverse=$DEMO_DIR/shadow.txt)
echo "✅ Result: $RESULT"
echo ""
sleep 3

# Demo 3: Thinkertoy banner reverse
echo "▶ Demo 3: Reverse thinkertoy banner ASCII art"
echo "  Step 1: Generate ASCII art"
echo "  Command: go run ./cmd \"Think\" thinkertoy > $DEMO_DIR/think.txt"
echo ""
sleep 1
go run ./cmd "Think" thinkertoy > $DEMO_DIR/think.txt
echo "📄 Generated ASCII art:"
cat $DEMO_DIR/think.txt
echo ""
sleep 3

echo "  Step 2: Reverse it back to text"
echo "  Command: go run ./cmd --reverse=$DEMO_DIR/think.txt"
echo ""
sleep 1
RESULT=$(go run ./cmd --reverse=$DEMO_DIR/think.txt)
echo "✅ Result: $RESULT"
echo ""
sleep 3

# Demo 4: Numbers and special characters
echo "▶ Demo 4: Reverse with numbers and special characters"
echo "  Step 1: Generate ASCII art"
echo "  Command: go run ./cmd \"2024!\" standard > $DEMO_DIR/special.txt"
echo ""
sleep 1
go run ./cmd "2024!" standard > $DEMO_DIR/special.txt
echo "📄 Generated ASCII art:"
cat $DEMO_DIR/special.txt
echo ""
sleep 3

echo "  Step 2: Reverse it back to text"
echo "  Command: go run ./cmd --reverse=$DEMO_DIR/special.txt"
echo ""
sleep 1
RESULT=$(go run ./cmd --reverse=$DEMO_DIR/special.txt)
echo "✅ Result: $RESULT"
echo ""
sleep 3

# Demo 5: Full alphabet
echo "▶ Demo 5: Reverse full alphabet"
echo "  Step 1: Generate ASCII art"
echo "  Command: go run ./cmd \"ABCXYZ\" standard > $DEMO_DIR/alphabet.txt"
echo ""
sleep 1
go run ./cmd "ABCXYZ" standard > $DEMO_DIR/alphabet.txt
echo "📄 Generated ASCII art (showing first 10 lines):"
head -10 $DEMO_DIR/alphabet.txt
echo "..."
echo ""
sleep 3

echo "  Step 2: Reverse it back to text"
echo "  Command: go run ./cmd --reverse=$DEMO_DIR/alphabet.txt"
echo ""
sleep 1
RESULT=$(go run ./cmd --reverse=$DEMO_DIR/alphabet.txt)
echo "✅ Result: $RESULT"
echo ""
sleep 3

# Demo 6: Multiline text
echo "▶ Demo 6: Reverse multiline text"
echo "  Step 1: Generate ASCII art with newlines"
echo "  Command: go run ./cmd \"First\\nSecond\" standard > $DEMO_DIR/multiline.txt"
echo ""
sleep 1
go run ./cmd "First\nSecond" standard > $DEMO_DIR/multiline.txt
echo "📄 Generated ASCII art:"
cat $DEMO_DIR/multiline.txt
echo ""
sleep 3

echo "  Step 2: Reverse it back to text"
echo "  Command: go run ./cmd --reverse=$DEMO_DIR/multiline.txt"
echo ""
sleep 1
RESULT=$(go run ./cmd --reverse=$DEMO_DIR/multiline.txt)
echo "✅ Result (with newline preserved):"
echo "$RESULT"
echo ""
sleep 3

# Demo 7: Auto-banner detection
echo "▶ Demo 7: Auto-detection of banner type"
echo "  Creating ASCII art with different banners..."
echo ""
sleep 1

go run ./cmd "Test" standard > $DEMO_DIR/test_standard.txt
go run ./cmd "Test" shadow > $DEMO_DIR/test_shadow.txt
go run ./cmd "Test" thinkertoy > $DEMO_DIR/test_thinkertoy.txt

echo "  Reversing each file (automatic banner detection):"
echo ""
sleep 1

echo "  Standard banner → $(go run ./cmd --reverse=$DEMO_DIR/test_standard.txt)"
sleep 1
echo "  Shadow banner   → $(go run ./cmd --reverse=$DEMO_DIR/test_shadow.txt)"
sleep 1
echo "  Thinkertoy      → $(go run ./cmd --reverse=$DEMO_DIR/test_thinkertoy.txt)"
echo ""
echo "✅ All banners automatically detected!"
echo ""
sleep 3

# Demo 8: Error handling - invalid flag format
echo "▶ Demo 8: Error handling - invalid flag format"
echo "  Command: go run ./cmd --reverse example.txt (missing =)"
echo ""
sleep 1
go run ./cmd --reverse example.txt 2>&1 | head -5
echo ""
echo "✅ Proper error message shown"
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

echo "════════════════════════════════════════════════════════════════════════════════"
echo "                          ✅ REVERSE DEMO COMPLETE!                            "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""
echo "💡 Key Features Demonstrated:"
echo "   • Standard, shadow, and thinkertoy banner support"
echo "   • Automatic banner detection"
echo "   • Numbers and special characters"
echo "   • Multiline text preservation"
echo "   • Error handling for invalid flags"
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