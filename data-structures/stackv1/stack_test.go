package stackv1

import "testing"

func TestStack(t *testing.T) {
	s := New()

	s.Push(1)
	s.Push(2)

	if s.stack[0] != 2 && s.stack[1] != 1 {
		t.Fatalf("stack push failed")
	}

	if v, ok := s.Pop(); !ok || v != 2 {
		t.Fatalf("stack pop failed")
	}
}
