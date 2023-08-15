package fileop_test

import (
	"log"
	"os"
	"testing"
)

var dirName = "testdata"

func TestMain(m *testing.M) {
	if err := os.MkdirAll(dirName, 0755); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}
	m.Run()
	os.RemoveAll(dirName)
}
