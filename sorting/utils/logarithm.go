package utils

import (
	"fmt"
	"math/rand"
	"sort"
)

func Logarithm(maxSize, maxValue, maxTime int, funcSort func(arr []int)) error {
	for i := 0; i < maxTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		funcSort(arr1)
		sort.Ints(arr2)
		if err := isEqual(arr1, arr2); err != nil {
			return err
		}
	}
	return nil
}

func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr = append(arr, ((maxValue+1)*rand.Intn(maxValue) - maxValue*rand.Intn(maxValue)))
	}
	return arr
}

func copyArray(arr []int) []int {
	_arr := make([]int, len(arr))
	for k, v := range arr {
		_arr[k] = v
	}
	return _arr
}

func isEqual(arr1, arr2 []int) error {
	if arr1 == nil && arr2 == nil {
		return nil
	}
	if (arr1 == nil || arr2 == nil) ||
		(len(arr1) != len(arr2)) {
		return fmt.Errorf("arr1 is not equal to arr2")
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return fmt.Errorf("arr1(%v) is not equal to arr2(%v) index %v", arr1[i], arr2[i], i)
		}
	}
	return nil
}
