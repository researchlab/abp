package less

//给定一个数组arr和值num，要把小于num的数排在数组左边，等于num的数排在数组中间，大于num的数排序在数组的右边
//时间复杂度要为O(N)，空间复杂度为O(1)

func LessV1(arr []int, num int) (int, int) {
	less, cur, more := -1, 0, len(arr)
	for cur < more {
		if arr[cur] == num {
			cur++
		} else if arr[cur] < num {
			arr[less+1], arr[cur] = arr[cur], arr[less+1]
			less++
			cur++
		} else {
			arr[more-1], arr[cur] = arr[cur], arr[more-1]
			more--
		}
	}
	return less, more
}
