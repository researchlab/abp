package bucketsort

//桶排序适用于与数据状态有关，建议适用于知道数组数值最大值在0-200范围的数据进行排序
//桶排序过程
//1.找出待排序数组中最大值 (建议最大值返回在0-200之间才采用桶排序，实现思路为桶排序中的计数排序)
//2.申请一个最大值加+1的数组A
//3.再次遍历数组,将数值相应的数组A下标位置值累加1
//4.遍历数组A还原待排序数组
func BucketSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	bucket := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	i := 0
	for j := 0; j < max+1; j++ {
		for bucket[j] > 0 {
			arr[i] = j
			i++
			bucket[j]--
		}
	}
}
