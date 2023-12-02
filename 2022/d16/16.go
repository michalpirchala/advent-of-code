package d16

import (
	"aoc2022/common"
	"fmt"
	"strconv"
	"strings"
)

type valve struct {
	rate    int
	leadsTo []string
}

type distance struct {
	distance int
	valve    string
}

var (
	valves   = make(map[string]valve, 0)
	dists    = make(map[string]map[string]int)
	nonempty = make([]string, 0)
	indices  = make(map[string]int)
	cache    = make(map[string]int)
)

// P1 basically rewrites hyper-neutrino's excellent solution
// source: https://www.youtube.com/watch?v=bLMj50cpOug&t=131s
func P1() {
	common.FileIter("16", func(s string) {
		name, valve := scan(s)
		valves[name] = valve
	})

	for valve := range valves {
		if valve != "AA" && valves[valve].rate == 0 {
			continue
		}
		if valve != "AA" {
			nonempty = append(nonempty, valve)
		}
		dists[valve] = map[string]int{
			"AA":  0,
			valve: 0,
		}
		visited := make(map[string]bool)

		q := make([]distance, 0)
		q = append(q, distance{distance: 0, valve: valve})

		for len(q) > 0 {
			a := q[0]
			q = q[1:]
			for _, neighbor := range valves[a.valve].leadsTo {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				if valves[neighbor].rate > 0 {
					dists[valve][neighbor] = a.distance + 1
				}
				q = append(q, distance{distance: a.distance + 1, valve: neighbor})
			}
		}

		if valve != "AA" {
			delete(dists[valve], "AA")
		}
		delete(dists[valve], valve)
	}

	fmt.Println(dists)

	for i, v := range nonempty {
		indices[v] = i
	}

	fmt.Println(dfs(30, "AA", 0))

	b := (1 << len(nonempty)) - 1
	m := 0
	for i := 0; i < (b+1)/2; i++ {
		a := dfs(26, "AA", i) + dfs(26, "AA", b^i)
		if a > m {
			m = a
		}
	}
	fmt.Println(m)
}

func dfs(time int, valve string, bitmask int) int {
	cacheKey := strconv.Itoa(time) + valve + strconv.Itoa(bitmask)
	if val, ok := cache[cacheKey]; ok {
		return val
	}
	maxval := 0
	for neighbor := range dists[valve] {
		bit := 1 << indices[neighbor]
		if bitmask&bit > 0 {
			continue
		}
		remtime := time - dists[valve][neighbor] - 1
		if remtime <= 0 {
			continue
		}
		maxval = common.Max(maxval, dfs(remtime, neighbor, bitmask|bit)+valves[neighbor].rate*remtime)
	}
	cache[cacheKey] = maxval
	return maxval
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
