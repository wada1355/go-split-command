package splitters_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wata1355/go-split-command/splitters"
)

func TestSplitByLines(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		fileContent   string
		linesPerFile  int
		expectedFiles map[string]string
	}{
		{
			name:         "Basic split",
			fileName:     "testByLines_basic.txt",
			fileContent:  "line1\nline2\nline3\nline4\nline5\n",
			linesPerFile: 2,
			expectedFiles: map[string]string{
				"testByLines_basic_part1.txt": "line1\nline2\n",
				"testByLines_basic_part2.txt": "line3\nline4\n",
				"testByLines_basic_part3.txt": "line5\n",
			},
		},
		{
			name:         "Lines less than linesPerFile",
			fileName:     "testByLines_less.txt",
			fileContent:  "line1\n",
			linesPerFile: 2,
			expectedFiles: map[string]string{
				"testByLines_less_part1.txt": "line1\n",
			},
		},
		{
			name:         "Lines equal to linesPerFile",
			fileName:     "testByLines_equal.txt",
			fileContent:  "line1\nline2\n",
			linesPerFile: 2,
			expectedFiles: map[string]string{
				"testByLines_equal_part1.txt": "line1\nline2\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFilePath := filepath.Join(dirName, tt.fileName)
			if err := os.WriteFile(testFilePath, []byte(tt.fileContent), 0644); err != nil {
				t.Fatalf("Failed to create test file. err is %v", err)
			}
			if err := splitters.SplitByLines(testFilePath, tt.linesPerFile); err != nil {
				t.Fatalf("Failed to split by lines: %v", err)
			}
			for expectedFileName, expectedContent := range tt.expectedFiles {
				checkFileContent(t, filepath.Join(dirName, expectedFileName), expectedContent)
			}
		})
	}
}
