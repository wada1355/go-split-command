package splitter

import (
	"bufio"
	"os"
)

func splitByNumFiles(filePath string, numFiles int) error {
	linesPerFile := getLines(filePath) / numFiles
	return splitByLines(filePath, linesPerFile)
}

func getLines(filePath string) int {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}
