package d02

import (
	"aoc2022/common"
	"fmt"
)

// this day I was showing programming to someone
// otherwise I wouldn't bother so much :D

const ROCK = 1
const PAPER = 2
const SCISS = 3

func P1() {
	points := 0

	common.FileIter("2", func(line string) {
		if line == "" {
			return
		}
		opp := z1(string(line[0]))
		me := z2(string(line[2]))

		if won(opp, me) {
			points += 6
		} else if opp == me {
			points += 3
		}

		points += me
	})

	fmt.Println(points)
}

func won(opp int, me int) bool {
	return (me == ROCK && opp == SCISS) ||
		(me == PAPER && opp == ROCK) ||
		(me == SCISS && opp == PAPER)
}

func z1(s string) int {
	switch s {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISS
	}
	panic("unknown")
}

func z2(s string) int {
	switch s {
	case "X":
		return ROCK
	case "Y":
		return PAPER
	case "Z":
		return SCISS
	}
	panic("unknown")
}

func P2() {
	points := 0

	common.FileIter("2", func(line string) {
		if line == "" {
			return
		}
		opp := z1(string(line[0]))
		me := coMamDac(opp, string(line[2]))

		if won(opp, me) {
			points += 6
		} else if opp == me {
			points += 3
		}

		points += me
	})

	fmt.Println(points)
}

func coMamDac(opp int, v string) int {
	znaky := []int{ROCK, PAPER, SCISS}
	if v == "Y" { // remiza
		return opp
	}

	if v == "X" { // prehrac
		return znaky[(opp-1-1+3)%3]
	}

	if v == "Z" { // vyhrac
		return znaky[(opp)%3]
	}

	panic("shouldn't happen")
}
