# Ascii art Task list

---
### 🟩 Phase 1: Setup & Input Handling
- [x] Create the project structure (`main.go`, `parser/`, `renderer/`, `banners/`, `tests/`, `experiment.go`)
- [x] Read and validate command-line arguments
- [x] Convert literal `\n` in input string to actual newlines
- [x] Add support for selecting different banners via optional second argument (currently hardcoded to `"standard"`)

---

### 🟨 Phase 2: Banner File Parsing
- [ ] Load the correct banner file (e.g., `standard.txt`) from the `banners/` directory
- [ ] Parse each 8-line character block into a map like `map[rune][]string`
- [ ] Verify correctness by printing a known character (e.g. “A”) using the parsed data

---

### 🟧 Phase 3: Rendering Logic
- [ ] Loop through each character in the input string
- [ ] For each row (0–7), concatenate that row of each character to form complete lines
- [ ] Handle line breaks (`\n`) correctly to output multi-line ASCII art

---

### 🟥 Phase 4: Output & Display
- [ ] Print the final ASCII art result to the terminal
- [ ] Handle empty lines (multiple `\n`) correctly
- [ ] Decide how to handle unsupported characters (e.g., error, skip, or blank)

---

### 🧪 Phase 5: Testing
- [ ] Create unit tests for the `ParseInput` function
- [ ] Create unit tests for banner parsing logic
- [ ] Create unit tests for rendering logic
- [ ] Test special cases like empty string, just `\n`, unsupported characters, etc.

---