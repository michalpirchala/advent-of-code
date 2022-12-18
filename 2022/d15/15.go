package d15

import (
	"aoc2022/common"
	"fmt"
	"strings"
)

func P1() {
	const ROW = 2000000
	covered := make(map[int]struct{})
	bar := make(map[common.Position]struct{}) //bacons at row
	common.FileIter("15", func(s string) {
		com := strings.Split(s, ":")
		sensor := parseXY(com[0])
		beacon := parseXY(com[1])
		if beacon.R == ROW {
			bar[*beacon] = struct{}{}
		}
		vDist := common.Abs(sensor.R - ROW)
		mhDist := sensor.MHDist(beacon)
		rr := mhDist - vDist
		if rr < 0 {
			return
		}
		for i := sensor.C - rr; i <= sensor.C+rr; i++ {
			covered[i] = struct{}{}
		}
	})
	fmt.Println(len(covered) - len(bar))
}

func P2() {
	mapa := make(common.BoolMap)
	sensors := make(map[common.Position]int)
	const LIMIT = 4000000
	common.FileIter("15", func(s string) {
		com := strings.Split(s, ":")
		sensor := parseXY(com[0])
		beacon := parseXY(com[1])
		mapa.AssignToMap(beacon.R, beacon.C)
		mhDist := sensor.MHDist(beacon)
		sensors[*sensor] = mhDist
	})
	for row := 0; row <= LIMIT; row++ {
	PointSearch:
		for col := 0; col <= LIMIT; col++ {
			if mapa[row][col] {
				continue
			}
			for sensor, mhDist := range sensors {
				vDist := common.Abs(sensor.R - row)
				rr := mhDist - vDist
				if rr < 0 {
					continue
				}
				if col >= sensor.C-rr && col <= sensor.C+rr {
					col = sensor.C + rr
					continue PointSearch
				}
			}
			fmt.Println(col*4000000 + row)
			return
		}
	}
}

func parseXY(s string) *common.Position {
	com := strings.Split(s, ", ")
	x := strings.Split(com[0], "=")
	y := strings.Split(com[1], "=")
	return &common.Position{
		R: common.StrToInt(y[1]),
		C: common.StrToInt(x[1]),
	}
}
