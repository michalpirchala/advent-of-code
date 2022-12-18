package d16

import (
	"aoc2022/common"
	"fmt"
	"sort"
	"strings"
)

type valve struct {
	rate    int
	leadsTo []string
}

type sim struct {
	time      int
	opened    map[string]bool
	valve     string
	pressure  int
	pressureT int
}

func P1() {
	valves := make(map[string]valve, 0)
	q := make([]sim, 0)
	v := 0
	common.FileIter("16", func(s string) {
		name, valve := scan(s)
		valves[name] = valve
		if v == 0 {
			q = append(q, sim{
				time:   30,
				opened: make(map[string]bool),
				valve:  name,
			})
		}
		v++
	})

	ps := make([]int, 0)
	for len(q) > 0 {
		a := q[0]
		q = q[1:]

		//check for time
		if a.time <= 0 {
			ps = append(ps, a.pressureT)
			continue
		}

		// skip if actual is opened
		if a.opened[a.valve] || valves[a.valve].rate == 0 {
			//move to others without opening the valve
			for _, next := range valves[a.valve].leadsTo {
				q = append(q, sim{
					time:      a.time - 1,
					opened:    copyOpened(a.opened),
					valve:     next,
					pressure:  a.pressure,
					pressureT: a.pressureT + a.pressure,
				})
			}
			continue
		}

		// open actual valve
		a.time -= 1
		a.pressureT += a.pressure
		a.opened[a.valve] = true
		a.pressure += valves[a.valve].rate

		//fmt.Println(a.pressure)

		if a.time == 0 {
			q = append(q, a)
			continue
		}

		for _, next := range valves[a.valve].leadsTo {
			q = append(q, sim{
				time:      a.time - 1,
				opened:    copyOpened(a.opened, a.valve),
				valve:     next,
				pressure:  a.pressure,
				pressureT: a.pressureT + a.pressure,
			})
		}
		//fmt.Println(len(q))
	}
	sort.Ints(ps)
	fmt.Println(ps[len(ps)-1])
}

func scan(l string) (string, valve) {
	sep := strings.Split(l, "; tunnels lead to valves ")
	if len(sep) == 1 {
		sep = strings.Split(l, "; tunnel leads to valve ")
	}
	val := strings.Split(sep[0], " has flow rate=")
	valves := strings.Split(sep[1], ", ")
	name := strings.Split(val[0], " ")
	return name[1], valve{
		rate:    common.StrToInt(val[1]),
		leadsTo: valves,
	}
}

func copyOpened(o map[string]bool, valves ...string) map[string]bool {
	openedCopy := make(map[string]bool)
	for s, _ := range o {
		openedCopy[s] = true
	}
	for _, s := range valves {
		openedCopy[s] = true
	}
	return openedCopy
}

func P2() {

}
