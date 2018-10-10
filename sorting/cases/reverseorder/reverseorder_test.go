package reverseorder

import "testing"

func TestReverseOrder(t *testing.T) {
	array := []int{1, 3, 4, 2, 5}
	ReversOrder(array)
	t.Log("expected {3,2} {4,2}")
}
