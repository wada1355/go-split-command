package validators_test

import (
	"testing"

	"github.com/wata1355/go-split-command/splitters"
	"github.com/wata1355/go-split-command/validators"
)

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		name      string
		fileArgs  []string
		options   splitters.Options
		expectErr bool
	}{
		{
			name:      "valid arguments",
			fileArgs:  []string{"file.txt"},
			options:   splitters.Options{Lines: 10},
			expectErr: false,
		},
		{
			name:      "no file argument",
			fileArgs:  []string{},
			options:   splitters.Options{Lines: 10},
			expectErr: true,
		},
		{
			name:      "multiple file arguments",
			fileArgs:  []string{"file1.txt", "file2.txt"},
			options:   splitters.Options{Lines: 10},
			expectErr: true,
		},
		{
			name:      "nonexistent file",
			fileArgs:  []string{"nonexistent.txt"},
			options:   splitters.Options{Lines: 10},
			expectErr: true,
		},
		{
			name:      "empty file",
			fileArgs:  []string{"emptyFile.txt"},
			options:   splitters.Options{Lines: 10},
			expectErr: true,
		},
		{
			name:      "no split option",
			fileArgs:  []string{"file.txt"},
			options:   splitters.Options{},
			expectErr: false,
		},
		{
			name:      "multiple split options",
			fileArgs:  []string{"file.txt"},
			options:   splitters.Options{Lines: 10, Bytes: 10},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validators.ValidateArgs(tt.fileArgs, &tt.options)
			if tt.expectErr && err == nil {
				t.Error("Expected error, got nil")
			} else if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
