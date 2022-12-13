package d09

import (
	"aoc2022/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// P1 possible that works only for my input, better is to use solution like in P2
func P1() {
	pos := make(map[string]struct{})
	hi, hj, ti, tj := 0, 0, 0, 0
	addPos(pos, ti, tj)

	common.FileIter("9", func(s string) {
		instr := strings.Split(s, " ")
		ch, _ := strconv.Atoi(instr[1])
		for i := 0; i < ch; i++ {
			switch instr[0] {
			case "U":
				hi += 1
			case "D":
				hi -= 1
			case "R":
				hj += 1
			case "L":
				hj -= 1
			}

			if common.Abs(hi-ti) > 1 {
				tj = hj
				if hi > ti {
					ti = hi - 1
				} else {
					ti = hi + 1
				}
			}

			if common.Abs(hj-tj) > 1 {
				ti = hi
				if hj > tj {
					tj = hj - 1
				} else {
					tj = hj + 1
				}
			}

			addPos(pos, ti, tj)
		}
	})

	fmt.Println(len(pos))
}

func P2() {
	pos := make(map[string]struct{})

	const ropeSize = 10
	tail := ropeSize - 1
	var rope []common.Position
	for i := 0; i < ropeSize; i++ {
		rope = append(rope, common.Position{})
	}

	addPos(pos, rope[tail].R, rope[tail].C)

	common.FileIter("9", func(s string) {
		instr := strings.Split(s, " ")
		ch, _ := strconv.Atoi(instr[1])
		for i := 0; i < ch; i++ {
			//move head
			switch instr[0] {
			case "U":
				rope[0] = rope[0].Up()
			case "D":
				rope[0] = rope[0].Down()
			case "R":
				rope[0] = rope[0].Right()
			case "L":
				rope[0] = rope[0].Left()
			}
			// move everything but head
			for a := 0; a < tail; a++ {
				if !rope[a+1].Touching(rope[a]) {
					movePos(&rope[a], &rope[a+1])
					//drawMap2(rope)
				}
			}
			//drawMap2(rope)
			addPos(pos, rope[tail].R, rope[tail].C)
		}
		//drawMap2(rope)
	})

	fmt.Println(len(pos))
}

func movePos(head *common.Position, tail *common.Position) {
	diffR := head.R - tail.R
	diffC := head.C - tail.C
	tail.R += int(math.Round(float64(diffR) / 2))
	tail.C += int(math.Round(float64(diffC) / 2))
}

func drawMap2(rope []common.Position) {
	const head = 0
	x, mx := 0, 0
	y, my := 0, 0

	for i := 0; i < len(rope); i++ {
		if rope[i].R > x {
			x = rope[i].R
		}
		if rope[i].R < mx {
			mx = rope[i].R
		}
		if rope[i].C > y {
			y = rope[i].C
		}
		if rope[i].C < my {
			my = rope[i].C
		}
	}

	fmt.Println("====")
	for i := x; i >= mx; i-- {
		for j := my; j <= y; j++ {
			inRope := 0
			for x := 1; x < len(rope); x++ {
				if rope[x].R == i && rope[x].C == j {
					inRope = x
					break
				}
			}

			if i == rope[head].R && j == rope[head].C {
				fmt.Print("H")
			} else if inRope > 0 {
				fmt.Print(inRope)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("====")
}

func addPos(pos map[string]struct{}, ti int, tj int) {
	pos[strconv.Itoa(ti)+strconv.Itoa(tj)] = struct{}{}
}
