package d07

import (
	"aoc2022/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Dir struct {
	files map[string]int
	subs  map[string]*Dir
	Name  string
}

func (fs *Dir) getSize() int {
	size := 0
	for _, filesize := range fs.files {
		size += filesize
	}
	for _, sub := range fs.subs {
		size += sub.getSize()
	}
	return size
}

func newDir() *Dir {
	nd := Dir{}
	nd.files = make(map[string]int)
	nd.subs = make(map[string]*Dir)
	return &nd
}

func getActiveDir(path []string, fs *Dir) *Dir {
	if len(path) == 0 {
		return fs
	}
	actual := fs
	for _, dir := range path {
		actual = actual.subs[dir]
	}

	return actual
}

func buildFs() *Dir {
	var path []string
	fs := newDir()

	common.FileIter("7", func(s string) {
		ad := getActiveDir(path, fs)
		if strings.HasPrefix(s, "$ cd") {
			switch s[5:] {
			case "..":
				path = path[:len(path)-1]
			case "/":
			default:
				path = append(path, s[5:])

			}
			return
		}
		if strings.HasPrefix(s, "$ ls") {
			return
		}
		file := strings.Split(s, " ")

		if size, err := strconv.Atoi(file[0]); err == nil {
			ad.files[file[1]] = size
		} else {
			if _, ok := ad.subs[file[1]]; !ok {
				ad.subs[file[1]] = newDir()
			}

		}
	})
	return fs
}

func P1() {
	fs := buildFs()
	fmt.Println(sumLowerT(100000, fs))
}

func sumLowerT(i int, fs *Dir) int {
	sum := 0

	size := fs.getSize()
	if size < i {
		sum += size
	}

	for _, subdir := range fs.subs {
		sum += sumLowerT(i, subdir)
	}

	return sum
}

func P2() {
	fs := buildFs()
	need := 30000000 - (70000000 - fs.getSize())
	sizes := buildFilesizes(fs)
	sort.Ints(sizes)
	for _, size := range sizes {
		if size >= need {
			fmt.Println(size)
			return
		}
	}
}

func buildFilesizes(fs *Dir) []int {
	var sizes []int
	sizes = append(sizes, fs.getSize())
	for _, d := range fs.subs {
		subSizes := buildFilesizes(d)
		sizes = append(sizes, subSizes...)
	}
	return sizes
}
