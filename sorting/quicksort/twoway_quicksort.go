package quicksort

import "fmt"

// 当待排序数组是一个数值重复率非常高的序列时，容易产生最坏O(n^2)的时间复杂度。
// 重复非常高的时候 容易产生左右子树不均衡
// 此时优化方案之一就是  双路快排, 将小于v的部分放到指针i（i为遍历用指针）的左边，将大于v的部分都放到指针j的右边，延用随机选base的思路。等于v时使用交换的方式处理。

/* 双路快排-赋值
原数据
5    6    3    4    7    8    9
取出第一个数据5 放到一边
然后从右往左找到一个比  5数据小的数  放到5的位置
4    6    3    4    7    8    9
然后从左往右找  找一个比5 大的数据   放到刚从右边取到的4的位置
4    6    3    6    7    8    9
从右往左找   找一个比5 小的数据  3   放到刚从左边取到的6的位置
4    3    3    6    7    8    9
此时两边往中间挤 数据全部找完， 然后把第一个去出的5 放到3的原位置
4    3    5    6    7    8    9
第一次排序完  后的数据

####二
5 左边的数据 按同样的原理进行排序
取出 4 放到一边
从右往左查找比4 小的数据
3    3
把4放到 原3的位置
3    4

这样原数据为
3    4    5    6    7    8    9

####三
然后取6右边的数据 同样的原理进行排序
比6小的数据没有
然后比较比6大的数据

没有

然后比较  7
然后比较  8
*/

/*
序列元素重复，是序列已经排好序的一种特殊情况，如果一个序列中的元素全部相同，也将出现最差情况。
如果序列中存在大量重复元素，在普通快排中，相等的元素会被全部放到分区点的一边，这样会大大增加快排的时间复杂度，双路快排就是用来解决这个问题的。
能够将序列均衡分开的分区点才是好的分区点。均匀分开意味着保持O(logn)的复杂度。
*/

func TwowayQuicksort(data []int, left int, right int) {
	if right-left <= 1 {
		return
	}
	p := twowayPartition(data, left, right)
	TwowayQuicksort(data, left, p)
	TwowayQuicksort(data, p+1, right)
}

func twowayPartition(data []int, left, right int) int {
	v := data[left]
	i, j := left+1, right-1
	for {
		for i < right && data[i] < v {
			i++
		}
		fmt.Println("j:", j)
		fmt.Println("i:", i)
		for j > left && data[j] > v {
			j--
		}
		fmt.Println("i,j", i, " -- ", j)
		if i >= j {
			break
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	data[left], data[j] = data[j], data[left]
	return j
}
