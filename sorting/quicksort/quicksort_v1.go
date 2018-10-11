package quicksort

func QuickSortV1(data []int) {
	if len(data) < 2 {
		return
	}
	quickSortV1(data, 0, len(data)-1)
}
func quickSortV1(data []int, L, R int) {
	s := NewStack()
	s.Push(&Point{Less: L, More: R})
	for !s.IsEmpty() {
		po := s.Pop()
		if po.Less < po.More {
			less, more := partitionV1(data, po.Less, po.More)
			if less > po.Less { //返回的左边界少于上一层的左边界, 将返回的左边界至上一层的左边界中间的'右边区域'压栈'
				s.Push(&Point{Less: po.Less, More: less - 1})
			}
			if more < po.More {
				s.Push(&Point{Less: more + 1, More: po.More})
			}
		}
	}
}

func partitionV1(data []int, L, R int) (int, int) {
	less := L - 1
	more := R //先划定data[len(R)-1]位置值为'大于'data[len(R)-1]的值,即len(data)-1位置不参与排序，最后再与more位置替换一下
	for L < more {
		if data[L] > data[R] {
			data[L], data[more-1] = data[more-1], data[L]
			more--
		} else if data[L] < data[R] {
			data[L], data[less+1] = data[less+1], data[L]
			less++
			L++
		} else {
			L++
		}
	}
	data[more], data[R] = data[R], data[more]
	return less + 1, more
}
