# ASCII Art Project Roadmap

A Go program that renders input strings as stylised ASCII art using different banner styles.

---

## 📌 Phase 1: Core CLI + Argument Parsing
> **Goal:** Handle input from the user and validate it.

- [ ] Accept a string input via command-line arguments
- [ ] Parse and validate input (ensure only printable characters, handle `\n`)
- [ ] Accept and validate banner style input (`shadow`, `standard`, `thinkertoy`)
- [ ] Print helpful error messages for invaid input or style

---

## 🧩 Phase 2: Banner File Loader
> **Goal:** Load and parse the font/banner data.

- [ ] Open the appropriate banner file from `/banners`
- [ ] Read and store ASCII representations into a usable data structure (e.g., `map[rune][]string`)
- [ ] Handle edge cases (e.g., missing files, malformed banners)

---

## 🖼 Phase 3: ASCII Rendering Engine
> **Goal:** Translate input strings into ASCII art using loaded banners.

- [ ] Implement render logic (match characters to font lines)
- [ ] Handle line breaks (`\n`) properly
- [ ] Optimize spacing and alignment
- [ ] Add support for special characters

---

## 🧪 Phase 4: Testing & Validation
> **Goal:** Ensure reliability through tests.

- [ ] Write unit tests for banner loading
- [ ] Write unit tests for render logic
- [ ] Test CLI behavior with various inputs
- [ ] Validate output against known-good samples

---
