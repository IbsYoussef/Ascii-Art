# ASCII-ART

<div align="center">

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

![Demo](assets/demo.gif)

**Generate beautiful colored ASCII art from text with multiple banner styles**

</div>

## 📋 Table of Contents

1. [🎯 About](#-about)
2. [📁 File Structure](#-file-structure)
3. [✨ Features](#-features)
4. [🚀 Usage Instructions](#-usage-instructions)
5. [🎨 Color Support](#-color-support)
6. [📚 Examples](#-examples)
7. [🧪 Testing](#-testing)
8. [🙏 Acknowledgements](#-acknowledgements)
9. [📄 License](#-license)

---

## 🎯 About

**ASCII Art Generator** transforms text into stylized ASCII art using multiple banner templates with full color support. Built in Go as part of the **01 Founders** curriculum, this project demonstrates:

- Clean modular architecture
- File parsing and rendering techniques
- Advanced color processing (RGB, HSL, Hex)
- Comprehensive unit and E2E testing
- Command-line interface design

The tool handles multi-line input, offers three distinct banner styles, supports multiple color formats, and is structured for easy extension — making it both practical and fun to use.

---

## 📁 File Structure

```
.
├── README.md
├── banners/                # ASCII art font files
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── cmd/
│   └── main.go             # Entry point
├── internal/
│   ├── ascii/              # Core ASCII logic
│   │   ├── input.go
│   │   ├── loadBanner.go
│   │   └── renderAscii.go
│   ├── ascii-color/        # Color expansion
│   │   ├── color.go
│   │   ├── inputColor.go
│   │   └── renderColor.go
│   └── files/              # File utilities
│       └── readFile.go
├── test/
│   ├── unit/               # Unit tests
│   └── e2e/                # End-to-end tests
└── assets/
    └── demo.gif            # Demo recording
```

---

## ✨ Features

### 🖊️ **ASCII Art Generation**

- Takes user input from the command line and renders it as ASCII art
- Supports multiple banner styles (`standard`, `shadow`, `thinkertoy`)
- Handles multi-line input using `\n` escape characters

### 🎨 **Color Support**

- Color entire strings or specific substrings
- Multiple color format support:
  - **Named colors**: red, blue, green, yellow, orange, cyan, magenta, white, black
  - **Hex colors**: `#FF0000`, `#F00` (short form)
  - **RGB colors**: `rgb(255,0,0)`
  - **HSL colors**: `hsl(0,100%,50%)`
- Case-sensitive substring matching
- True RGB color accuracy

### 💡 **Code Quality**

- Clean and modular Go codebase
- Comprehensive test coverage
- Backwards compatible (works without color flags)

---

## 🚀 Usage Instructions

### 📦 Clone the Repository

```bash
git clone https://learn.01founders.co/git/iyoussef/Ascii-Art-Color-V2.git
cd Ascii-Art-Color-V2
```

### ⚙️ Basic Usage

```bash
# Without color
go run ./cmd "Hello World" <banner-choice>

# With color (entire string)
go run ./cmd --color=<color> "Hello World" <banner-choice>

# With color (substring only)
go run ./cmd --color=<color> <substring> "Hello World" <banner-choice>
```

---

## 🎨 Color Support

### Syntax

```bash
go run ./cmd --color=<color> [substring] "text" [banner]
```

### Color Formats

**Named Colors:**

```bash
go run ./cmd --color=red "Hello"
go run ./cmd --color=blue "World"
```

**Hex Colors:**

```bash
go run ./cmd --color=#FF0000 "Red Text"
go run ./cmd --color=#0F0 "Green Text"
```

**RGB Colors:**

```bash
go run ./cmd --color=rgb(255,0,0) "Red"
go run ./cmd --color=rgb(0,255,0) "Green"
```

**HSL Colors:**

```bash
go run ./cmd --color=hsl(0,100%,50%) "Red"
go run ./cmd --color=hsl(120,100%,50%) "Green"
```

### Substring Coloring

Color only specific parts of your text:

```bash
# Color "kit" in "kitten"
go run ./cmd --color=blue kit "a king kitten have kit"

# Color "B" in "RGB()"
go run ./cmd --color=red B "RGB()"
```

**Note:** Substring matching is case-sensitive!

---

## 📚 Examples

### Example 1: Basic ASCII Art

```bash
go run ./cmd "Hello"
```

### Example 2: Colored Text

```bash
go run ./cmd --color=red "Hello World"
```

### Example 3: Substring Coloring

```bash
go run ./cmd --color=green kit "a king kitten"
```

### Example 4: With Different Banner

```bash
go run ./cmd --color=cyan "Hello" shadow
```

### Example 5: Multiple Color Formats

```bash
# Named
go run ./cmd --color=orange "Test"

# Hex
go run ./cmd --color=#FF00FF "Test"

# RGB
go run ./cmd --color=rgb(0,255,255) "Test"

# HSL
go run ./cmd --color=hsl(240,100%,50%) "Test"
```

---

## 🧪 Testing

### Run All Tests

```bash
# Unit tests
go test -v ./test/unit/

# E2E tests
go test -v ./test/e2e/

# All tests
go test -v ./...

# With coverage
go test -v -cover ./...
```

### Test Coverage

- ✅ Color parsing (named, hex, RGB, HSL)
- ✅ Input validation and flag parsing
- ✅ Substring matching and coloring
- ✅ Banner loading and rendering
- ✅ End-to-end CLI testing

---

## 🙏 Acknowledgements

- Created as part of my Go learning journey at 01 Founders
- Inspired by classic ASCII art and terminal aesthetics
- Thank you to the Go Community for excellent documentation

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

**[⬆ Back to Top](#ascii-art)**

</div>
