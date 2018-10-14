package queue

import "testing"

func TestQueue(t *testing.T) {
	arr := []int{1, 2, 5, 1, 3}
	nqueue := New()
	for _, v := range arr {
		nqueue.Push(v)
	}
	got := make([]int, 0, 5)
	for i := 0; i < len(arr); i++ {
		if data, ok := nqueue.Pop(); ok {
			if !ok {
				break
			}
			got = append(got, data.(int))
		}
	}
	for i, v := range got {
		if arr[i] != v {
			t.Fatalf("queue pop failed")
		}
	}
}
