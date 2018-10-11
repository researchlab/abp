package arraytoheap

import (
	"testing"

	"github.com/researchlab/algorithms-cs/sorting/utils"
)

func TestArrayToMaxHeap(t *testing.T) {
	arr := []int{2, 1, 3, 6, 0, 4}
	expected := []int{6, 3, 4, 1, 0, 2}
	ArrayToMaxHeap(arr)
	for i := 0; i < len(arr); i++ {
		if expected[i] != arr[i] {
			t.Fatalf("array to heap failed %v", arr)
		}
	}
}

func benchmarkArrayToMaxHeap(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		ArrayToMaxHeap(list)
	}
}

func BenchmarkArrayToMaxHeap100(b *testing.B)   { benchmarkArrayToMaxHeap(100, b) }
func BenchmarkArrayToMaxHeap1000(b *testing.B)  { benchmarkArrayToMaxHeap(1000, b) }
func BenchmarkArrayToMaxHeap10000(b *testing.B) { benchmarkArrayToMaxHeap(10000, b) }
