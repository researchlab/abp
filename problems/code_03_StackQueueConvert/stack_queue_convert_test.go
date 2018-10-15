package code_03_StackQueueConvert

import "testing"

func TestTwoStacksQueue(t *testing.T) {
	tsq := NewTSQ()
	queue := []int{1, 2, 3, 4, 5}
	dequeue := []int{1, 2, 3, 4, 5}
	for _, v := range queue {
		tsq.Push(v)
	}

	for _, v := range dequeue {
		if v != tsq.Peek() || v != tsq.Poll() {
			t.Fatalf("two stack queue failed %v", v)
		}
	}

	enqueue := []int{1, 2, 3}
	enqueue1 := []int{5, 6}
	for _, v := range enqueue {
		tsq.Push(v)
	}
	if tsq.Poll() != enqueue[0] {
		t.Fatalf("two stack queue push faield")
	}
	for _, v := range enqueue1 {
		tsq.Push(v)
	}
	dequeue1 := []int{2, 3, 5, 6}
	for _, v := range dequeue1 {
		if v != tsq.Peek() || v != tsq.Poll() {
			t.Fatalf("two stack queue pop failed")
		}
	}
}

func TestTwoQueuesStack(t *testing.T) {
	tqs := NewTQS()
	pushStack := []int{1, 2, 3, 4, 5}
	popStack := []int{5, 4, 3, 2, 1}
	for _, v := range pushStack {
		tqs.Push(v)
	}

	for _, v := range popStack {
		if v != tqs.Peek() || v != tqs.Pop() {
			t.Fatalf("two queues stack failed")
		}
	}

	pushStack1 := []int{1, 2, 3}
	for _, v := range pushStack1 {
		tqs.Push(v)
	}
	if tqs.Peek() != pushStack1[2] || tqs.Pop() != pushStack1[2] {
		t.Fatalf("two queues stack pop failed")
	}
	for _, v := range pushStack1 {
		tqs.Push(v)
	}
	popStack1 := []int{3, 2, 1, 2, 1}
	for _, v := range popStack1 {
		if v != tqs.Peek() || v != tqs.Pop() {
			t.Fatalf("two queues stack failed")
		}
	}
}
