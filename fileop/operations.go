package fileop

import (
	"bufio"
	"errors"
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
	if filePath == "" {
		return nil, errors.New("invalid file path")
	}
	if newFileNum < 1 {
		return nil, errors.New("invalid file number")
	}
	outputFileName := getOutputFileName(filePath, newFileNum)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %s: %v", outputFileName, err)
	}
	return outputFile, nil
}

func GetLines(filePath string) (int, error) {
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

func getOutputFileName(originalPath string, fileNum int) string {
	dir, filename := filepath.Split(originalPath)
	ext := filepath.Ext(filename)
	baseName := filename[:len(filename)-len(ext)]
	return filepath.Join(dir, fmt.Sprintf("%s_part%d%s", baseName, fileNum, ext))
}
