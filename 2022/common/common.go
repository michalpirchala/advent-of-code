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
func (p *Position) Up() Position {
	return Position{p.R + 1, p.C}
}
func (p *Position) Down() Position {
	return Position{p.R - 1, p.C}
}
func (p *Position) Equals(p2 Position) bool {
	return p.R == p2.R && p.C == p2.C
}
func (p *Position) UpR() Position {
	return Position{p.R - 1, p.C}
}
func (p *Position) DownR() Position {
	return Position{p.R + 1, p.C}
}
func (p *Position) Left() Position {
	return Position{p.R, p.C - 1}
}
func (p *Position) Right() Position {
	return Position{p.R, p.C + 1}
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
