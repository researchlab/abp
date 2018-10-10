package merge

import (
	"testing"

	"github.com/researchlab/algorithms-cs/sorting/utils"
)

func TestMergeV1Sort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	list = MergeV1Sort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("merge v1 sort failed index %v", i)
		}
	}
}

func benchmarkMergeV1Sort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		MergeV1Sort(list)
	}
}

func BenchmarkMergeV1Sort100(b *testing.B)   { benchmarkMergeV1Sort(100, b) }
func BenchmarkMergeV1Sort1000(b *testing.B)  { benchmarkMergeV1Sort(1000, b) }
func BenchmarkMergeV1Sort10000(b *testing.B) { benchmarkMergeV1Sort(10000, b) }
