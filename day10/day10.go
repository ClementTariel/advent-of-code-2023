package main

import (
	"strconv"
)

func connected(prevL string, l string, di int, dj int) (bool) {
	if (di == 1 && dj == 0) {
		return (prevL=="S"||prevL=="|"||prevL=="7"||prevL=="F")&&(l=="S"||l=="|"||l=="L"||l=="J")
	}
	if (di == -1 && dj == 0) {
		return (prevL=="S"||prevL=="|"||prevL=="L"||prevL=="J")&&(l=="S"||l=="|"||l=="7"||l=="F")
	}
	if (di == 0 && dj == 1) {
		return (prevL=="S"||prevL=="-"||prevL=="L"||prevL=="F")&&(l=="S"||l=="-"||l=="7"||l=="J")
	}
	if (di == 0 && dj == -1) {
		return (prevL=="S"||prevL=="-"||prevL=="7"||prevL=="J")&&(l=="S"||l=="-"||l=="L"||l=="F")
	}
	return false
}

func adjacent(i int, j int, ni int, nj int) ([][2]int){
	pipes := [][2]int{}
	if i > 0 {
		pipes = append(pipes, [2]int{i-1,j})
	}
	if i < ni - 1 {
		pipes = append(pipes, [2]int{i+1,j})
	}
	if j > 0 {
		pipes = append(pipes, [2]int{i,j-1})
	}
	if j < nj - 1 {
		pipes = append(pipes, [2]int{i,j+1})
	}
	return pipes
}

func Resolve1(lines []string) (string) {
	res := 0
	prevX := 0
	prevY := 0
	pipes := [][]string{}
	for i, line := range lines {
		pipes = append(pipes, []string{})
		for j, c := range line {
			l := string(c)
			pipes[len(pipes)-1] = append(pipes[len(pipes)-1], l)
			if l == "S" {
				prevX = i
				prevY = j
			}
		}
	}
	x := prevX
	y := prevY
	finish := false
	count := 0
	for !finish {
		for _, coord := range adjacent(x, y, len(pipes), len(pipes[x])) {
			i := coord[0]
			j := coord[1]
			if (i != prevX || j != prevY) && connected(pipes[x][y], pipes[i][j], i-x, j-y) {
				prevX = x
				prevY = y
				x = i
				y = j
				break
			}
		}
		count += 1
		if pipes[x][y] == "S" {
			finish = true
		}
	}
	res = count >> 1
	return strconv.Itoa(res)
}

func Resolve2(lines []string) (string) {
	res := 0
	prevX := 0
	prevY := 0
	pipes := [][]string{}
	hBorder := [][]bool{}
	hBorderCount := [][]int{}
	for i, line := range lines {
		pipes = append(pipes, []string{})
		hBorder = append(hBorder, []bool{})
		hBorderCount = append(hBorderCount, []int{})
		for j, c := range line {
			l := string(c)
			hBorder[len(hBorder)-1] = append(hBorder[len(hBorder)-1], false)
			hBorderCount[len(hBorderCount)-1] = append(hBorderCount[len(hBorderCount)-1], 0)
			pipes[len(pipes)-1] = append(pipes[len(pipes)-1], l)
			if l == "S" {
				prevX = i
				prevY = j
			}
		}
	}
	x := prevX
	y := prevY
	adj := adjacent(x, y, len(pipes), len(pipes[x]))
	adj0 := adj[0]
	adj1 := adj[1]
	di0 := x - adj0[0]
	dj0 := y - adj0[1]
	di1 := x - adj1[0]
	dj1 := y - adj1[1]
	if (dj0 == 1 && dj1 == -1) || (dj1 == 1 && dj0 == -1) {
		hBorder[x][y] = true
	}
	if (dj0 == 1 && di1 == 1) || (dj1 == 1 && di0 == 1) {
		hBorder[x][y] = true
	}
	if (dj0 == 1 && di1 == -1) || (dj1 == 1 && di0 == -1) {
		hBorder[x][y] = true
	}

	finish := false
	for !finish {
		for _, coord := range adjacent(x, y, len(pipes), len(pipes[x])) {
			i := coord[0]
			j := coord[1]
			l := pipes[i][j]
			if (i != prevX || j != prevY) && connected(pipes[x][y], l, i-x, j-y) {
				pipes[prevX][prevY] = "S"
				prevX = x
				prevY = y
				x = i
				y = j
				if l == "-" || l == "7" || l == "J" {
					hBorder[i][j] = true
				}
				break
			}
		}
		if pipes[x][y] == "S" {
			finish = true
		}
	}
	pipes[prevX][prevY] = "S"
	count := 0
	for i, _ := range hBorder {
		for j, _ := range hBorder[i] {
			hCount := 0
			if i > 0 {
				hCount = hBorderCount[i-1][j]
			}
			if hBorder[i][j] {
				hCount += 1
			}
			hBorderCount[i][j] = hCount
			if pipes[i][j] != "S" && hCount%2 == 1{
				count += 1
			}
		}
	}
	res = count
	return strconv.Itoa(res)
}

