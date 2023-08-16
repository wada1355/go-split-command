package splitters_test

import (
	"testing"

	"github.com/wata1355/go-split-command/splitters"
)

// 通常処理 3.353s
// 並行化（最初の実装）12.596s
// 並行化（2回目の実装）3.081s
func BenchmarkSplitByNumFiles(b *testing.B) {
	filePath := "./testfile.txt" // 1GB
	numFiles := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		splitters.SplitByNumFiles(filePath, numFiles)
	}
}
