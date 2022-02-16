package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name   string
		args   []int
		wanted []int
	}{{
		name:   "建设最大堆",
		args:   []int{2, 1, 3, 6, 0, 4},
		wanted: []int{0, 1, 3, 6, 2, 4},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewHeap(len(tt.args))
			for _, data := range tt.args {
				heap.Push(data)
			}
			assert.Equal(t, tt.wanted, heap.Elems)
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name     string
		args     []int
		wanted   int
		wantedEx []int
	}{{
		name:     "Pop",
		args:     []int{2, 1, 3, 6, 0, 4},
		wanted:   0,
		wantedEx: []int{1, 2, 3, 6, 4, 4},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewHeap(len(tt.args))
			for _, v := range tt.args {
				heap.Push(v)
			}
			assert.Equal(t, tt.wanted, heap.Pop())
			assert.Equal(t, tt.wantedEx, heap.Elems)
		})
	}
}
