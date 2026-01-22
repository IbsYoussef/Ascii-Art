#!/bin/bash

# demo_standard.sh - Showcase standard ASCII art features
# This script demonstrates the core functionality without any flags

clear

echo "════════════════════════════════════════════════════════════════════════════════"
echo "                     ASCII ART - STANDARD FEATURES DEMO                        "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""
sleep 2

# Demo 1: Simple text with standard banner
echo "▶ Demo 1: Simple text with standard banner"
echo "  Command: go run ./cmd \"Hello\" standard"
echo ""
sleep 1
go run ./cmd "Hello" standard
echo ""
sleep 3

# Demo 2: Text with shadow banner
echo "▶ Demo 2: Text with shadow banner"
echo "  Command: go run ./cmd \"World\" shadow"
echo ""
sleep 1
go run ./cmd "World" shadow
echo ""
sleep 3

# Demo 3: Text with thinkertoy banner
echo "▶ Demo 3: Text with thinkertoy banner"
echo "  Command: go run ./cmd \"ASCII\" thinkertoy"
echo ""
sleep 1
go run ./cmd "ASCII" thinkertoy
echo ""
sleep 3

# Demo 4: Default banner (omit banner argument)
echo "▶ Demo 4: Default banner (standard is default)"
echo "  Command: go run ./cmd \"Default\""
echo ""
sleep 1
go run ./cmd "Default"
echo ""
sleep 3

# Demo 5: Multiline text
echo "▶ Demo 5: Multiline text with \\n"
echo "  Command: go run ./cmd \"First\\nLine\" standard"
echo ""
sleep 1
go run ./cmd "First\nLine" standard
echo ""
sleep 3

# Demo 6: Special characters
echo "▶ Demo 6: Special characters"
echo "  Command: go run ./cmd \"123 -> #\$%\" standard"
echo ""
sleep 1
go run ./cmd "123 -> #\$%" standard
echo ""
sleep 3

# Demo 7: Long text
echo "▶ Demo 7: Longer text"
echo "  Command: go run ./cmd \"Art!\" shadow"
echo ""
sleep 1
go run ./cmd "Art!" shadow
echo ""
sleep 2

echo "════════════════════════════════════════════════════════════════════════════════"
echo "                          ✅ DEMO COMPLETE!                                    "
echo "════════════════════════════════════════════════════════════════════════════════"
echo ""