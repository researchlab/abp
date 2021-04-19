package heap

import (
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestHeapSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	HeapSort(arr)
	for i, v := range arr {
		if expected[i] != v {
			t.Fatalf("heap sort failed %v %v", i, arr)
		}
	}
}

func TestHeapSort2(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	HeapSort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("heap sort failed %v", i)
		}
	}
}

func TestHeapSortByLogarithm(t *testing.T) {
	if err := utils.Logarithm(100, 100, 5000, HeapSort); err != nil {
		t.Fatal(err)
	}
}

func benchmarkHeapSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i <= b.N; i++ {
		HeapSort(list)
	}
}

func BenchmarkHeapSort100(b *testing.B)   { benchmarkHeapSort(100, b) }
func BenchmarkHeapSort1000(b *testing.B)  { benchmarkHeapSort(1000, b) }
func BenchmarkHeapSort10000(b *testing.B) { benchmarkHeapSort(10000, b) }
