package d04

import (
	"aoc2022/common"
	"fmt"
	"strings"
)

func P1() {
	amt := 0
	common.FileIter("4", func(s string) {
		sections := strings.Split(s, ",")
		as := strings.Split(sections[0], "-")
		bs := strings.Split(sections[1], "-")
		a := common.SliceStrToInt(as)
		b := common.SliceStrToInt(bs)

		if (a[0] >= b[0] && a[1] <= b[1]) ||
			(b[0] >= a[0] && b[1] <= a[1]) {
			amt++
		}
	})
	fmt.Println(amt)
}

func P2() {
	amt := 0
	common.FileIter("4", func(s string) {
		sections := strings.Split(s, ",")
		as := strings.Split(sections[0], "-")
		bs := strings.Split(sections[1], "-")
		a := common.SliceStrToInt(as)
		b := common.SliceStrToInt(bs)

		if (a[0] >= b[0] && a[0] <= b[1]) ||
			(a[1] >= b[0] && a[1] <= b[1]) ||
			(a[0] >= b[0] && a[1] <= b[1]) ||
			(b[0] >= a[0] && b[1] <= a[1]) {
			amt++
		}
	})
	fmt.Println(amt)
}
