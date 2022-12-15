package common

import (
	"bufio"
	"os"
	"strconv"
)

func FileIter(num string, f func(string)) {
	file, _ := os.Open("inputs/" + num + ".txt")
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		f(scan.Text())
	}
	file.Close()
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func Sum(max []int) int {
	sum := 0
	for _, v := range max {
		sum = sum + v
	}
	return sum
}

type Position struct {
	R, C int
}

func (p *Position) Touching(p2 Position) bool {
	return Abs(p.R-p2.R) <= 1 && Abs(p.C-p2.C) <= 1
}
func (p *Position) MoveTo(p2 Position) {
	p.R = p2.R
	p.C = p2.C
}
func (p *Position) Up() *Position {
	return &Position{p.R + 1, p.C}
}
func (p *Position) Down() *Position {
	return &Position{p.R - 1, p.C}
}
func (p *Position) Equals(p2 Position) bool {
	return p.R == p2.R && p.C == p2.C
}
func (p *Position) UpR() *Position {
	return &Position{p.R - 1, p.C}
}
func (p *Position) DownR() *Position {
	return &Position{p.R + 1, p.C}
}
func (p *Position) Left() *Position {
	return &Position{p.R, p.C - 1}
}
func (p *Position) Right() *Position {
	return &Position{p.R, p.C + 1}
}
func (p *Position) ForEachUntil(c Position, f func(int, int)) {
	if p.R == c.R {
		a, b := Min(p.C, c.C), Max(p.C, c.C)
		for a <= b {
			f(p.R, a)
			a++
		}
		return
	} else if p.C == c.C {
		a, b := Min(p.R, c.R), Max(p.R, c.R)
		for a <= b {
			f(a, p.C)
			a++
		}
		return
	}
	panic("not for diagonally positioned")
}
func (p *Position) MHDist(c *Position) int {
	return Abs(p.R-c.R) + Abs(p.C-c.C)
}

func SliceStrToInt(s []string) (ints []int) {
	for _, v := range s {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		ints = append(ints, val)
	}
	return
}

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type BoolMap map[int]map[int]bool

func (mapa BoolMap) AssignToMap(r int, c int) {
	if _, ok := mapa[r]; !ok {
		mapa[r] = make(map[int]bool)
	}
	mapa[r][c] = true
}
