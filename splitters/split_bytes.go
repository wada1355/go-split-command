package splitters

import (
	"io"
	"os"

	"github.com/wata1355/go-split-command/fileop"
)

func (splitter *Splitter) SplitByBytes() error {
	file, err := os.Open(splitter.FileArgs.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, splitter.Options.Bytes)
	newFileNum := 0

	for {
		bytesRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		newFileNum++
		newFile, err := fileop.CreateNewFile(splitter.FileArgs.FilePath, newFileNum)
		if err != nil {
			return err
		}
		defer newFile.Close()

		if _, err := newFile.Write(buffer[:bytesRead]); err != nil {
			return err
		}
	}
	return nil
}
