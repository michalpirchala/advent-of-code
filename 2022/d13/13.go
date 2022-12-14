package d13

import (
	"aoc2022/common"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
)

type packet any

type packets []packet

func (p packets) Len() int {
	return len(p)
}

func (p packets) Less(i, j int) bool {
	less, _ := rightOrder(p[i], p[j])
	return less
}

func (p packets) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func P1() {
	file, _ := os.Open("inputs/13.txt")
	defer file.Close()
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	var sum, i int
	for scan.Scan() {
		s := scan.Bytes()
		if len(s) == 0 {
			continue
		}
		i++
		var l, r packet
		json.Unmarshal(s, &l)
		scan.Scan()
		json.Unmarshal(scan.Bytes(), &r)

		right, _ := rightOrder(l, r)
		if right {
			sum += i
		}
	}
	fmt.Println(sum)
}

func P2() {
	var sep1, sep2 packet
	json.Unmarshal([]byte("[[2]]"), &sep1)
	json.Unmarshal([]byte("[[6]]"), &sep2)
	lines := packets{sep1, sep2}
	common.FileIter("13", func(s string) {
		if len(s) == 0 {
			return
		}
		b := []byte(s)
		var l packet
		json.Unmarshal(b, &l)
		lines = append(lines, l)
	})
	sort.Sort(lines)
	s := 1
	for i, line := range lines {
		if reflect.DeepEqual(line, sep1) || reflect.DeepEqual(line, sep2) {
			s *= i + 1
		}
	}
	fmt.Println(s)
}

func rightOrder(l, r packet) (bool, bool) {
	switch lv := l.(type) {
	case float64:
		switch rv := r.(type) {
		case float64:
			if lv == rv {
				return true, false
			}
			return lv < rv, true
		case []any:
			return rightOrder([]any{lv}, r)
		}
	case []any:
		switch rv := r.(type) {
		case float64:
			return rightOrder(lv, []any{r})
		case []any:
			i := 0
			for {
				if len(lv) == 0 && len(rv) > 0 {
					return true, true
				}
				if len(lv) > 0 && len(rv) == 0 {
					return false, true
				}
				if len(lv) == 0 && len(rv) == 0 {
					return false, false
				}
				right, valid := rightOrder(lv[i], rv[i])
				if valid {
					return right, valid
				}
				lv = lv[1:]
				rv = rv[1:]
			}
			return true, true
		}
	}
	return true, false
}
