package splitter

import (
	"errors"
)

type Options struct {
	Lines    int
	NumFiles int
	Bytes    int
	FilePath string
}

func Split(options Options) error {
	switch {
	case options.Lines > 0:
		return SplitByLines(options.FilePath, options.Lines)
	case options.NumFiles > 0:
		return SplitByNumFiles(options.FilePath, options.NumFiles)
	case options.Bytes > 0:
		return SplitByBytes(options.FilePath, options.Bytes)
	default:
		return errors.New("please specify a split option (-l, -n, or -b)")
	}
}
