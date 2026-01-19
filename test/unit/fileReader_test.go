package unit

import (
	output "ascii-art/internal/ascii-reverse"
	"os"
	"path/filepath"
	"testing"
)

func TestReadAsciiArtFile(t *testing.T) {
	// Create temporary directory for test files
	tempDir := t.TempDir()

	tests := []struct {
		name        string
		filename    string
		content     string
		setup       func(string) error
		wantContent string
		wantErr     bool
		errContains string
	}{
		{
			name:     "read simple ASCII art",
			filename: filepath.Join(tempDir, "simple.txt"),
			content: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
			setup: func(path string) error {
				return os.WriteFile(path, []byte(` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`), 0644)
			},
			wantContent: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
			wantErr: false,
		},
		{
			name:     "read multiline ASCII art",
			filename: filepath.Join(tempDir, "multi.txt"),
			content: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
 _   _   _                 _       _  
| | | | | |               | |     | | 
| |_| | | |   ___    ___  | |   __| | 
| | | | | |  / _ \  / __| | |  / _  | 
| | | | | | | (_) | |     | | | (_| | 
|_| |_| |_|  \___/  |     |_|  \__._| 
                                      
                                      
`,
			setup: func(path string) error {
				return os.WriteFile(path, []byte(` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
 _   _   _                 _       _  
| | | | | |               | |     | | 
| |_| | | |   ___    ___  | |   __| | 
| | | | | |  / _ \  / __| | |  / _  | 
| | | | | | | (_) | |     | | | (_| | 
|_| |_| |_|  \___/  |     |_|  \__._| 
                                      
                                      
`), 0644)
			},
			wantContent: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
 _   _   _                 _       _  
| | | | | |               | |     | | 
| |_| | | |   ___    ___  | |   __| | 
| | | | | |  / _ \  / __| | |  / _  | 
| | | | | | | (_) | |     | | | (_| | 
|_| |_| |_|  \___/  |     |_|  \__._| 
                                      
                                      
`,
			wantErr: false,
		},
		{
			name:     "read empty file",
			filename: filepath.Join(tempDir, "empty.txt"),
			content:  "",
			setup: func(path string) error {
				return os.WriteFile(path, []byte(""), 0644)
			},
			wantContent: "",
			wantErr:     false,
		},
		{
			name:        "file does not exist",
			filename:    filepath.Join(tempDir, "nonexistent.txt"),
			setup:       nil,
			wantContent: "",
			wantErr:     true,
			errContains: "file not found",
		},
		{
			name:     "file with special characters",
			filename: filepath.Join(tempDir, "special.txt"),
			content:  "  _   ____    _____  \n / | |___ \\  |___ /  \n| |   __) |   |_ \\  \n| |  / __/   ___) | \n| | |_____| |____/  \n|_|                 \n",
			setup: func(path string) error {
				return os.WriteFile(path, []byte("  _   ____    _____  \n / | |___ \\  |___ /  \n| |   __) |   |_ \\  \n| |  / __/   ___) | \n| | |_____| |____/  \n|_|                 \n"), 0644)
			},
			wantContent: "  _   ____    _____  \n / | |___ \\  |___ /  \n| |   __) |   |_ \\  \n| |  / __/   ___) | \n| | |_____| |____/  \n|_|                 \n",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: Create file if needed
			if tt.setup != nil {
				err := tt.setup(tt.filename)
				if err != nil {
					t.Fatalf("Setup failed: %v", err)
				}
			}

			// Execute
			got, err := output.ReadAsciiArtFile(tt.filename)

			// Check error expectations
			if tt.wantErr {
				if err == nil {
					t.Errorf("ReadAsciiArtFile() expected error, got nil")
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("ReadAsciiArtFile() error = %v, want error containing %q", err, tt.errContains)
				}
				return
			}

			// No error expected
			if err != nil {
				t.Errorf("ReadAsciiArtFile() unexpected error = %v", err)
				return
			}

			// Verify content
			if got != tt.wantContent {
				t.Errorf("ReadAsciiArtFile() content mismatch\ngot:\n%q\nwant:\n%q", got, tt.wantContent)
			}
		})
	}
}
