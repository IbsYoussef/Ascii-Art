# ASCII-ART

A CLI application written in Go that renders ASCII art using formatted banners from input text.

## 📑 Table of Contents

1. [📝 About](#-about)  
2. [📁 File Structure](#-file-structure)  
3. [✨ Features](#-features)  
4. [🚀 Usage Instructions](#-usage-instructions)  
   - [📦 Clone the Repository](#-clone-the-repository)  
   - [⚙️ Run the CLI (Go Users)](#-run-the-cli-go-users)  
   - [📥 Download the Executables (Non-Go Users)](#-download-the-executables-non-go-users)  
5. [🔭 Future Plans](#-future-plans)  
6. [🤝 Contributions](#-contributions)  
7. [🙏 Acknowledgements](#-acknowledgements)
7. [📄 License](#-license)

---

## 📝 About
```
                               _____    _____   _____   _____                    _____    _______                       
                      /\      / ____|  / ____| |_   _| |_   _|           /\     |  __ \  |__   __|                      
                     /  \    | (___   | |        | |     | |            /  \    | |__) |    | |                         
                    / /\ \    \___ \  | |        | |     | |           / /\ \   |  _  /     | |                         
                   / ____ \   ____) | | |____   _| |_   _| |_         / ____ \  | | \ \     | |                         
                  /_/    \_\ |_____/   \_____| |_____| |_____|       /_/    \_\ |_|  \_\    |_|                         
```

This project is a command-line tool written in Go that takes user input and transforms it into styled ASCII art using various banner template styles like `standard`, `shadow`, and `thinkertoy`.

It was created as part of my coding course at 01 Founders to deepen my understanding of Go, file parsing, and modular programming — all while keeping things fun and creative. The tool handles multi-line input, clean formatting, and is structured for easy extension and maintenance.

---

## 📁 File structure
```
.
├── README.md
├── ROADMAP.md
├── banners // Banner files used for reading, parsing,and rendering ASCII art text
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── build // Builder that cross-compiles and zips output
│   ├── embedmain
│   │   └── main.go
│   └── main.go
├── cmd // Main command point to run program
│   └── main.go
├── go.mod
├── internal
│   ├── ascii // Core ascii logic
│   │   ├── input.go
│   │   ├── input_test.go
│   │   ├── loadBanner.go
│   │   ├── loadBanner_test.go
│   │   ├── renderAscii.go
│   │   └── renderAscii_test.go
│   ├── asciiembed // Embedded banner logic
│   │   ├── banners
│   │   │   ├── shadow.txt
│   │   │   ├── standard.txt
│   │   │   └── thinkertoy.txt
│   │   ├── banners.go
│   │   ├── loadBanner.go
│   │   └── parser.go
│   ├── e2e // End to End integration test
│   │   └── e2e_test.go
│   └── files // File reading and helpers
│       ├── readFile.go
│       └── readFile_test.go
└── output_tests // ASCII output test samples
    ├── output.txt
    ├── output2.txt
    ├── output3.txt
    ├── output4_shadow.txt
    ├── output5_thinkertoy.txt
    ├── output6_thinkertoy.txt
    └── output7.txt

12 directories, 31 files
```
---

## ✨ Features
- 🖊️ Takes user input from the command line and renders it as ASCII art  
- 🎨 Supports multiple banner styles (`standard`, `shadow`, `thinkertoy`)  
- 📜 Handles multi-line input using `\n` escape characters  
- 💡 Clean and modular Go codebase for easy readability and testing  
- 🧩 Embeds banner files into the binary for portability  
- ⚙️ Cross-platform builder with auto `.zip` packaging support  
---

## 🚀 Usage Instructions
- ### 📦 Clone the repository
First, clone the repository to your local machine:

```bash
git clone https://github.com/IbsYoussef/Ascii-Art.git
cd ascii-art
```

- ### ⚙️ Run the CLI (Go Users)
```bash
go run ./cmd "Hello World" <banner-choice>
```
You can use either `standard`, `shadow` or `thinkertoy` as you banner choice for the styling, if omitted the standard banner will be used by default.

To print text on multiple lines use \n in your string input:
```bash
go run ./cmd "Hello\nWorld" <banner-choice>
```

- ### 📥 Download the executables (Non-Go Users)
Precompiled .zip packages are available for each platform in the Releases tab.
Example .zip packages:
- ascii-art-windows.zip
- ascii-art-linux.zip
- ascii-art-macos-intel.zip
- ascii-art-macos-arm64.zip

Each .zip contains a single, ready-to-run binary.
Just extract it and run from your terminal:

```bash
./ascii-art-linux "Hello\nWorld" shadow
```

---

## 🔭 Future Plans
Here are a few enhancements I plan to add in future updates:

- 🎨 **Color Output**: Add a `--color` flag so users can stylize their ASCII art with terminal color codes (e.g., red, green, cyan, etc.)
- 🔁 **Reverse Mode**: Option to reverse the input text before rendering it in ASCII format
- 💾 **Output to File**: Allow users to save the ASCII art output to a file of their choice using an `--output` or `-o` flag
- 📐 **Text Alignment**: Add flags for aligning text output (`--left`, `--center`, `--right`) for better formatting control
- 🌐 **Web Version**: Add a hosted web version for users to interact with directly
- 🛠 **CLI flag support** in builder (e.g., --os linux, --zip, --clean)
- 📦 **Auto-zipping** in the build step for smoother releases
- 🤖 **Optional GitHub/Gitea CI integration** to automate building and releasing

---

##  🤝 Contributions
Contributions are welcome! If you'd like to help improve **ascii-art**, please follow these steps:

1. **Fork the Repository:**  
   Click the "Fork" button at the top-right of the repository page to create your own copy of the project.

2. **Create a New Branch:**  
   Create a new branch for your feature or bug fix:
   ```bash
    git checkout -b feature-or-bugfix-description
   ```
3. **Make your Changes:**
Implement your changes and ensure that your code adheres to the project's style guidelines.
Tip: Write or update tests as needed.

4. **Commit and Push your Changes**:
Commit your changes with a clear, descriptive message and push your branch to your forked repository:
    ```bash
    git commit -m "Add: description of your changes"
    git push origin feature-or-bugfix-description
    ```
5. **Open a Pull Request**:
Open a pull request (PR) from your branch to the main repository. Please include a clear description of your changes and the motivation behind them.
If you're not sure about a major change, open an issue first to discuss your ideas.

Thank you for helping make ascii-art even better!

---
## 🙏 Acknowledgements
- Created as part of my Go learning journey at 01 Founders
---

## 📄 License
This project is licensed under the [MIT License](LICENSE).

Acknowledgements
Special Thanks:
Thanks to all contributors, mentors, and peers who provided feedback and support during the development of go-reloaded.

Inspiration:
This project was inspired by best practices in Go development and the need for automated text formatting solutions.

Resources:

The MIT License
Various open-source projects and communities that encourage collaboration and learning.
Thank you for checking out go-reloaded! We hope this tool helps streamline your text processing tasks and that you find it both useful and easy to contribute to.


