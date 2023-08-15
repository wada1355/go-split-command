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
	defer os.Remove(testFile.Name())
	tests := []struct {
		name    string
		writer  *bufio.Writer
		newFile *os.File
		wantErr bool
	}{
		{
			name:    "valid writer and file",
			writer:  bufio.NewWriter(testFile),
			newFile: testFile,
			wantErr: false,
		},
		{
			name:    "nil writer",
			writer:  nil,
			newFile: testFile,
			wantErr: false,
		},
		{
			name:    "nil file",
			writer:  bufio.NewWriter(testFile),
			newFile: nil,
			wantErr: false,
		},
		{
			name:    "both nil",
			writer:  nil,
			newFile: nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fileop.WriteToFile(tt.writer, tt.newFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateNewFile(t *testing.T) {
	tests := []struct {
		name             string
		fileName         string
		newFileNum       int
		expectedErr      bool
		expectedFilePath string
	}{
		{
			name:             "basic test",
			fileName:         "testdata/testfile.txt",
			newFileNum:       1,
			expectedErr:      false,
			expectedFilePath: "testdata/testfile_part1.txt",
		},
		{
			name:             "invalid file path",
			fileName:         "",
			newFileNum:       1,
			expectedErr:      true,
			expectedFilePath: "",
		},
		{
			name:             "invalid file number",
			fileName:         "",
			newFileNum:       0,
			expectedErr:      true,
			expectedFilePath: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := fileop.CreateNewFile(test.fileName, test.newFileNum)
			if err == nil && test.expectedErr {
				t.Error("Expected error, got nil")
			}
			if err != nil && !test.expectedErr {
				t.Errorf("unexpected error: %v", err)
			}
			if test.expectedFilePath != "" {
				if _, err := os.Stat(test.expectedFilePath); os.IsNotExist(err) {
					t.Errorf("expected file %s does not exist", test.expectedFilePath)
				}
			}
		})
	}
}

func TestGetLines(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		content  string
		expected int
	}{
		{
			name:     "normal case",
			fileName: "testfile.txt",
			content:  "line1\nline2\nline3",
			expected: 3,
		},
		{
			name:     "empty file",
			fileName: "emptyFile.txt",
			content:  "",
			expected: 0,
		},
		{
			name:     "single line",
			fileName: "singleLine.txt",
			content:  "line1",
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filePath := filepath.Join(dirName, test.fileName)
			if err := os.WriteFile(filePath, []byte(test.content), 0644); err != nil {
				t.Fatalf("Failed to create test file. err is %v", err)
			}
			lineCount, err := fileop.GetLines(filePath)
			if err != nil {
				t.Errorf("Failed to get lines. err is %v", err)
			}
			if lineCount != test.expected {
				t.Errorf("Expected line count %d, got %d", test.expected, lineCount)
			}
		})
	}
}
