package reverseorder

import (
	"log"
)

func ReversOrder(data []int) {
	if len(data) < 2 {
		return
	}
	MergeSort(data, 0, len(data)-1)
}

func MergeSort(data []int, L, R int) {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)
	MergeSort(data, L, mid)
	MergeSort(data, mid+1, R)
	Merge(data, L, mid, R)
}

func Merge(data []int, L, mid, R int) {
	help := make([]int, 0, (R - L + 1))
	left, right := L, mid+1
	for left <= mid && right <= R {
		if data[left] < data[right] {
			help = append(help, data[left])
			left++
		} else {
			//当遇到第一个左边的数比右边大时，则从左边的这个数开始到左边数组结束中所有数字就都会比右边这个数字大了
			for i := left; i <= mid-left+1; i++ {
				log.Printf("reverse-order %v %v\n", data[i], data[right])
			}
			help = append(help, data[right])
			right++
		}
	}
	for left <= mid {
		help = append(help, data[left])
		left++
	}
	for right <= R {
		help = append(help, data[right])
		right++
	}
	for i := 0; i < len(help); i++ {
		data[L+i] = help[i]
	}
}
