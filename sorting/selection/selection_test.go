package selection

import (
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestSelectionSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	SelectionSort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("selection sort failed index %v", i)
		}
	}
}

func TestSelectionSortByLogarithm(t *testing.T) {
	if err := utils.Logarithm(100, 100, 5000, SelectionSort); err != nil {
		t.Fatal(err)
	}
}
