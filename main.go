// TODO
// 1. 動くものを作る　100%
// 2. 可読性を上げる 70%
// 3. パフォーマンス性を上げる 20%
// 4. イレギュラーな入力に対応する 80%
// 5. 単体テストの導入 0%
// 6. 追加のオプションの実装 0%

// 工夫した点

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

	if err := validators.ValidateArgs(fileArgs, options); err != nil {
		log.Fatal(err)
	}

	filePath := fileArgs[0]
	err := splitters.Split(options, filePath)
	if err != nil {
		log.Fatalf("Failed to split file. error message is %s", err)
	}
}
