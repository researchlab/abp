package smallsum

//采用归并思想求解
//时间复杂度为归并算法复杂度 O(NlgN)
//算法是计算某个指定值其右边有多少个值比起大，则这个值存在的小何是 这个值*比它大的值的个数N
func SmallSum(data []int) int {
	if len(data) < 2 {
		return 0
	}
	return MergeSort(data, 0, len(data)-1)
}

func MergeSort(data []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := L + ((R - L) >> 1)
	leftSum := MergeSort(data, L, mid)
	rightSum := MergeSort(data, mid+1, R)
	mergeSum := merge(data, L, mid, R)
	return leftSum + rightSum + mergeSum
}

func merge(data []int, L, mid, R int) int {
	help := make([]int, 0, (R - L + 1))
	pleft, pright := L, mid+1
	res := 0
	for pleft <= mid && pright <= R {
		if data[pleft] < data[pright] {
			res += data[pleft] * (R - pright + 1)
			help = append(help, data[pleft])
			pleft++
		} else {
			help = append(help, data[pright])
			pright++
		}
	}
	for pleft <= mid {
		help = append(help, data[pleft])
		pleft++
	}
	for pright <= R {
		help = append(help, data[pright])
		pright++
	}
	for i := 0; i < len(help); i++ {
		data[L+i] = help[i]
	}
	return res
}
