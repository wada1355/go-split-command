package splitters_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wata1355/go-split-command/splitters"
)

func TestSplitByNumFiles(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		fileContent   string
		numFiles      int
		expectedFiles map[string]string
	}{
		{
			name:        "Divisible lines",
			fileName:    "testByNumFiles_divisible.txt",
			fileContent: "line1\nline2\nline3\nline4\n",
			numFiles:    2,
			expectedFiles: map[string]string{
				"testByNumFiles_divisible_part1.txt": "line1\nline2\n",
				"testByNumFiles_divisible_part2.txt": "line3\nline4\n",
			},
		},
		{
			name:        "Indivisible lines",
			fileName:    "testByNumFiles_indivisible.txt",
			fileContent: "line1\nline2\nline3\nline4\nline5\n",
			numFiles:    2,
			expectedFiles: map[string]string{
				"testByNumFiles_indivisible_part1.txt": "line1\nline2\n",
				"testByNumFiles_indivisible_part2.txt": "line3\nline4\n",
				"testByNumFiles_indivisible_part3.txt": "line5\n",
			},
		},
		{
			name:        "Less lines than files",
			fileName:    "testByNumFiles_less.txt",
			fileContent: "line1\nline2\n",
			numFiles:    3,
			expectedFiles: map[string]string{
				"testByNumFiles_less_part1.txt": "line1\nline2\n",
			},
		},
		{
			name:        "Lines equal to files",
			fileName:    "testByNumFiles_equal.txt",
			fileContent: "line1\nline2\n",
			numFiles:    2,
			expectedFiles: map[string]string{
				"testByNumFiles_equal_part1.txt": "line1\n",
				"testByNumFiles_equal_part2.txt": "line2\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFilePath := filepath.Join(dirName, tt.fileName)
			err := os.WriteFile(testFilePath, []byte(tt.fileContent), 0644)
			if err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}
			splitter := splitters.Splitter{
				Options:  splitters.Options{NumFiles: tt.numFiles},
				FileArgs: splitters.FileArgs{FilePath: testFilePath},
			}

			if err := splitter.SplitByNumFiles(); err != nil {
				t.Fatalf("Failed to split by bytes: %v", err)
			}
			for expectedFileName, expectedContent := range tt.expectedFiles {
				checkFileContent(t, filepath.Join(dirName, expectedFileName), expectedContent)
			}
		})
	}
}
