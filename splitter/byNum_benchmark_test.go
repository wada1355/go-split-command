package splitter_test

import (
	"testing"

	"github.com/wata1355/go-split-command/splitter"
)

// 通常処理 3.353s
// 並行化（1回目）12.596s
// 並行化（2回目）3.081s
func BenchmarkSplitByNumFiles(b *testing.B) {
	filePath := "./testfile.txt"
	numFiles := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		splitter.SplitByNumFiles(filePath, numFiles)
	}
}
