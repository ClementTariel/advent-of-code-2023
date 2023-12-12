package main

import (
	"strconv"
)

func Insert[T any](l []T,i int, e T) ([]T) {
	n := len(l)
	i = (i%n + n) % n
	if i == n {
		return append(l, e)
	} else {
		l = append(l[:i+1], l[i:]...)
		l[i] = e
		return l
	}
}

func Abs(n int) (int) {
	if n < 0 {
		return -n
	}
	return n
}

func ManhattanDist(c1 [2]int, c2 [2]int) (int) {
	return Abs(c1[0] - c2[0]) + Abs(c1[1] - c2[1])
}

func Resolve1(lines []string) (string) {
	res := 0
	space := [][]string{}
	for _, line := range lines {
		space = append(space, []string{})
		for _, c := range line {
			l := string(c)
			space[len(space)-1] = append(space[len(space)-1], l)
		}
	}
	emptyLines := []int{}
	emptyColumns := []int{}
	for i, _ := range space {
		isEmpty := true
		for _, l := range space[i] {
			if l == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyLines = append(emptyLines, i)
		}
	}
	for j, _ := range space[0] {
		isEmpty := true
		for _, l := range space {
			if l[j] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyColumns = append(emptyColumns, j)
		}
	}
	for i := len(emptyLines)-1 ; i >= 0 ; i-- {
		newLine := []string{}
		k := emptyLines[i]
		for j := 0 ; j < len(space[k]) ; j++ {
			newLine = append(newLine, ".")
		}
		space = Insert(space, k, newLine)
	}
	for j := len(emptyColumns)-1 ; j >= 0 ; j-- {
		k := emptyColumns[j]
		for i, _ := range space {
			space[i] = Insert(space[i], k, ".")
		}
	}
	coords := [][2]int{}
	for i, _ := range space {
		for j, _ := range space[i] {
			if space [i][j] == "#" {
				coords = append(coords, [2]int{i,j})
			}
		}
	}
	for _, coord1 := range coords {
		for _, coord2 := range coords {
			res += ManhattanDist(coord1, coord2)
		}
	}
	res = res >> 1
	return strconv.Itoa(res)
}

func Expand(coord [2]int, emptyLines []int, emptyColumns []int) ([2]int) {
	di := 0
	for _, i := range emptyLines {
		if i < coord[0] {
			di += 1
		}
	}
	di *= 1000000 - 1
	dj := 0
	for _, j := range emptyColumns {
		if j < coord[1] {
			dj += 1
		}
	}
	dj *= 1000000 - 1
	return [2]int{coord[0]+di, coord[1]+dj}
}

func Resolve2(lines []string) (string) {
	res := 0
	space := [][]string{}
	for _, line := range lines {
		space = append(space, []string{})
		for _, c := range line {
			l := string(c)
			space[len(space)-1] = append(space[len(space)-1], l)
		}
	}
	emptyLines := []int{}
	emptyColumns := []int{}
	for i, _ := range space {
		isEmpty := true
		for _, l := range space[i] {
			if l == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyLines = append(emptyLines, i)
		}
	}
	for j, _ := range space[0] {
		isEmpty := true
		for _, l := range space {
			if l[j] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyColumns = append(emptyColumns, j)
		}
	}
	coords := [][2]int{}
	for i, _ := range space {
		for j, _ := range space[i] {
			if space [i][j] == "#" {
				coords = append(coords, Expand([2]int{i,j}, emptyLines, emptyColumns))
			}
		}
	}
	for _, coord1 := range coords {
		for _, coord2 := range coords {
			res += ManhattanDist(coord1, coord2)
		}
	}
	res = res >> 1
	return strconv.Itoa(res)
}
