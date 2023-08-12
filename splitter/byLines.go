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

	// 一度ファイル全体をスキャンしているが、ゴルーチン内で再度ファイルを開いて部分的な読み取りを行っているため、結果的に二重の操作となっている
	totalLines, err := utils.GetLines(filePath)
	if err != nil {
		return err
	}
	if totalLines == 0 {
		return errors.New("Specified file was empty")
	}

	var tasks []splitTask
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
			processSplitTask(filePath, task)
		}(t)
	}

	wg.Wait()

	return nil
}

func processSplitTask(filePath string, task splitTask) error {
	// 各ゴルーチンごとにファイルをオープンしてしまうと以下のような問題が生じる
	// 例1：ファイルをオープンできない（OSには同時にオープンできるファイルの最大数という制限がある）
	// 例2：パフォーマンスの低下
	//     ・同時に多数のI/O操作が発生するため
	//     ・複数のゴルーチンが同時に同じファイルにアクセスすると、競合状態やディスクのシーク時間が増加するため
	//     ・メモリ使用量が増加するため
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var writer *bufio.Writer
	var newFile *os.File

	for i := 0; i <= task.endLine; i++ {
		if !scanner.Scan() {
			break
		}
		// 不要なループが発生している
		if i < task.startLine {
			continue
		}
		if i == task.startLine {
			newFile, err = utils.CreateNewFile(filePath, task.fileNum)
			if err != nil {
				return err
			}
			writer = bufio.NewWriter(newFile)
		}
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}
	}
	if err := utils.WriteToFile(writer, newFile); err != nil {
		return err
	}
	return scanner.Err()
}
