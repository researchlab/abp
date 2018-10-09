package utils

import (
	"bufio"
	"go/build"
	"os"
	"path/filepath"
	"strconv"
)

var datapath = "github.com/researchlab/algorithms-cs/sorting/utils"

func GetArrayOfSize(n int) []int {
	p, err := build.Default.Import(datapath, "", build.FindOnly)
	if err != nil {
		// handle error
		return nil
	}
	fname := filepath.Join(p.Dir, "IntegerArray.txt")
	f, _ := os.Open(fname)
	defer f.Close()

	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, s)
	}
	return numbers[0:n]
}
