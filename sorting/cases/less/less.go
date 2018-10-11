package less

//给定一个数组arr和值num, 要求将arr中大于num的值编排到arr的右边，小于等于的值排到arr左边输出分界线下标x
//要求 时间复杂度O(N), 空间复杂度O(1)
func LessCase(arr []int, num int) int {
	x := -1
	for i, v := range arr {
		if v <= num {
			arr[x+1], arr[i] = arr[i], arr[x+1]
			x++
		}
	}
	return x
}
