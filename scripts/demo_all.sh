#!/bin/bash

# demo_all.sh - Master script to run all demos
# Usage: ./demo_all.sh [standard|color|output|reverse|all]

clear

show_menu() {
    echo "════════════════════════════════════════════════════════════════════════════════"
    echo "                        ASCII ART - DEMO SELECTOR                              "
    echo "════════════════════════════════════════════════════════════════════════════════"
    echo ""
    echo "  Choose a demo to run:"
    echo ""
    echo "    1) Standard Features Demo  (demo_standard.sh)"
    echo "    2) Color Features Demo     (demo_color.sh)"
    echo "    3) Output Features Demo    (demo_output.sh)"
    echo "    4) Reverse Features Demo   (demo_reverse.sh)  ⭐ NEW!"
    echo "    5) Run All Demos           (all four in sequence)"
    echo "    6) Exit"
    echo ""
    echo "════════════════════════════════════════════════════════════════════════════════"
    echo ""
}

run_standard() {
    echo "🚀 Running Standard Features Demo..."
    sleep 1
    bash scripts/demo_standard.sh
}

run_color() {
    echo "🚀 Running Color Features Demo..."
    sleep 1
    bash scripts/demo_color.sh
}

run_output() {
    echo "🚀 Running Output Features Demo..."
    sleep 1
    bash scripts/demo_output.sh
}

run_reverse() {
    echo "🚀 Running Reverse Features Demo..."
    sleep 1
    bash scripts/demo_reverse.sh
}

run_all() {
    echo "🚀 Running All Demos in Sequence..."
    echo ""
    sleep 2
    
    run_standard
    echo ""
    echo "Press Enter to continue to Color Demo..."
    read
    clear
    
    run_color
    echo ""
    echo "Press Enter to continue to Output Demo..."
    read
    clear
    
    run_output
    echo ""
    echo "Press Enter to continue to Reverse Demo..."
    read
    clear
    
    run_reverse
    
    echo ""
    echo "════════════════════════════════════════════════════════════════════════════════"
    echo "                     ✅ ALL DEMOS COMPLETE!                                    "
    echo "════════════════════════════════════════════════════════════════════════════════"
    echo ""
}

# Handle command line argument
if [ $# -eq 1 ]; then
    case "$1" in
        standard)
            run_standard
            exit 0
            ;;
        color)
            run_color
            exit 0
            ;;
        output)
            run_output
            exit 0
            ;;
        reverse)
            run_reverse
            exit 0
            ;;
        all)
            run_all
            exit 0
            ;;
        *)
            echo "Invalid option: $1"
            echo "Usage: $0 [standard|color|output|reverse|all]"
            exit 1
            ;;
    esac
fi

# Interactive menu
while true; do
    show_menu
    read -p "Enter your choice [1-6]: " choice
    echo ""
    
    case $choice in
        1)
            clear
            run_standard
            echo ""
            read -p "Press Enter to return to menu..."
            clear
            ;;
        2)
            clear
            run_color
            echo ""
            read -p "Press Enter to return to menu..."
            clear
            ;;
        3)
            clear
            run_output
            echo ""
            read -p "Press Enter to return to menu..."
            clear
            ;;
        4)
            clear
            run_reverse
            echo ""
            read -p "Press Enter to return to menu..."
            clear
            ;;
        5)
            clear
            run_all
            echo ""
            read -p "Press Enter to return to menu..."
            clear
            ;;
        6)
            echo "Goodbye! 👋"
            exit 0
            ;;
        *)
            echo "❌ Invalid choice. Please enter 1-6."
            sleep 2
            clear
            ;;
    esac
done