package splitters

import (
	"errors"

	"github.com/wata1355/go-split-command/fileop"
)

func (splitter *Splitter) SplitByNumFiles() error {
	lineNum, err := fileop.CountLines(splitter.FileArgs.FilePath)
	if err != nil {
		return err
	}
	if lineNum == 0 {
		return errors.New("file has no lines")
	}
	var linesPerFile int
	if lineNum < splitter.Options.NumFiles {
		linesPerFile = lineNum
	} else {
		linesPerFile = lineNum / splitter.Options.NumFiles
	}
	splitter.Options.Lines = linesPerFile
	return splitter.SplitByLines()
}
