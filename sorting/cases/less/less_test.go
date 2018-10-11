package less

import "testing"

func TestLessCase(t *testing.T) {
	arr := []int{3, 1, 5, 9, 2}
	x := LessCase(arr, 5)
	if x != 3 {
		t.Fatalf("less case failed %v", x)
	}
	if arr[0] != 3 && arr[1] != 1 && arr[2] != 5 && arr[3] != 2 && arr[4] != 9 {
		t.Fatalf("less case failed arr %v", arr)
	}
}
