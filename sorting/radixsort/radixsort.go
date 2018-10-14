package radixsort

import (
	"math"
)

//桶排序的另一种实现， 基数排序
//其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较。由于整数也可以表达字符串（比如名字或日期）和特定格式的浮点数，所以基数排序也不是只能使用于整数
//实现思路:将所有待比较数值（正整数）统一为同样的数字长度，数字较短的数前面补零。然后，从最低位开始，依次进行一次排序。这样从最低位排序一直到最高位排序完成以后，数列就变成一个有序序列。
//基数排序适应于 整数排序(非负数), 字符串/浮点数等自定义排序，多关键字排序
func RadixSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	radixSort(arr, 0, len(arr)-1, maxbits(arr))
}

func maxbits(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	res := 0
	for max != 0 {
		res++
		max /= 10
	}
	return res
}

func radixSort(arr []int, begin, end, digit int) {
	radix := 10
	i, j := 0, 0
	count := make([]int, 10)
	bucket := make([]int, end-begin+1)
	for d := 1; d <= digit; d++ {
		for i := 0; i < radix; i++ { //初始化10个桶
			count[i] = 0
		}
		for i = begin; i <= end; i++ {
			j = getDigit(arr[i], d) //计算对应桶的id
			count[j]++              //桶中应该存储了多少个数字
		}
		for i = 1; i < radix; i++ { //计算到i位置 当前以及存在的元素个数
			count[i] = count[i] + count[i-1]
		}
		for i = end; i >= begin; i-- { //依次进行排序
			j = getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}
		for i, j = begin, 0; i <= end; { //复制到arr对应位置上
			arr[i] = bucket[j]
			i++
			j++
		}
	}
}

func getDigit(x, d int) int {
	//radix :=10
	//(x / int(math.Pow10(d-10))) % radix
	return (x / int(math.Pow10(d-1))) % 10
}
