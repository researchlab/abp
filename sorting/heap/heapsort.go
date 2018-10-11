package heap

func HeapSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	size := len(arr) //堆大小
	//交换大根堆根节点与尾节点的值
	size--
	arr[0], arr[size] = arr[size], arr[0]
	for size > 0 {
		heapify(arr, 0, size)
		size--
		arr[0], arr[size] = arr[size], arr[0]
	}
}

func heapInsert(arr []int, i int) {
	for arr[i] > arr[(i-1)/2] {
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
		i = (i - 1) / 2
	}
}

func heapify(arr []int, index, size int) {
	left := 2*index + 1
	for left < size {
		//找出左右节点最大值节点下标
		if (left+1 < size) && arr[left+1] > arr[left] {
			left = left + 1
		}
		//找出当前节点与其左右节点最大值下标
		if arr[index] > arr[left] {
			left = index
		}
		if left == index {
			break
		}
		arr[index], arr[left] = arr[left], arr[index]
		index = left
		left = 2*index + 1
	}
}

//func heapify(arr []int, index, heapSize int) {
//	left := 2*index + 1
//	for left < heapSize {
//		largest := left
//		if (left+1 < heapSize) && arr[left+1] > arr[left] {
//			largest = left + 1
//		}
//		if arr[index] > arr[largest] {
//			largest = index
//		}
//		if largest == index {
//			break
//		}
//		arr[index], arr[largest] = arr[largest], arr[index]
//		index = largest
//		left = 2*index + 1
//	}
//}
