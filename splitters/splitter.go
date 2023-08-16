package splitters

import (
	"errors"
)

type Options struct {
	Lines    int
	NumFiles int
	Bytes    int
}

type FileArgs struct {
	FilePath string
}

type Splitter struct {
	Options  Options
	FileArgs FileArgs
}

func (splitter *Splitter) Split() error {
	switch {
	case splitter.Options.Lines > 0:
		return splitter.SplitByLines()
	case splitter.Options.NumFiles > 0:
		return splitter.SplitByNumFiles()
	case splitter.Options.Bytes > 0:
		return splitter.SplitByBytes()
	default:
		return errors.New("please specify a split option (-l, -n, or -b)")
	}
}
