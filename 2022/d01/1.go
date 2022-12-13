package d01

import (
	"aoc2022/common"
	"fmt"
	"sort"
	"strconv"
)

func P1() {
	max := 0
	a := 0

	common.FileIter("1", func(s string) {
		if s == "" {
			if a > max {
				max = a
			}
			a = 0
		}
		val, _ := strconv.Atoi(s)
		a = a + val
	})

	fmt.Println(max)
}

func P2() {
	max := make([]int, 0)
	a := 0

	common.FileIter("1", func(s string) {
		if s == "" {
			max = append(max, a)
			a = 0
		}
		val, _ := strconv.Atoi(s)
		a = a + val
	})
	sort.Ints(max)
	max = max[len(max)-3:]
	fmt.Println(max)
	fmt.Println(common.Sum(max))
}
