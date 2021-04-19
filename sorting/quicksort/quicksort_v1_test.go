package quicksort

import (
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestQuickSortV1(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	QuickSortV1(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("quick sort v1 failed %v", i)
		}
	}
}

func TestQuickSortV1ByLogarithm(t *testing.T) {
	if err := utils.Logarithm(100, 100, 5000, QuickSortV1); err != nil {
		t.Fatal(err)
	}
}

func benchmarkQuickSortV1(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i <= b.N; i++ {
		QuickSortV1(list)
	}
}

func BenchmarkQuickSortV1100(b *testing.B)   { benchmarkQuickSortV1(100, b) }
func BenchmarkQuickSortV11000(b *testing.B)  { benchmarkQuickSortV1(1000, b) }
func BenchmarkQuickSortV110000(b *testing.B) { benchmarkQuickSortV1(10000, b) }
