package merge

import (
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestMerge(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	MergeSort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("merge sort failed index %v", i)
		}
	}
}

func TestMergeByLogarithm(t *testing.T) {
	if err := utils.Logarithm(100, 100, 5000, MergeSort); err != nil {
		t.Fatalf("merge sort failed %v", err)
	}
}

func benchmarkMergeSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		MergeSort(list)
	}
}

func BenchmarkMergeSort100(b *testing.B)   { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)  { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B) { benchmarkMergeSort(10000, b) }
