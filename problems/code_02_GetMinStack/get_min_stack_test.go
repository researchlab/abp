package code_02_GetMinStack

import (
	"fmt"
	"testing"
)

func TestGetMinStack(t *testing.T) {
	mystack := New()
	pushStack(mystack, []int{3, 5, 2, 6, 1})
	if err := comparator(mystack, []input{
		{1, 1},
		{6, 2},
		{2, 2},
		{5, 3},
		{3, 3},
	}); err != nil {
		t.Fatal(err)
	}
	if err := isEmpty(mystack); err != nil {
		t.Fatal(err)
	}
}

func pushStack(mystack *MyStack, data []int) {
	for _, v := range data {
		mystack.Push(v)
	}
}

type input struct {
	pop, min int
}

func comparator(mystack *MyStack, inputs []input) error {
	for _, i := range inputs {
		if v, ok := mystack.GetMin(); !ok || v != i.min {
			return fmt.Errorf("get min stack failed %v", v)
		}
		if v, ok := mystack.Pop(); !ok || v != i.pop {
			return fmt.Errorf("get min stack pop failed %v", v)
		}
	}
	return nil
}

func isEmpty(mystack *MyStack) error {
	if v, ok := mystack.GetMin(); ok {
		return fmt.Errorf("get empty min stack failed %v", v)
	}
	if v, ok := mystack.Pop(); ok {
		return fmt.Errorf("get empty min stack pop failed %v", v)
	}
	return nil
}
