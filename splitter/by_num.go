package splitter

import (
	"errors"

	"github.com/wata1355/go-split-command/fileop"
)

func SplitByNumFiles(filePath string, numFiles int) error {
	lineNum, err := fileop.GetLines(filePath)
	if err != nil {
		return err
	}
	if lineNum == 0 {
		return errors.New("file has no lines")
	}
	var linesPerFile int
	if lineNum < numFiles {
		linesPerFile = lineNum
	} else {
		linesPerFile = lineNum / numFiles
	}
	return SplitByLines(filePath, linesPerFile)
}
