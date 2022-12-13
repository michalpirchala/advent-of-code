package d10

import (
	"aoc2022/common"
	"fmt"
	"strconv"
	"strings"
)

func P1() {
	c := 0
	X := 1
	score := 0
	var crt [240]uint8
	common.FileIter("10", func(s string) {
		cycle := func() {
			c++
			updateCrt(&crt, X, c)
			score += scoreCheck(c, X)

		}
		instr := strings.Split(s, " ")
		v := 0
		switch instr[0] {
		case "noop":
			break
		case "addx":
			v, _ = strconv.Atoi(instr[1])
			cycle()
		}
		cycle()
		X += v
	})
	fmt.Println(score)

	drawCrt(crt)
}

func drawCrt(crt [240]uint8) {
	for i := 0; i < 240; i++ {
		fmt.Printf("%c", crt[i])
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}

func updateCrt(crt *[240]uint8, X, c int) {
	p := c - 1
	if common.Abs((p%40)-X) <= 1 {
		(*crt)[p] = '#'
	} else {
		(*crt)[p] = '.'
	}
}

func scoreCheck(c, X int) int {
	if (c+20)%40 == 0 {
		return c * X
	}
	return 0
}
