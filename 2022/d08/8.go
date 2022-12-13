package d08

import (
	"aoc2022/common"
	"fmt"
)

func P1() {
	mapa, rows, cols := loadMap()
	c := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if vidnoHo(mapa, i, j) {
				c++
			}
		}
	}
	fmt.Println(c)
}

func P2() {
	mapa, rows, cols := loadMap()
	max := 0
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			score := score(mapa, i, j)
			if score > max {
				max = score
			}
		}
	}
	fmt.Println(max)
}

func vidnoHo(mapa map[int]map[int]int, i int, j int) bool {
	if i == 0 || i == len(mapa)-1 || j == 0 || j == len(mapa[i])-1 {
		return true
	}

	zl := l(mapa, i, j)
	zp := p(mapa, i, j)
	zh := h(mapa, i, j)
	zd := d(mapa, i, j)

	return zl || zp || zh || zd
}

func l(mapa map[int]map[int]int, i int, j int) bool {
	for j2 := 0; j2 < j; j2++ {
		if mapa[i][j] <= mapa[i][j2] {
			return false
		}
	}
	return true
}

func p(mapa map[int]map[int]int, i int, j int) bool {
	for j2 := j + 1; j2 < len(mapa[i]); j2++ {
		if mapa[i][j] <= mapa[i][j2] {
			return false
		}
	}
	return true
}

func h(mapa map[int]map[int]int, i int, j int) bool {
	for i2 := 0; i2 < i; i2++ {
		if mapa[i][j] <= mapa[i2][j] {
			return false
		}
	}
	return true
}

func d(mapa map[int]map[int]int, i int, j int) bool {
	for i2 := i + 1; i2 < len(mapa[i]); i2++ {
		if mapa[i][j] <= mapa[i2][j] {
			return false
		}
	}
	return true
}

func loadMap() (map[int]map[int]int, int, int) {
	mapa := make(map[int]map[int]int)
	rows := 0
	cols := 0
	common.FileIter("8", func(s string) {
		cols = len(s)
		mapa[rows] = make(map[int]int)
		for i, c := range s {
			mapa[rows][i] = int(c - '0')
		}
		rows++
	})
	return mapa, rows, cols
}

func score(mapa map[int]map[int]int, i int, j int) int {
	zl := sl(mapa, i, j)
	zp := sp(mapa, i, j)
	zh := sh(mapa, i, j)
	zd := sd(mapa, i, j)

	return zl * zp * zh * zd
}

func sl(mapa map[int]map[int]int, i int, j int) int {
	s := 0
	for j2 := j - 1; j2 >= 0; j2-- {
		s++
		if mapa[i][j] <= mapa[i][j2] {
			break
		}
	}
	return s
}

func sp(mapa map[int]map[int]int, i int, j int) int {
	s := 0
	for j2 := j + 1; j2 < len(mapa[i]); j2++ {
		s++
		if mapa[i][j] <= mapa[i][j2] {
			break
		}
	}
	return s
}

func sh(mapa map[int]map[int]int, i int, j int) int {
	s := 0
	for i2 := i - 1; i2 >= 0; i2-- {
		s++
		if mapa[i][j] <= mapa[i2][j] {
			break
		}
	}
	return s
}

func sd(mapa map[int]map[int]int, i int, j int) int {
	s := 0
	for i2 := i + 1; i2 < len(mapa[i]); i2++ {
		s++
		if mapa[i][j] <= mapa[i2][j] {
			break
		}
	}
	return s
}
