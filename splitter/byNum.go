package splitter

import (
	"github.com/wata1355/go-split-command/utils"
)

func SplitByNumFiles(filePath string, numFiles int) error {
	lineNum, err := utils.GetLines(filePath)
	if err != nil {
		return err
	}
	linesPerFile := lineNum / numFiles
	return SplitByLines(filePath, linesPerFile)
}
