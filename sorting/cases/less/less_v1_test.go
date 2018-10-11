package less

import "testing"

func TestLessV1(t *testing.T) {
	arr := []int{3, 1, 5, 2, 5, 9, 6, 7}
	num := 5
	l, m := LessV1(arr, num)
	if l != 3 && m != 5 {
		t.Fatalf("less v1 failed %v %v", l, m)
	}
	expected := []int{3, 1, 2, 5, 5, 6, 7, 9}
	for i, v := range arr {
		if expected[i] != v {
			t.Fatalf("less v1 failed arr %v expected %v %v", arr, expected, i)
		}
	}
}
