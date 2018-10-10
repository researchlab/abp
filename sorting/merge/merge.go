package merge

//下面merge 实现思路是
//将待排序数组分为左右两边，先对左边和右边进行递归排序
//将上述排序完成的左右两边在用外排进行排序并申请一个辅助数组进行接收，外排结束后将排序的数组复制覆盖到待排数组data的相应位置data[L+0, L+len(help)-1]
//时间复杂度分析
//将待排数组一切为左右两个待排数组进行递归排序，则子过程规模复杂度为 2T(N/2)
//外排将全部数据重新排序一次 时间复杂度为 O(N), 辅助空间为O(N)
//所以时间复杂度为 T(N) = 2T(N/2) + O(N) 应用Master定理得 时间复杂度为 O(NlgN) , 空间复杂度为O(N)
func MergeSort(data []int) {
	if len(data) < 2 {
		return
	}
	processSort(data, 0, len(data)-1)
}

func processSort(data []int, L, R int) {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)   // L+R中点的位置 (L+R)/2
	processSort(data, L, mid)   // T(N/2)
	processSort(data, mid+1, R) //T(N/2)
	merge(data, L, mid, R)      // O(N)
}

func merge(data []int, L, mid, R int) {
	help := make([]int, 0, (R - L + 1)) //特别注意如果 help := make([]int, (R-L+1)) 然后进行help = append(help,value) 那value 将从help[R-L+1]的起始位置往上加, 而help的0位置到(R-L+1)位置已经被初始化为0
	//当help := make([]int, 0, (R-L+1)) 则表示从help[0] 位置往上加value
	pleft, pright := L, mid+1
	for pleft <= mid && pright <= R {
		if data[pleft] > data[pright] {
			help = append(help, data[pright])
			pright++
		} else {
			help = append(help, data[pleft])
			pleft++
		}
	}

	for pright <= R {
		help = append(help, data[pright])
		pright++
	}

	for pleft <= mid {
		help = append(help, data[pleft])
		pleft++
	}

	for i, v := range help {
		data[L+i] = v
	}
}
