# ASCII Art Project Roadmap

A Go program that renders input strings as stylised ASCII art using different banner styles.

---

## 📌 Phase 1: Core CLI + Argument Parsing
> **Goal:** Handle input from the user and validate it.

- [x] Accept a string input via command-line arguments
- [x] Parse and validate input (ensure only printable characters, handle `\n`)
- [x] Accept and validate banner style input (`shadow`, `standard`, `thinkertoy`)
- [x] Print helpful error messages for invaid input or style

---

## 🧩 Phase 2: Banner File Loader
> **Goal:** Load and parse the font/banner data.

- [x] Open the appropriate banner file from `/banners`
- [x] Read and store ASCII representations into a usable data structure (e.g., `map[rune][]string`)
- [x] Handle edge cases (e.g., missing files, malformed banners)

---

## 🖼 Phase 3: ASCII Rendering Engine
> **Goal:** Translate input strings into ASCII art using loaded banners.

- [x] Implement render logic (match characters to font lines)
- [x] Handle line breaks (`\n`) properly
- [x] Optimize spacing and alignment

---

## 🧪 Phase 4: Testing & Validation
> **Goal:** Ensure reliability through tests.

- [ ] Write unit tests for banner loading
- [ ] Write unit tests for render logic
- [ ] Test CLI behavior with various inputs
- [ ] Validate output against known-good samples

---
