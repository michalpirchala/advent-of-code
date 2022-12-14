package d14

import (
	"aoc2022/common"
	"fmt"
	"strings"
)

type BoolMap map[int]map[int]bool

func (mapa BoolMap) AssignToMap(r int, c int) {
	if _, ok := mapa[r]; !ok {
		mapa[r] = make(map[int]bool)
	}
	mapa[r][c] = true
}

func P1() {
	rockPos := make(map[int]int)
	mapa := make(BoolMap)
	common.FileIter("14", func(s string) {
		dirs := strings.Split(s, " -> ")

		for i := 0; i < len(dirs)-1; i++ {
			d := strings.Split(dirs[i], ",")
			p1 := common.Position{R: common.StrToInt(d[1]), C: common.StrToInt(d[0])}
			d = strings.Split(dirs[i+1], ",")
			p2 := common.Position{R: common.StrToInt(d[1]), C: common.StrToInt(d[0])}

			p1.ForEachUntil(p2, func(r int, c int) {
				mapa.AssignToMap(r, c)
				if rockPos[c] < r {
					rockPos[c] = r
				}
			})
		}
	})

	n := 0
SandLoop:
	for {
		p := &common.Position{
			R: 0,
			C: 500,
		}
		for {
			next := fall(mapa, p)
			if next == nil {
				break
			}
			if rockPos[next.C] < next.R {
				break SandLoop
			}
			p = next
		}
		mapa.AssignToMap(p.R, p.C)
		n++
	}
	fmt.Println(n)
}

func fall(mapa BoolMap, p *common.Position) *common.Position {
	dirs := []*common.Position{p.DownR(), p.DownR().Left(), p.DownR().Right()}
	for _, dir := range dirs {
		if !mapa[dir.R][dir.C] {
			return dir
		}
	}
	return nil
}

func P2() {
	mapa := make(BoolMap)
	h := 0
	common.FileIter("14", func(s string) {
		dirs := strings.Split(s, " -> ")

		for i := 0; i < len(dirs)-1; i++ {
			d := strings.Split(dirs[i], ",")
			p1 := common.Position{R: common.StrToInt(d[1]), C: common.StrToInt(d[0])}
			d = strings.Split(dirs[i+1], ",")
			p2 := common.Position{R: common.StrToInt(d[1]), C: common.StrToInt(d[0])}

			p1.ForEachUntil(p2, func(r int, c int) {
				mapa.AssignToMap(r, c)
				if h < r {
					h = r
				}
			})
		}
	})
	h += 2

	n := 0
SandLoop:
	for {
		p := &common.Position{
			R: 0,
			C: 500,
		}
		for {
			next := fall2(mapa, p, h)
			if next == nil {
				break
			}
			p = next
		}
		mapa.AssignToMap(p.R, p.C)
		n++
		if p.R == 0 && p.C == 500 {
			break SandLoop
		}
	}
	fmt.Println(n)
}

func fall2(mapa BoolMap, p *common.Position, h int) *common.Position {
	dirs := []*common.Position{p.DownR(), p.DownR().Left(), p.DownR().Right()}
	for _, dir := range dirs {
		if !mapa[dir.R][dir.C] && dir.R < h {
			return dir
		}
	}
	return nil
}
