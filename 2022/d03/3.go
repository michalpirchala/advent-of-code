package d03

import (
	"aoc2022/common"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func P1() {
	points := 0

	common.FileIter("3", func(s string) {
		l := len(s)
		for i := 0; i < l/2; i++ {
			c := s[i]
			for j := l / 2; j < l; j++ {
				c2 := s[j]
				if c == c2 {
					prior := prio(int(c))
					points += prior
					return
				}
			}
		}
	})

	fmt.Println(points)
}

func prio(c int) int {
	if c >= 'A' && c <= 'Z' {
		return c - 38
	}
	return c - 96
}

func P2() {
	file, _ := os.Open("inputs/3.txt")
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	body := 0

	for scan.Scan() {
		l1 := scan.Text()
		scan.Scan()
		l2 := scan.Text()
		scan.Scan()
		l3 := scan.Text()
	out:
		for i := 0; i < len(l1); i++ {
			c := l1[i]
			if strings.Contains(l2, string(c)) && strings.Contains(l3, string(c)) {
				body += prio(int(c))
				break out
			}
		}
	}
	file.Close()
	fmt.Println(body)
}
