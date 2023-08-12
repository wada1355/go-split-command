package splitter

import (
	"bufio"
	"os"
)

func SplitByNumFiles(filePath string, numFiles int) error {
	lineNum, err := getLines(filePath)
	if err != nil {
		return err
	}
	linesPerFile := lineNum / numFiles
	return SplitByLines(filePath, linesPerFile)
}

func getLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	return lineCount, nil
}
