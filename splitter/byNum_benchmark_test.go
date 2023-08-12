package splitter_test

import (
	"testing"

	"github.com/wata1355/go-split-command/splitter"
)

// 通常処理 3.353s
func BenchmarkSplitByLines(b *testing.B) {
	filePath := "./testfile.txt"
	numFiles := 100

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		splitter.SplitByNumFiles(filePath, numFiles)
	}
}
