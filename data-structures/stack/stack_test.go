package stack

import "testing"

func TestStack(t *testing.T) {
	arr := []int{1, 3, 5, 2, 7}
	s := New()
	for _, v := range arr {
		s.Push(v)
	}
	got := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		data, ok := s.Pop()
		if !ok {
			break
		}
		got = append(got, data.(int))
	}

	for i, v := range arr {
		if got[len(arr)-1-i] != v {
			t.Fatalf("stack pop failed %v", got)
		}
	}
}
