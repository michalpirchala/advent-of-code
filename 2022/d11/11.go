package d11

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	mod     = 0
	rounds  = 20
	divider = 3
)

func P1() {
	eleven()
}

func P2() {
	rounds = 10000
	divider = 1
	eleven()
}

func eleven() {
	file, _ := os.Open("inputs/11.txt")
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	monkeys := make(map[int]*monkey)

	m := 0
	for scan.Scan() {
		s := scan.Text()

		if len(s) > 0 && s[0] == 'M' {
			lm := newMonkey()

			// items
			scan.Scan()
			s = scan.Text()
			r := strings.Split(s, ": ")
			items := strings.Split(r[1], ", ")
			for _, item := range items {
				itemI, _ := strconv.Atoi(item)
				lm.items.add(monkeyItem(itemI))
			}

			// operation
			scan.Scan()
			s = scan.Text()
			r = strings.Split(s, " ")
			lm.operation = operationFunction(r[6], r[7])

			//test
			scan.Scan()
			s = scan.Text()
			r = strings.Split(s, "by ")
			v, _ := strconv.Atoi(r[1])
			lm.testDiv = v

			//true
			scan.Scan()
			s = scan.Text()
			r = strings.Split(s, "monkey ")
			v, _ = strconv.Atoi(r[1])
			lm.testTrueMonkey = v

			//false
			scan.Scan()
			s = scan.Text()
			r = strings.Split(s, "monkey ")
			v, _ = strconv.Atoi(r[1])
			lm.testFalseMonkey = v

			if m == 0 {
				mod = lm.testDiv
			} else {
				mod *= lm.testDiv
			}

			monkeys[m] = &lm
			m++
		}
	}
	file.Close()

	for r := 0; r < rounds; r++ {
		for i := 0; i < m; i++ {
			for !monkeys[i].items.empty() {
				item := monkeys[i].inspectItem()
				next := monkeys[i].nextMonkey(item)
				monkeys[next].items.add(item)
			}
		}
		fmt.Printf("round %v inspections %v\n", r+1, getInspections(monkeys))
	}

	inspections := getInspections(monkeys)
	sort.Ints(inspections)
	fmt.Println(inspections[m-1] * inspections[m-2])
}

type monkeyItems []monkeyItem
type monkeyItem int

func (is *monkeyItems) add(i monkeyItem) {
	*is = append(*is, i)
}

func (is *monkeyItems) get() monkeyItem {
	item := (*is)[0]
	*is = (*is)[1:len(*is)]
	return item
}

func (is *monkeyItems) empty() bool {
	return len(*is) == 0
}

type monkey struct {
	items           *monkeyItems
	operation       func(item monkeyItem) monkeyItem
	testDiv         int
	testTrueMonkey  int
	testFalseMonkey int
	inspections     int
}

func (m *monkey) inspectItem() monkeyItem {
	m.inspections++
	item := m.items.get()
	item = m.operation(item)
	item = item % monkeyItem(mod)
	item /= monkeyItem(divider)
	return item
}

func (m *monkey) nextMonkey(item monkeyItem) int {
	if m.test(item) {
		return m.testTrueMonkey
	} else {
		return m.testFalseMonkey
	}
}

func (m *monkey) test(item monkeyItem) bool {
	return int(item)%int(m.testDiv) == 0
}

func newMonkey() monkey {
	m := monkey{}
	m.items = &monkeyItems{}
	return m
}

func getInspections(monkeys map[int]*monkey) []int {
	m := len(monkeys)
	inspections := make([]int, m)
	for i := 0; i < m; i++ {
		inspections[i] = monkeys[i].inspections
	}

	return inspections
}

func operationFunction(op, value string) func(item monkeyItem) monkeyItem {
	return func(old monkeyItem) monkeyItem {
		var v monkeyItem
		if value == "old" {
			v = old
		} else {
			i, _ := strconv.Atoi(value)
			v = monkeyItem(i)
		}
		switch op {
		case "+":
			return old + v
		case "-":
			return old - v
		case "*":
			return old * v
		case "/":
			return old / v
		}
		panic("no case defined in operation for " + op)
	}
}
