package splitters_test

import (
	"log"
	"os"
	"testing"
)

var dirName = "testdata"

func checkFileContent(t *testing.T, filePath string, expectedContent string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %s", filePath)
	}
	if string(content) != expectedContent {
		t.Errorf("Unexpected content in %s. got %s but want %s", filePath, string(content), expectedContent)
	}
}

func TestMain(m *testing.M) {
	if err := os.MkdirAll(dirName, 0755); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}
	m.Run()
	os.RemoveAll(dirName)
}
