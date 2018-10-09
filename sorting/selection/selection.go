package selection

func SelectionSort(data []int) {
	if len(data) < 2 || data == nil {
		return
	}
	//遍历所有元素
	for i := 0; i < len(data)-1; i++ {
		minIndex := i
		//查找从i+1 到最后一个元素的最小值
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
