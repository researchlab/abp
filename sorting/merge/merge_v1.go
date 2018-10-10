package merge

//归并排序v1思路是 递归并外排后不再替换之前的数组值，而是直接返回新数组
func MergeV1Sort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	m := len(data) / 2
	left := MergeV1Sort(data[0:m]) //左边数组递归排序
	right := MergeV1Sort(data[m:]) //右边数组递归排序
	return MergeV1(left, right)    //左右两边外排
}

func MergeV1(left, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			tmp = append(tmp, right[j])
			j++
		} else {
			tmp = append(tmp, left[i])
			i++
		}
	}
	tmp = append(tmp, right[j:]...)
	tmp = append(tmp, left[i:]...)
	return tmp
}
