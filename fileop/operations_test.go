package fileop_test

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"

	"github.com/wata1355/go-split-command/fileop"
)

func TestWriteToFile(t *testing.T) {
	testFile, err := os.CreateTemp(dirName, "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tests := []struct {
		name    string
		writer  *bufio.Writer
		newFile *os.File
		isErr   bool
	}{
		{
			name:    "valid writer and file",
			writer:  bufio.NewWriter(testFile),
			newFile: testFile,
			isErr:   false,
		},
		{
			name:    "nil writer",
			writer:  nil,
			newFile: testFile,
			isErr:   false,
		},
		{
			name:    "nil file",
			writer:  bufio.NewWriter(testFile),
			newFile: nil,
			isErr:   false,
		},
		{
			name:    "both nil",
			writer:  nil,
			newFile: nil,
			isErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fileop.WriteToFile(tt.writer, tt.newFile)
			if (err != nil) != tt.isErr {
				t.Errorf("isErr %v but error: %v", tt.isErr, err)
			}
		})
	}
}

func TestCreateNewFile(t *testing.T) {
	tests := []struct {
		name         string
		fileName     string
		newFileNum   int
		isErr        bool
		wantFilePath string
	}{
		{
			name:         "basic test",
			fileName:     "testdata/testfile.txt",
			newFileNum:   1,
			isErr:        false,
			wantFilePath: "testdata/testfile_part1.txt",
		},
		{
			name:         "invalid file path",
			fileName:     "",
			newFileNum:   1,
			isErr:        true,
			wantFilePath: "",
		},
		{
			name:         "invalid file number",
			fileName:     "",
			newFileNum:   0,
			isErr:        true,
			wantFilePath: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := fileop.CreateNewFile(test.fileName, test.newFileNum)
			if err == nil && test.isErr {
				t.Error("Expected error, got nil")
			}
			if err != nil && !test.isErr {
				t.Errorf("unexpected error: %v", err)
			}
			if test.wantFilePath != "" {
				if _, err := os.Stat(test.wantFilePath); os.IsNotExist(err) {
					t.Errorf("%s does not exist", test.wantFilePath)
				}
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		name        string
		fileName    string
		fileContent string
		expected    int
	}{
		{
			name:        "normal case",
			fileName:    "testfile.txt",
			fileContent: "line1\nline2\nline3",
			expected:    3,
		},
		{
			name:        "empty file",
			fileName:    "emptyFile.txt",
			fileContent: "",
			expected:    0,
		},
		{
			name:        "single line",
			fileName:    "singleLine.txt",
			fileContent: "line1",
			expected:    1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filePath := filepath.Join(dirName, test.fileName)
			if err := os.WriteFile(filePath, []byte(test.fileContent), 0644); err != nil {
				t.Fatalf("Failed to create test file. err is %v", err)
			}
			lineCount, err := fileop.CountLines(filePath)
			if err != nil {
				t.Errorf("Failed to get lines. err is %v", err)
			}
			if lineCount != test.expected {
				t.Errorf("Expected line count %d, got %d", test.expected, lineCount)
			}
		})
	}
}
