package d05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []uint8

func (s *stack) Pop() (uint8, bool) {
	if len(*s) == 0 {
		return '0', false
	}
	i := len(*s) - 1
	val := (*s)[i]
	*s = (*s)[:i]
	return val, true
}

func (s *stack) Push(v uint8) {
	*s = append(*s, v)
}

func P1() {
	file, _ := os.Open("inputs/5.txt")
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	stackNum := 0
	var stacks []stack
	for scan.Scan() {
		l := scan.Text()
		if stackNum == 0 {
			stackNum = (len(l) + 1) / 4
			stacks = make([]stack, stackNum)
		}
		if l[1] == '1' {
			scan.Scan()
			break
		}

		stackI := 0
		for i := 1; i < len(l); i += 4 {
			c := l[i]
			if c != ' ' {
				stacks[stackI].Push(c)
			}
			stackI++
		}
	}

	//reverse stacks
	for i := 0; i < stackNum; i++ {
		var st stack
		for {
			v, ok := stacks[i].Pop()
			if !ok {
				break
			}
			st.Push(v)
		}
		stacks[i] = st
	}

	for scan.Scan() {
		l := scan.Text()
		if len(l) == 0 {
			break
		}
		w := strings.Split(l, " ")
		num, _ := strconv.Atoi(w[1])
		from, _ := strconv.Atoi(w[3])
		to, _ := strconv.Atoi(w[5])

		for i := 0; i < num; i++ {
			val, ok := stacks[from-1].Pop()
			if !ok {
				panic("wtf man")
			}
			stacks[to-1].Push(val)
		}
	}

	for i := 0; i < stackNum; i++ {
		val, _ := stacks[i].Pop()
		fmt.Print(string(val))
	}

	file.Close()
}

func P2() {
	file, _ := os.Open("inputs/5.txt")
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	stackNum := 0
	var stacks []stack
	for scan.Scan() {
		l := scan.Text()
		if stackNum == 0 {
			stackNum = (len(l) + 1) / 4
			stacks = make([]stack, stackNum)
		}
		if l[1] == '1' {
			scan.Scan()
			break
		}

		stackI := 0
		for i := 1; i < len(l); i += 4 {
			c := l[i]
			if c != ' ' {
				stacks[stackI].Push(c)
			}
			stackI++
		}
	}

	//reverse stacks
	for i := 0; i < stackNum; i++ {
		var st stack
		for {
			v, ok := stacks[i].Pop()
			if !ok {
				break
			}
			st.Push(v)
		}
		stacks[i] = st
	}

	for scan.Scan() {
		l := scan.Text()
		if len(l) == 0 {
			break
		}
		w := strings.Split(l, " ")
		num, _ := strconv.Atoi(w[1])
		from, _ := strconv.Atoi(w[3])
		to, _ := strconv.Atoi(w[5])

		var temp stack

		for i := 0; i < num; i++ {
			val, ok := stacks[from-1].Pop()
			if !ok {
				panic("wtf man")
			}
			temp.Push(val)
		}

		for i := 0; i < num; i++ {
			val, ok := temp.Pop()
			if !ok {
				panic("wtf man")
			}
			stacks[to-1].Push(val)
		}
	}

	for i := 0; i < stackNum; i++ {
		val, _ := stacks[i].Pop()
		fmt.Print(string(val))
	}

	file.Close()
}
