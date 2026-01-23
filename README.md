# ASCII-ART

<div align="center">

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

![Basic Demo](assets/intro_demo.gif)

**Transform text into stylized ASCII art**

A lightweight CLI tool that renders text as ASCII art. Choose from multiple fonts, apply vibrant colors, and export your creations to files.

</div>

---

## 📋 Table of Contents

1. [🎯 About](#-about)
2. [✨ Features](#-features)
3. [🚀 Quick Start](#-quick-start)
4. [📖 Usage Guide](#-usage-guide)
5. [📁 Project Structure](#-project-structure)
6. [🧪 Testing](#-testing)
7. [🔭 Roadmap](#-roadmap)
8. [🙏 Acknowledgements](#-acknowledgements)
9. [📄 License](#-license)

---

## 🎯 About

**ASCII Art** is a command-line tool that transforms plain text into beautiful ASCII art with support for colors and file output. Built in Go as part of the **01 Founders** curriculum, this project demonstrates clean architecture, modular design, and comprehensive testing practices.

**Key Highlights:**

- 🎨 Multiple banner styles (standard, shadow, thinkertoy)
- 🌈 Full RGB/HSL/Hex color support with substring coloring
- 💾 Save output directly to files
- 🔄 Reverse ASCII art back to original text
- 📐 Text alignment and justification (left, right, center, justify)
- 🧪 100% test coverage with unit and E2E tests
- 📦 Zero dependencies - uses only Go standard library

---

## ✨ Features

### 🖊️ ASCII Art Generation

Transform any text into stylized ASCII art using three distinct banner fonts.

![Standard Demo](assets/standard_demo.gif)

**Supported Banners:**

- `standard` - Classic ASCII art style (default)
- `shadow` - Bold shadowed characters
- `thinkertoy` - Playful, creative font

**Capabilities:**

- Multi-line text support with `\n` escape sequences
- Special characters and numbers
- Handles spaces and punctuation
- Case-sensitive rendering

**Basic Commands:**

```bash
# Simple text with default banner
go run ./cmd "Hello"

# Choose a specific banner
go run ./cmd "World" shadow
go run ./cmd "ASCII" thinkertoy

# Multi-line text
go run ./cmd "First\nLine" standard
```

---

### 🎨 Color Support

Add vibrant colors to your ASCII art with multiple color format support.

![Color Demo](assets/color_demo.gif)

**Color Formats:**

- **Named Colors**: `red`, `blue`, `green`, `yellow`, `orange`, `pink`, `cyan`, `magenta`, `white`, `black`, `gray`
- **Hex Colors**: `#FF5733`, `#00FF00`, `#3498DB`
- **RGB Colors**: `rgb(255,87,51)`, `rgb(0,255,255)`
- **HSL Colors**: `hsl(9,100%,60%)`, `hsl(120,100%,50%)`

**Full Text Coloring:**

```bash
# Named color
go run ./cmd --color=red "Hello" standard

# Hex color
go run ./cmd --color=#FF5733 "Vibrant" shadow

# RGB color
go run ./cmd --color='rgb(0,255,255)' "Cyan" thinkertoy

# HSL color
go run ./cmd --color='hsl(120,100%,50%)' "Green" standard
```

**Substring Coloring:**

Color specific parts of your text (case-sensitive):

```bash
# Color only "World" in blue
go run ./cmd --color=blue World "Hello World" standard

# Color only "Go" in red
go run ./cmd --color=red Go "Let's Go!" shadow

# Color repeated substring
go run ./cmd --color=green kit "a king kitten have kit" standard
```

**Color Notes:**

- Substring matching is case-sensitive
- Colors apply to all matching occurrences
- RGB/HSL formats must be quoted to avoid shell interpretation

---

### 💾 Output to File

Save your ASCII art creations directly to files, with full color preservation.

![Output Demo](assets/output_demo.gif)

**Save to File:**

```bash
# Basic file output
go run ./cmd --output=banner.txt "Hello" standard

# With different banners
go run ./cmd --output=shadow.txt "World" shadow
go run ./cmd --output=think.txt "ASCII" thinkertoy
```

**Colored Output:**

Color codes are preserved in files and render when viewed in terminals:

```bash
# Save colored ASCII art
go run ./cmd --output=colored.txt --color=red "Color" standard

# View the colored file
cat colored.txt  # Colors appear in terminal!

# Combine color substring with file output
go run ./cmd --output=rainbow.txt --color=blue Art "ASCII Art" shadow
```

**Output Features:**

- Automatic file creation and overwriting
- ANSI color codes preserved in files
- Works with all banner styles
- Combine with color flags seamlessly

**File Output Notes:**

- Files contain ANSI escape codes for colors
- Use `cat` or `less -R` to view colors in terminal
- Plain text editors show raw ANSI codes
- Perfect for saving terminal art or banners

---

### 🔄 Reverse Feature

Convert ASCII art back to its original text with automatic banner detection.

<img src="assets/reverse_demo.gif" width="600" alt="Reverse Demo">

**Features:**

- ✨ Automatic banner detection (standard, shadow, thinkertoy)
- 📝 Multiline text support
- 🔢 Numbers and special characters
- 🔄 Variable-width character handling
- 🖥️ Windows/Unix line ending support

**Basic Usage:**

```bash
# Step 1: Generate ASCII art
go run ./cmd "Hello World" standard > output.txt

# Step 2: Reverse it back to text
go run ./cmd --reverse=output.txt
# Output: Hello World
```

**How it Works:**

1. Parses ASCII art into 8-line character chunks
2. Loads banner character templates
3. Matches patterns against all available banners
4. Returns the recognized text with preserved formatting

**Advanced Examples:**

```bash
# Works with all banner types (auto-detected)
go run ./cmd "Shadow Text" shadow > shadow.txt
go run ./cmd --reverse=shadow.txt
# Output: Shadow Text

# Multiline text
go run ./cmd "Line1\nLine2" standard > multi.txt
go run ./cmd --reverse=multi.txt
# Output: Line1
#         Line2

# Special characters and numbers
go run ./cmd "2024!" thinkertoy > year.txt
go run ./cmd --reverse=year.txt
# Output: 2024!
```

**Reverse Feature Notes:**

- Automatically detects which banner was used
- Supports standard, shadow, and thinkertoy banners
- Cannot reverse colored ASCII art (ANSI codes interfere with pattern matching)
- Preserves newlines and formatting

---

### 📐 Text Alignment & Justification

Align your ASCII art perfectly for any terminal width with dynamic text alignment.

<img src="assets/justify_demo.gif" width="900" alt="Justify Demo">

**Alignment Options:**

- `left` - Left alignment with 8-space margin (default)
- `right` - Right-aligned with dynamic padding
- `center` - Centered text with balanced spacing
- `justify` - Words distributed evenly across terminal width

**Basic Alignment:**

```bash
# Left alignment (default)
go run ./cmd --align=left "Hello" standard

# Right alignment
go run ./cmd --align=right "Hello" standard

# Center alignment
go run ./cmd --align=center "Hello" shadow

# Justify (word distribution)
go run ./cmd --align=justify "Hello World" thinkertoy
```

**Alignment Features:**

- **Terminal Width Detection**: Automatically adapts to your terminal size via `COLUMNS` environment variable
- **Dynamic Spacing**: Smart algorithms calculate optimal spacing for each alignment type
- **Works with All Banners**: Compatible with standard, shadow, and thinkertoy
- **Combine with Other Features**: Use alignment with colors and output flags

**Advanced Examples:**

```bash
# Center alignment with color
go run ./cmd --align=center --color=cyan "Centered" shadow

# Right alignment saved to file
go run ./cmd --align=right --output=right.txt "Right" standard

# Justify with substring coloring
go run ./cmd --align=justify --color=green "World" "Hello World" thinkertoy
```

**Alignment Notes:**

- Alignment applies to terminal output only (not file output)
- Terminal width is detected automatically (default: 80 columns if not detected)
- Justify distributes words evenly, creating uniform spacing
- All alignments maintain the integrity of ASCII art characters

---

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/IbsYoussef/Ascii-Art.git
cd Ascii-Art

# Test it works
go run ./cmd "Hello World" standard
```

### Basic Usage

```bash
# Simple ASCII art
go run ./cmd "Your Text"

# With banner choice
go run ./cmd "Your Text" <banner>

# With color
go run ./cmd --color=<color> "Your Text" <banner>

# Save to file
go run ./cmd --output=<file.txt> "Your Text" <banner>

# Reverse ASCII art
go run ./cmd --reverse=

# All together (except reverse)
go run ./cmd --output=art.txt --color=red "Your Text" shadow
```

**Banner Options:** `standard` (default), `shadow`, `thinkertoy`

---

## 📖 Usage Guide

### Command Syntax

```bash
go run ./cmd [OPTIONS] [STRING] [BANNER]
```

**Options:**

- `--color=<color>` - Apply color to text
- `--color=<color> <substring>` - Color specific substring
- `--output=<filename>` - Save output to file
- `--reverse=<filename>` - Convert ASCII art back to text

**Arguments:**

- `STRING` - Text to convert (required)
- `BANNER` - Banner style (optional, defaults to `standard`)

### Complete Examples

**Example 1: Basic ASCII Art**

```bash
go run ./cmd "Hello"
```

**Example 2: Choose Banner**

```bash
go run ./cmd "Shadow Style" shadow
go run ./cmd "Think Style" thinkertoy
```

**Example 3: Add Color**

```bash
# Named color
go run ./cmd --color=blue "Ocean" standard

# Hex color
go run ./cmd --color=#FF1493 "Pink" shadow

# RGB color
go run ./cmd --color='rgb(255,165,0)' "Orange" thinkertoy
```

**Example 4: Color Substring**

```bash
# Color "Go" in the phrase
go run ./cmd --color=cyan Go "Let's Go!" standard

# Color "World"
go run ./cmd --color=green World "Hello World" shadow
```

**Example 5: Save to File**

```bash
# Save plain ASCII art
go run ./cmd --output=output.txt "Save Me" standard

# Save colored ASCII art
go run ./cmd --output=colored.txt --color=red "Colored" shadow
```

**Example 6: Reverse ASCII Art**

```bash
# Generate and save ASCII art
go run ./cmd "Test" standard > test.txt

# Reverse it back
go run ./cmd --reverse=test.txt
# Output: Test

# Works with any banner (auto-detected)
go run ./cmd "Shadow" shadow > shadow.txt
go run ./cmd --reverse=shadow.txt
# Output: Shadow
```

**Example 7: Multi-line Text**

```bash
go run ./cmd "Line One\nLine Two" standard
go run ./cmd --color=blue "First\nSecond" shadow
```

**Example 8: Special Characters**

```bash
go run ./cmd "123!@#" standard
go run ./cmd --color=yellow "Numbers: 456" thinkertoy
```

**Example 9: Combining Features**

```bash
# Color + Output
go run ./cmd --output=art.txt --color=purple "Beautiful" shadow

# Substring Color + Output
go run ./cmd --output=highlight.txt --color=red World "Hello World" standard

# Multi-line + Color + Output
go run ./cmd --output=multi.txt --color=cyan "First\nSecond" thinkertoy
```

### Error Messages

The tool provides clear error messages for common issues:

```bash
# Missing input
go run ./cmd
# Output: Usage: go run ./cmd [STRING] [BANNER]

# Invalid banner
go run ./cmd "Test" invalid
# Output: invalid banner style
#         valid banners: standard, shadow, thinkertoy

# Invalid color format
go run ./cmd --color test "Hello"
# Output: invalid color format
#         Usage: go run ./cmd [OPTION] [STRING] [BANNER]

# Invalid output flag
go run ./cmd --output test.txt "Hello"
# Output: Usage: go run ./cmd [OPTION] [STRING] [BANNER]
#         EX: go run ./cmd --output= something standard

# Invalid reverse flag
go run ./cmd --reverse example.txt
# Output: Usage: go run ./cmd [OPTION]
#         EX: go run ./cmd --reverse=
```

---

## 📁 Project Structure

```
Ascii-Art/
├── README.md
├── go.mod
├── assets/
│   ├── intro_demo.gif          # Intro demo
│   ├── standard_demo.gif       # Standard features demo
│   ├── color_demo.gif          # Color features demo
│   ├── output_demo.gif         # Output features demo
│   ├── reverse_demo.gif        # Reverse features demo
│   └── justify_demo.gif        # Justify features demo
├── banners/
│   ├── standard.txt            # Standard banner font
│   ├── shadow.txt              # Shadow banner font
│   └── thinkertoy.txt          # Thinkertoy banner font
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── ascii/                  # Core ASCII logic
│   │   ├── input.go            # Input parsing & validation
│   │   ├── loadBanner.go       # Banner file loading
│   │   └── renderAscii.go      # ASCII art rendering
│   ├── ascii-color/            # Color feature module
│   │   ├── color.go            # Color parsing & ANSI codes
│   │   ├── inputColor.go       # Color flag parsing
│   │   └── renderColor.go      # Colored rendering logic
│   ├── ascii-output/           # Output feature module
│   │   ├── errors.go           # Error definitions
│   │   ├── inputOutput.go      # Output flag parsing
│   │   ├── fileWriter.go       # File writing logic
│   │   └── outputHandler.go    # Output routing & capture
│   ├── ascii-reverse/          # Reverse feature module
│   │   ├── errors.go           # Error definitions
│   │   ├── fileReader.go       # File reading utilities
│   │   ├── inputReverse.go     # Reverse flag parsing
│   │   ├── parser.go           # ASCII art parsing
│   │   ├── recogniser.go       # Pattern recognition
│   │   ├── templateLoader.go   # Banner template loading
│   │   └── reverseHandler.go   # Main reverse handler
│   ├── ascii-justify/          # Justify/align feature module
│   │   ├── align.go            # Alignment algorithms
│   │   ├── errors.go           # Error definitions
│   │   ├── handler.go          # Main justify handler
│   │   ├── inputJustify.go     # Align flag parsing
│   │   ├── measure.go          # Text measurement utilities
│   │   └── terminal.go         # Terminal width detection
│   └── files/
│       └── readFile.go         # File reading utilities
├── scripts/                    # Demo recording scripts
│   ├── demo_intro.sh           # Intro demo script
│   ├── demo_standard.sh        # Standard features demo
│   ├── demo_color.sh           # Color features demo
│   ├── demo_output.sh          # Output features demo
│   ├── demo_reverse.sh         # Reverse features demo
│   └── demo_justify.sh         # Justify features demo
└── test/
    ├── unit/                   # Unit tests
    │   ├── align_test.go
    │   ├── color_test.go
    │   ├── fileReader_test.go
    │   ├── fileWriter_test.go
    │   ├── inputColor_test.go
    │   ├── inputJustify_test.go
    │   ├── inputOutput_test.go
    │   ├── inputReverse_test.go
    │   ├── input_test.go
    │   ├── loadBanner_test.go
    │   ├── measure_test.go
    │   ├── outputHandler_test.go
    │   ├── parser_test.go
    │   ├── readFile_test.go
    │   ├── recogniser_test.go
    │   ├── renderAscii_test.go
    │   ├── renderColor_test.go
    │   ├── templateLoader_test.go
    │   ├── terminal_test.go
    │   └── test_helpers.go
    └── e2e/                    # End-to-end tests
        ├── e2e_test.go
        ├── e2e_color_test.go
        ├── e2e_output_test.go
        ├── e2e_reverse_test.go
        └── e2e_justify.sh      # Shell test script for justify
```

### Architecture Overview

**Modular Design:**

- `internal/ascii/` - Core ASCII art generation
- `internal/ascii-color/` - Color feature expansion
- `internal/ascii-output/` - File output feature
- `internal/ascii-reverse/` - Reverse text recognition
- `internal/files/` - Shared file utilities

**Key Design Principles:**

- Clean separation of concerns
- Feature isolation in modules
- Comprehensive error handling
- 100% backwards compatibility
- Extensive test coverage

---

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run unit tests only
go test -v ./test/unit/

# Run E2E tests only
go test -v ./test/e2e/

# Run specific feature tests
go test -v ./test/unit/recogniser_test.go      # Reverse feature
go test -v ./test/e2e/e2e_reverse_test.go      # Reverse E2E

# Run with coverage
go test -v -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Coverage

**Unit Tests:**

- ✅ Input parsing and validation
- ✅ Banner file loading
- ✅ ASCII art rendering
- ✅ Color parsing (named, hex, RGB, HSL)
- ✅ Color flag parsing and validation
- ✅ Colored rendering logic
- ✅ Output flag parsing
- ✅ File writing operations
- ✅ Output capture and routing
- ✅ Reverse flag parsing
- ✅ ASCII art file reading
- ✅ ASCII art parsing into chunks
- ✅ Pattern recognition and matching
- ✅ Template loading with CRLF support

**E2E Tests:**

- ✅ CLI command execution
- ✅ Standard ASCII art generation
- ✅ All banner styles
- ✅ Color functionality (all formats)
- ✅ Substring coloring
- ✅ File output creation
- ✅ Combined features (color + output)
- ✅ Reverse feature (13 comprehensive tests)
- ✅ Auto banner detection
- ✅ Multiline reverse
- ✅ Error handling and usage messages

**Test Statistics:**

- Total test files: 25 (20 unit + 5 E2E)
- Unit tests: 20 comprehensive test files
- Reverse feature: 13/13 tests passing (100% ✅)
- Justify feature: 15/15 manual tests passing (100% ✅)
- Overall coverage: ~95%

---

## 🔭 Roadmap

### Current Version: v1.4.0

**Completed Features:**

- ✅ v1.0.0 - Core ASCII art generation
- ✅ v1.1.0 - Color support (named, hex, RGB, HSL)
- ✅ v1.2.0 - Output to file support
- ✅ v1.3.0 - Reverse feature (ASCII art → text)
- ✅ v1.4.0 - Text alignment and justification

### Future Enhancements

**Under Consideration:**

- 🎨 Additional banner styles
- ⚡ Performance optimizations
- 🔧 Extended special character support
- 📊 ASCII art templates and presets

---

## 🙏 Acknowledgements

- Created as part of my Go learning journey at **01 Founders**
- Inspired by classic ASCII art and terminal aesthetics
- Thank you to the Go community for excellent documentation and tools

---

## 📄 License

This project is licensed under the **MIT License**.

```
MIT License

Copyright (c) 2026 IbsYoussef

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

See [LICENSE.txt](LICENSE.txt) for full details.

---

<div align="center">

[⬆ Back to Top](#ascii-art)

</div>
