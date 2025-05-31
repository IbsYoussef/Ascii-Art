package tests

import (
	"ascii-art/util"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantData  util.InputData
		wantError bool
	}{
		{
			name:      "No arguments",
			args:      []string{},
			wantError: true,
		},
		{
			name:      "Empty text input",
			args:      []string{""},
			wantError: true,
		},
		{
			name: "Valid input with default banner",
			args: []string{"Hello"},
			wantData: util.InputData{
				Text:   "Hello",
				Banner: "standard",
			},
		},
		{
			name: "Valid input with custom banner",
			args: []string{"World", "shadow"},
			wantData: util.InputData{
				Text:   "World",
				Banner: "shadow",
			},
		},
		{
			name: "Newline handling",
			args: []string{"Line1\\nLine2"},
			wantData: util.InputData{
				Text:   "Line1\nLine2",
				Banner: "standard",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := util.ParseInput(tt.args)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if gotData != tt.wantData {
				t.Errorf("expected %+v, got %+v", tt.wantData, gotData)
			}
		})
	}
}
