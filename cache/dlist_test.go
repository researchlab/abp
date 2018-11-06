package cache

import (
	"container/list"
	"testing"
)

func TestDList(t *testing.T) {
	dlist := list.New()
	dlist.PushBack(1)
	dlist.PushBack(2)
	e3 := dlist.PushFront(3)
	for e := dlist.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
	t.Log("after move e3 to back")
	dlist.MoveToBack(e3)
	for e := dlist.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}

}
