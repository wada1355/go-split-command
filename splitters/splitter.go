package splitters

import (
	"errors"
)

type Options struct {
	Lines    int
	NumFiles int
	Bytes    int
}

func Split(options Options, filePath string) error {
	switch {
	case options.Lines > 0:
		return SplitByLines(filePath, options.Lines)
	case options.NumFiles > 0:
		return SplitByNumFiles(filePath, options.NumFiles)
	case options.Bytes > 0:
		return SplitByBytes(filePath, options.Bytes)
	default:
		return errors.New("please specify a split option (-l, -n, or -b)")
	}
}
