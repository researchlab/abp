package maxgap

import "sort"

func comparator(list []int) int {
	if len(list) < 2 {
		return 0
	}
	data := sort.IntSlice(list)
	sort.Sort(data)
	gap := data[1] - data[0]
	for i := 1; i < len(data); i++ {
		if (data[i] - data[i-1]) > gap {
			gap = data[i] - data[i-1]
		}
	}
	return gap
}
