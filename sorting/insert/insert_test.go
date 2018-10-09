package insert

import (
	"testing"

	"github.com/researchlab/algorithms-cs/sorting/utils"
)

func TestInsertSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	InsertSort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("insert sort failed index %v", i)
		}
	}
}

func TestInsertSortByLogarithm(t *testing.T) {
	if err := utils.Logarithm(100, 100, 5000, InsertSort); err != nil {
		t.Fatal(err)
	}
}

func benchmarkInsertSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		InsertSort(list)
	}
}

func BenchmarkInsertSort100(b *testing.B)   { benchmarkInsertSort(100, b) }
func BenchmarkInsertSort1000(b *testing.B)  { benchmarkInsertSort(1000, b) }
func BenchmarkInsertSort10000(b *testing.B) { benchmarkInsertSort(10000, b) }
