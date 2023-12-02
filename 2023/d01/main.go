package main

import (
	"aoc2022/common"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(P1())
	fmt.Println(P2())
}

func P1() int {
	sum := 0
	common.FileIt("input.txt", func(s string) {
		f, l := -1, 0
		for _, c := range s {
			if c >= '0' && c <= '9' {
				l = int(c - '0')
				if f < 0 {
					f = l
				}
			}
		}
		sum += f*10 + l
	})
	return sum
}

func P2() int {
	type Num struct {
		s string
		i int
	}
	nums := []Num{
		Num{"one", 1},
		Num{"two", 2},
		Num{"three", 3},
		Num{"four", 4},
		Num{"five", 5},
		Num{"six", 6},
		Num{"seven", 7},
		Num{"eight", 8},
		Num{"nine", 9},
	}
	sum := 0
	common.FileIt("input.txt", func(s string) {
		f, l := -1, 0
		for i, c := range s {
			if c >= '0' && c <= '9' {
				l = int(c - '0')
				if f < 0 {
					f = l
				}
			} else {
				for _, num := range nums {
					if strings.Index(s[i:], num.s) == 0 {
						l = num.i
						if f < 0 {
							f = l
						}
						break
					}
				}
			}
		}
		sum += f*10 + l
	})
	return sum
}
