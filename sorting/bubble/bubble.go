package bubble

func BubbleSort(data []int) {
	if data == nil || len(data) < 2 {
		return
	}
	for e := len(data) - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
			}
		}
	}
}
