package bucketsort

import (
	"sort"
	"testing"

	"github.com/researchlab/abp/sorting/utils"
)

func TestBucketSort(t *testing.T) {
	arr := []int{1, 3, 1, 7, 6, 2}
	expected := []int{1, 1, 2, 3, 6, 7}
	BucketSort(arr)
	for i, v := range arr {
		if expected[i] != v {
			t.Fatalf("bucket sort failed %v %v", i, arr)
		}
	}
}

func TestBucketSortByLogarithm(t *testing.T) {
	maxSize, maxValue, maxTime := 100, 100, 5000
	for i := 0; i < maxTime; i++ {
		arr := utils.GenerateRandomIntArray(maxSize, maxValue)
		arr1 := sort.IntSlice(utils.CopyArray(arr))
		BucketSort(arr)
		arr1.Sort()
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr[i] {
				t.Fatalf("bucket sort failed")
			}
		}
	}
}
