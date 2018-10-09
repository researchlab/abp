package insert

func InsertSort(data []int) {
	if data == nil || len(data) < 2 {
		return
	}
	for i := 1; i < len(data); i++ {
		for j := i - 1; j >= 0 && data[j] > data[j+1]; j-- {
			data[j], data[j+1] = data[j+1], data[j]
		}
	}
}
