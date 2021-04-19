package maxgap

import (
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestMaxGap(t *testing.T) {
	arr := []int{1, 1, 3, 6, 7}
	val := MaxGap(arr)
	if val != 3 {
		t.Fatalf("max gap failed %v", val)
	}
}

func TestMaxGapByLogarithm(t *testing.T) {
	maxSize := 100
	maxValue := 100
	maxTime := 5000
	for i := 0; i < maxTime; i++ {
		arr := utils.GenerateRandomArray(maxSize, maxValue)
		arr1 := utils.CopyArray(arr)
		if MaxGap(arr) != comparator(arr1) {
			t.Fatalf("max gap failed")
		}
	}
}
