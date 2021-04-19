package radixsort

import (
	"sort"
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestRadixSort(t *testing.T) {
	arr := []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 233, 144}
	RadixSort(arr)
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			t.Fatalf("radix sort failed")
		}
	}
	t.Log(arr)
}

func TestRadixSort2(t *testing.T) {
	arr := utils.GetArrayOfSize(100)
	RadixSort(arr)
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			t.Fatalf("radix sort failed")
		}
	}
}

func TestRadixSortByLogarithm(t *testing.T) {
	maxSize, maxValue, maxTime := 100, 100, 5000
	for i := 0; i < maxTime; i++ {
		arr := utils.GenerateRandomIntArray(maxSize, maxValue)
		arr1 := sort.IntSlice(utils.CopyArray(arr))
		arr1.Sort()
		RadixSort(arr)
		for i := 0; i < len(arr); i++ {
			if arr[i] != arr1[i] {
				t.Fatalf("radix sort by logarithm failed")
			}
		}

	}
}

func benchmarkRadixSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i <= b.N; i++ {
		RadixSort(list)
	}
}

func BenchmarkRadixSort100(b *testing.B)   { benchmarkRadixSort(100, b) }
func BenchmarkRadixSort1000(b *testing.B)  { benchmarkRadixSort(1000, b) }
func BenchmarkRadixSort10000(b *testing.B) { benchmarkRadixSort(10000, b) }
