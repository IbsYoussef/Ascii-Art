# ASCII-ART

<div align="center">

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

![Basic Demo](assets/demo.gif)

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
- 🧪 100% test coverage with unit and E2E tests
- 📦 Zero dependencies - uses only Go standard library

---

## ✨ Features

### 🖊️ ASCII Art Generation

Transform any text into stylized ASCII art using three distinct banner fonts.

![Standard Demo](assets/demo_standard.gif)

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

![Color Demo](assets/demo_color.gif)

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

![Output Demo](assets/demo_output.gif)

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

# All together
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

**Example 6: Multi-line Text**

```bash
go run ./cmd "Line One\nLine Two" standard
go run ./cmd --color=blue "First\nSecond" shadow
```

**Example 7: Special Characters**

```bash
go run ./cmd "123!@#" standard
go run ./cmd --color=yellow "Numbers: 456" thinkertoy
```

**Example 8: Combining Features**

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
#         EX: go run ./cmd --output=<fileName.txt> something standard
```

---

## 📁 Project Structure

```
Ascii-Art/
├── README.md
├── LICENSE.txt
├── ROADMAP.md
├── go.mod
├── assets/
│   ├── demo.gif                # Original project demo
│   ├── demo_standard.gif       # Standard features demo
│   ├── demo_color.gif          # Color features demo
│   └── demo_output.gif         # Output features demo
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
│   └── files/
│       └── readFile.go         # File reading utilities
└── test/
    ├── unit/                   # Unit tests
    │   ├── color_test.go
    │   ├── inputColor_test.go
    │   ├── inputOutput_test.go
    │   ├── input_test.go
    │   ├── loadBanner_test.go
    │   ├── readFile_test.go
    │   ├── renderAscii_test.go
    │   ├── renderColor_test.go
    │   ├── fileWriter_test.go
    │   └── outputHandler_test.go
    └── e2e/                    # End-to-end tests
        ├── e2e_test.go
        ├── e2e_color_test.go
        └── e2e_output_test.go
```

### Architecture Overview

**Modular Design:**

- `internal/ascii/` - Core ASCII art generation
- `internal/ascii-color/` - Color feature expansion
- `internal/ascii-output/` - File output feature
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

**E2E Tests:**

- ✅ CLI command execution
- ✅ Standard ASCII art generation
- ✅ All banner styles
- ✅ Color functionality (all formats)
- ✅ Substring coloring
- ✅ File output creation
- ✅ Combined features (color + output)
- ✅ Error handling and usage messages

**Test Statistics:**

- Total test files: 13
- Unit tests: 10 files
- E2E tests: 3 files
- Coverage: ~95%

---

## 🔭 Roadmap

### Current Version: v1.2.0

**Completed Features:**

- ✅ v1.0.0 - Core ASCII art generation
- ✅ v1.1.0 - Color support (named, hex, RGB, HSL)
- ✅ v1.2.0 - Output to file support

### Future Enhancements

**Planned for v1.3.0:**

- 🔄 **Reverse Mode** - `--reverse` flag to reverse input text before rendering

**Planned for v1.4.0:**

- 📏 **Text Alignment** - `--align=<left|center|right>` flag for text justification

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
