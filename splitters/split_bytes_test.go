package splitters_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wata1355/go-split-command/splitters"
)

func TestSplitByBytes(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		fileContent   string
		bytesPerFile  int
		expectedFiles map[string]string
	}{
		{
			name:         "Basic split",
			fileName:     "testByBytes_basic.txt",
			fileContent:  "Hello, World!",
			bytesPerFile: 5,
			expectedFiles: map[string]string{
				"testByBytes_basic_part1.txt": "Hello",
				"testByBytes_basic_part2.txt": ", Wor",
				"testByBytes_basic_part3.txt": "ld!",
			},
		},
		{
			name:         "Content less than bytesPerFile",
			fileName:     "testByBytes_less.txt",
			fileContent:  "Hello",
			bytesPerFile: 10,
			expectedFiles: map[string]string{
				"testByBytes_less_part1.txt": "Hello",
			},
		},
		{
			name:         "Content equal to bytesPerFile",
			fileName:     "testByBytes_equal.txt",
			fileContent:  "Hello",
			bytesPerFile: 5,
			expectedFiles: map[string]string{
				"testByBytes_equal_part1.txt": "Hello",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFilePath := filepath.Join(dirName, tt.fileName)
			if err := os.WriteFile(testFilePath, []byte(tt.fileContent), 0644); err != nil {
				t.Fatalf("Failed to create test file. err is %v", err)
			}

			if err := splitters.SplitByBytes(testFilePath, tt.bytesPerFile); err != nil {
				t.Fatalf("Failed to split by bytes: %v", err)
			}

			for expectedFileName, expectedContent := range tt.expectedFiles {
				checkFileContent(t, filepath.Join(dirName, expectedFileName), expectedContent)
			}
		})
	}
}
