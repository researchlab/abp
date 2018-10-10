package cases

import "testing"

func TestSmallSum(t *testing.T) {
	array := []int{1, 3, 4, 2, 5}
	smallSum := SmallSum(array)
	if smallSum != 16 {
		t.Fatalf("small sum failed expected 26 got %v", smallSum)
	}
}
