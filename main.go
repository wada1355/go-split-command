package main

import (
	"flag"
	"log"

	"github.com/wata1355/go-split-command/splitter"
)

// TODO
// 1. 動くものを作る　100%
// 2. 可読性を上げる 70%
// 3. パフォーマンス性を上げる 0%
// 4. イレギュラーな入力に対応する 0%
// 5. 単体テストの導入 0%
// 6. 追加のオプションの実装 0%

func main() {
	l := flag.Int("l", 0, "Number of lines per file")
	n := flag.Int("n", 0, "Number of files to split into")
	b := flag.Int("b", 0, "Number of bytes per file")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("Please specify file as argument")
	}

	filePath := args[0]

	options := splitter.Options{
		Lines:    *l,
		NumFiles: *n,
		Bytes:    *b,
		FilePath: filePath,
	}

	err := splitter.Split(options)
	if err != nil {
		log.Fatalf("Failed to split file. error message is %s", err)
	}
}
