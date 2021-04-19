package quicksort

import (
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestQuickSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	QuickSort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("quick sort failed %v", i)
		}
	}
}

func TestQuickSortByLogarithm(t *testing.T) {
	if err := utils.Logarithm(100, 100, 5000, QuickSort); err != nil {
		t.Fatal(err)
	}
}

func benchkmarkQuickSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		QuickSort(list)
	}
}

func BenchmarkQuickSort100(b *testing.B)   { benchkmarkQuickSort(100, b) }
func BenchmarkQuickSort1000(b *testing.B)  { benchkmarkQuickSort(1000, b) }
func BenchmarkQuickSort10000(b *testing.B) { benchkmarkQuickSort(10000, b) }
