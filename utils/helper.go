package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func WriteToFile(writer *bufio.Writer, newFile *os.File) error {
	if writer == nil || newFile == nil {
		return nil
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	if err := newFile.Close(); err != nil {
		return err
	}
	return nil
}

func CreateNewFile(filePath string, newFileNum int) (*os.File, error) {
	outputFileName := getOutputFileName(filePath, newFileNum)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return nil, err
	}
	return outputFile, nil
}

func getOutputFileName(originalPath string, fileNum int) string {
	base := filepath.Base(originalPath)
	return fmt.Sprintf("%s.part_%d", base, fileNum)
}
