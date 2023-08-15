package splitter

import (
	"bufio"
	"os"

	"github.com/wata1355/go-split-command/fileop"
)

func SplitByLines(filePath string, linesPerFile int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var writer *bufio.Writer
	var newFile *os.File
	lineCount := 0
	newFileNum := 0

	for scanner.Scan() {
		if lineCount == 0 || lineCount%linesPerFile == 0 {
			if err := fileop.WriteToFile(writer, newFile); err != nil {
				return err
			}
			newFileNum++

			newFile, err = fileop.CreateNewFile(filePath, newFileNum)
			if err != nil {
				return err
			}

			writer = bufio.NewWriter(newFile)
		}

		writer.WriteString(scanner.Text() + "\n")
		lineCount++
	}

	if err := fileop.WriteToFile(writer, newFile); err != nil {
		return err
	}

	return scanner.Err()
}
