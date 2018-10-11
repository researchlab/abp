package arraytoheap

func ArrayToMaxHeap(data []int) {
	if len(data) < 2 {
		return
	}
	for i := 2; i < len(data); i++ {
		swap(data, i)
	}
}

func swap(data []int, i int) {
	f := (i - 1) / 2
	if data[i] > data[f] {
		data[i], data[f] = data[f], data[i]
		swap(data, f)
	}
}
