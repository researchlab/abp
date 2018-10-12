package maxgap

//给定一个数组，求如果排序之后，相邻两数的最大差值，要求时 间复杂度O(N)，且要求不能用非基于比较的排序。
func MaxGap(data []int) int {
	if len(data) < 2 {
		return 0
	}
	size := len(data)
	min, max := data[0], data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	if min == max {
		return 0
	}
	hasNum := make([]bool, size+1)
	mins := make([]int, size+1)
	maxs := make([]int, size+1)
	bid := 0

	for i := 0; i < size; i++ {
		bid = bucket(data[i], size, min, max)
		if hasNum[bid] {
			if data[i] < mins[bid] {
				mins[bid] = data[i]
			}
			if data[i] > maxs[bid] {
				maxs[bid] = data[i]
			}
		} else {
			mins[bid], maxs[bid] = data[i], data[i]
			hasNum[bid] = true
		}
	}
	res := 0
	lastMax := maxs[0]
	for i := 1; i <= size; i++ {
		if hasNum[i] {
			if (mins[i] - lastMax) > res {
				res = mins[i] - lastMax
			}
			lastMax = maxs[i]
		}
	}
	return res
}

//将最大最小值区间分成size分求bucket id
func bucket(num, size, min, max int) int {
	return (num - min) * size / (max - min)
}
