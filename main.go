package main

import (
	"flag"
	"log"

	"github.com/wata1355/go-split-command/splitters"
	"github.com/wata1355/go-split-command/validators"
)

func main() {
	var l, n, b int

	flag.IntVar(&l, "l", 0, "Number of lines per file")
	flag.IntVar(&n, "n", 0, "Number of files to split into")
	flag.IntVar(&b, "b", 0, "Number of bytes per file")

	flag.Parse()

	fileArgs := flag.Args()
	options := splitters.Options{
		Lines:    l,
		NumFiles: n,
		Bytes:    b,
	}

	if err := validators.ValidateArgs(fileArgs, &options); err != nil {
		log.Fatal(err)
	}

	filePath := fileArgs[0]

	splitter := &splitters.Splitter{
		Options:  options,
		FileArgs: splitters.FileArgs{FilePath: filePath},
	}
	err := splitter.Split()
	if err != nil {
		log.Fatalf("Failed to split file: %s", err)
	}
}
