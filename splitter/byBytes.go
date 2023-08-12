package splitter

import (
	"io"
	"os"

	"github.com/wata1355/go-split-command/utils"
)

func SplitByBytes(filePath string, bytesPerFile int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, bytesPerFile)
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
		newFile, err := utils.CreateNewFile(filePath, newFileNum)
		if err != nil {
			return err
		}
		defer newFile.Close()

		newFile.Write(buffer[:bytesRead])
	}
	return nil
}
