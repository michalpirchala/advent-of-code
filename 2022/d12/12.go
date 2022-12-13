package d12

import (
	"aoc2022/common"
	"fmt"
)

type queuePos struct {
	s int //steps
	p common.Position
}

func P1() {
	mapa, start, point := loadMap()

	explored := make(map[common.Position]bool)
	q := []queuePos{{
		s: 0,
		p: start,
	}}
	explored[start] = true

	step := 0
	for len(q) > 0 {
		qp := q[0]
		c := qp.p //current
		q = q[1:]

		directions := []common.Position{c.UpR(), c.DownR(), c.Left(), c.Right()}
		for _, dir := range directions {
			if dir.R < 0 || dir.C < 0 || dir.R >= len(mapa) || dir.C >= len(mapa[0]) || explored[dir] {
				continue
			}
			if int(mapa[dir.R][dir.C])-int(mapa[c.R][c.C]) > 1 {
				continue
			}
			if dir.Equals(point) {
				fmt.Println(qp.s + 1)
				return
			}
			explored[dir] = true
			q = append(q, queuePos{qp.s + 1, dir})
		}
		step++
	}

	panic("shouldn't happen")
}

func P2() {
	mapa, _, point := loadMap()

	explored := make(map[common.Position]bool)
	q := []queuePos{{
		s: 0,
		p: point,
	}}
	explored[point] = true

	step := 0
	for len(q) > 0 {
		qp := q[0]
		c := qp.p //current
		q = q[1:]

		directions := []common.Position{c.UpR(), c.DownR(), c.Left(), c.Right()}
		for _, dir := range directions {
			if dir.R < 0 || dir.C < 0 || dir.R >= len(mapa) || dir.C >= len(mapa[0]) || explored[dir] {
				continue
			}
			if int(mapa[c.R][c.C])-int(mapa[dir.R][dir.C]) > 1 {
				continue
			}
			if mapa[dir.R][dir.C] == 'a' {
				fmt.Println(qp.s + 1)
				return
			}
			explored[dir] = true
			q = append(q, queuePos{qp.s + 1, dir})
		}
		step++
	}

	panic("shouldn't happen")
}

func loadMap() (map[int]map[int]uint8, common.Position, common.Position) {
	mapa := make(map[int]map[int]uint8)
	row := 0
	var start, point common.Position
	common.FileIter("12", func(s string) {
		mapa[row] = make(map[int]uint8)
		for col, c := range s {
			z := uint8(c)
			if z == 'S' {
				z = 'a'
				start.R = row
				start.C = col
			} else if z == 'E' {
				z = 'z'
				point.R = row
				point.C = col
			}
			mapa[row][col] = z
		}
		row++
	})
	return mapa, start, point
}

func drawMap(mapa map[int]map[int]uint8, explored map[common.Position]bool, c common.Position) {
	fmt.Println("===========")
	for i := 0; i < len(mapa); i++ {
		for j := 0; j < len(mapa[0]); j++ {
			a := common.Position{i, j}
			if c.Equals(a) {
				fmt.Print("#")
			} else if explored[a] {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", mapa[i][j])
			}
		}
		fmt.Println("")
	}
	fmt.Println("====")
}
