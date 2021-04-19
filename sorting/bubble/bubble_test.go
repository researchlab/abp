package bubble

import (
	"math/rand"
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestBubbleSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	BubbleSort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("sort failed index %v", i)
		}
	}
}

func TestBubbleSort2(t *testing.T) {
	list := make([]int, 100)
	for i := 0; i < 100; i++ {
		list = append(list, rand.Intn(100))
	}
	BubbleSort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("sort failed index %v", i)
		}
	}

}

func TestBubbleSortByLogarithm(t *testing.T) {
	if err := utils.Logarithm(100, 100, 5000, BubbleSort); err != nil {
		t.Fatal(err)
	}
}

func benchmarkBubbleSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		BubbleSort(list)
	}
}

func BenchmarkBubbleSort100(b *testing.B)    { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkBubbleSort(100000, b) }
