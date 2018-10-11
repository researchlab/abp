package quicksort

// 指定数组最后一个值x 作为划分值，递归遍历数组将小于等于x的都放到数组左边，大于的都放到左边;
func QuickSort(data []int) {
	if len(data) < 2 {
		return
	}
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, L, R int) {
	if L < R {
		less, more := partition(data, L, R)
		quickSort(data, L, less-1)
		quickSort(data, more+1, R)
	}
}

func partition(data []int, L, R int) (int, int) {
	less, more := L-1, R
	for L < more {
		if data[L] < data[R] {
			data[less+1], data[L] = data[L], data[less+1]
			less++
			L++
		} else if data[L] > data[R] {
			data[L], data[more-1] = data[more-1], data[L]
			more--
		} else {
			L++
		}
	}
	//因为是取data[R]为比较值，所以最后要把R与more位置值交互一下，此时less左边全是小于data[R],more右边全是大于data[R],less+1, more中间全是等于data[R]
	data[R], data[more] = data[more], data[R]
	return less + 1, more
}
