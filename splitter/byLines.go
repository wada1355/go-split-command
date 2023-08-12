package splitter

import (
	"bufio"
	"errors"
	"os"
	"sync"

	"github.com/wata1355/go-split-command/utils"
)

type splitTask struct {
	startLine int
	endLine   int
	fileNum   int
}

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
			if err := utils.WriteToFile(writer, newFile); err != nil {
				return err
			}
			newFileNum++
			newFile, err = utils.CreateNewFile(filePath, newFileNum)
			if err != nil {
				return err
			}
			writer = bufio.NewWriter(newFile)
		}

		writer.WriteString(scanner.Text() + "\n")
		lineCount++
	}

	if err := utils.WriteToFile(writer, newFile); err != nil {
		return err
	}

	return scanner.Err()
}

func SplitByLines_parallel(filePath string, linesPerFile int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	lines, err := readLines(filePath)
	if err != nil {
		return err
	}
	if len(lines) == 0 {
		return errors.New("Specified file was empty")
	}

	var tasks []splitTask
	totalLines := len(lines)
	for startLine := 0; startLine < totalLines; startLine += linesPerFile {
		endLine := startLine + linesPerFile
		if endLine > totalLines {
			endLine = totalLines
		}
		tasks = append(tasks, splitTask{startLine, endLine, startLine / linesPerFile})
	}

	var wg sync.WaitGroup
	for _, t := range tasks {
		wg.Add(1)
		go func(task splitTask) {
			defer wg.Done()
			processSplitTask(filePath, task, lines[task.startLine:task.endLine])
		}(t)
	}

	wg.Wait()

	return nil
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func processSplitTask(filePath string, task splitTask, lines []string) error {
	newFile, err := utils.CreateNewFile(filePath, task.fileNum)
	if err != nil {
		return err
	}
	defer newFile.Close()

	writer := bufio.NewWriter(newFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
