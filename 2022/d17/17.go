package d17

import (
	"aoc2022/common"
	"fmt"
	"os"
	"strings"
)

const W = 7

type shape struct {
	min [W]int
	max [W]int
}

func (s *shape) left(h [W]int, pos int) {
	if s.max[0]-s.min[0] != 0 {
		return
	}
	for i := 0; i < W-1; i++ {
		s.min[i] = s.min[i+1]
		s.max[i] = s.max[i+1]
	}
	s.min[W-1] = 0
	s.max[W-1] = 0

	for i := 0; i < W; i++ {
		if s.contains(i) && pos+s.min[i] < h[i] {
			s.right(h, pos)
			return
		}
	}
}

func (s *shape) right(h [W]int, pos int) {
	if s.max[W-1]-s.min[W-1] != 0 {
		return
	}
	for i := W - 1; i > 0; i-- {
		s.min[i] = s.min[i-1]
		s.max[i] = s.max[i-1]
	}
	s.min[0] = 0
	s.max[0] = 0

	for i := 0; i < W; i++ {
		if s.contains(i) && pos+s.min[i] < h[i] {
			s.left(h, pos)
			return
		}
	}
}
func (s *shape) contains(pos int) bool {
	return s.max[pos] != 0
}

var SHAPES = [...]shape{
	{
		min: [W]int{},
		max: [W]int{0, 0, 1, 1, 1, 1, 0},
	},
	{
		min: [W]int{0, 0, 1, 0, 1},
		max: [W]int{0, 0, 2, 3, 2, 0, 0},
	},
	{
		min: [W]int{},
		max: [W]int{0, 0, 1, 1, 3, 0, 0},
	},
	{
		min: [W]int{},
		max: [W]int{0, 0, 4},
	},
	{
		min: [W]int{},
		max: [W]int{0, 0, 2, 2},
	},
}

func generator(s string) func() int {
	i := 0
	return func() (movement int) {
		if s[i] == '>' {
			movement = 1
		} else {
			movement = -1
		}
		i++
		if i == len(s) {
			i = 0
		}
		return
	}
}

func shapeGenerator() func() shape {
	i := 0
	return func() (s shape) {
		s = SHAPES[i]
		i++
		if i == len(SHAPES) {
			i = 0
		}
		return
	}
}

// P1 is not correct
func P1() {
	input, _ := os.ReadFile("inputs/17.txt")
	s := strings.TrimSpace(string(input))
	g := generator(s)
	sg := shapeGenerator()
	var heights [W]int

	max := 0
	for i := 0; i < 2022; i++ {
		shape := sg()

		shapePosition := max + 3
		for {
			if g() > 0 {
				shape.right(heights, shapePosition)
			} else {
				shape.left(heights, shapePosition)
			}
			if touches(heights, shapePosition, shape) {
				for i, _ := range shape.max {
					if !shape.contains(i) {
						continue
					}
					heights[i] = common.Max(shapePosition+shape.max[i], heights[i])
					if heights[i] > max {
						max = heights[i]
					}
				}
				break
			}
			shapePosition--
		}
		fmt.Println(max)
		for i := max; i > 0; i-- {
			for j := 0; j < W; j++ {
				if heights[j] < i {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Println()
		}
		fmt.Println("======")
	}

	fmt.Println(max)
}

func touches(heights [W]int, position int, s shape) bool {
	for i, _ := range s.min {
		if !s.contains(i) {
			continue
		}
		if heights[i] == position+s.min[i] {
			return true
		}
	}
	return false
}
