package quicksort

import (
	"fmt"
	"testing"
)

func TestTwowayQuicksort(t *testing.T) {
	tests := []struct {
		args     []int
		expected []int
	}{
		{
			args:     []int{12, 1, 2},
			expected: []int{1, 2, 12},
		},
		{
			args:     []int{12, 1, 2, 13},
			expected: []int{1, 2, 12, 13},
		},
		{
			args:     []int{5, 6, 3, 4, 7, 8, 9},
			expected: []int{3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, test := range tests {
		TwowayQuicksort(test.args, 0, len(test.args))
		if fmt.Sprintf("%v", test.expected) != fmt.Sprintf("%v", test.args) {
			t.Errorf("twowayQuicksort() invalid, got:%v, expected:%v", test.args, test.expected)
		}
	}
	t.Logf("pass")
}
