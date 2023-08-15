package validators_test

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := os.WriteFile("file.txt", []byte("Hello World"), 0644); err != nil {
		log.Fatalf("Failed to create file.txt. err is %v", err)
	}
	_, err := os.Create("emptyFile.txt")
	if err != nil {
		log.Fatalf("Failed to Create emptyFile")
	}
	m.Run()
	os.Remove("file.txt")
	os.Remove("emptyFile.txt")
}
