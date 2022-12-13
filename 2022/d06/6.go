package d06

import (
	"aoc2022/common"
	"fmt"
)

func P1() {
	six(4)
}

func P2() {
	six(14)
}

func six(n int) {
	n -= 1
	common.FileIter("6", func(s string) {
		mapa := make(map[uint8]int)

		i := 0
		for ; i < len(s); i++ {
			mapa[s[i]]++
			if i >= n {
				if allOnce(mapa) {
					break
				}
				mapa[s[i-n]]--
			}
		}
		fmt.Println(i + 1)
	})
}

func allOnce(mapa map[uint8]int) bool {
	for _, c := range mapa {
		if c > 1 {
			return false
		}
	}
	return true
}
