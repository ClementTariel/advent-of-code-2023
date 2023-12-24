package main

import (
	"strconv"
	"strings"
)

func DiggingVector(dir string, n int) [2]int {
	if dir == "R" {
		return [2]int{0, n}
	} else if dir == "L" {
		return [2]int{0, -n}
	} else if dir == "U" {
		return [2]int{-n, 0}
	} else {
		return [2]int{n, 0}
	}
}

func Resolve1(lines []string) (string) {
	res := 0
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	x := 0
	y := 0
	ground := [][]string{[]string{"#"}}
	for _, line := range lines {
		data := strings.Split(line, " ")
		dir := data[0]
		n, _ := strconv.Atoi(data[1])
		v := DiggingVector(dir, n)
		if x + v[0] > maxX {
			for i := maxX; i < x + v[0]; i++{
				maxX += 1
				newLine := []string{}
				for j := 0; j < len(ground[0]); j++ {
					newLine = append(newLine, ".")
				}
				ground = append(ground, newLine)
			}
		}
		if x + v[0] < minX {
			for i := minX; i > x + v[0]; i--{
				minX -= 1
				newLine := []string{}
				for j := 0; j < len(ground[0]); j++ {
					newLine = append(newLine, ".")
				}
				ground = append(ground[:1], ground...)
				ground[0] = newLine
			}
		}
		if y + v[1] > maxY {
			for j := maxY; j < y + v[1]; j++{
				maxY += 1
				for i := 0; i < len(ground); i++ {
					ground[i] = append(ground[i], ".")
				}
			}
		}
		if y + v[1] < minY {
			for j := minY; j > y + v[1]; j--{
				minY -= 1
				for i := 0; i < len(ground); i++ {
					ground[i] = append(ground[i][:1], ground[i]...)
					ground[i][0] = "."
				}
			}
		}
		xSignPos := v[0] > 0
		ySignPos := v[1] > 0
		if xSignPos {
			for i := 0; i < v[0]; i++ {
				x += 1
				ground[x-minX][y-minY] = "#"
			}
		} else {
			for i := 0; i > v[0]; i-- {
				x -= 1
				ground[x-minX][y-minY] = "#"
			}
		}
		if ySignPos {
			for j := 0; j < v[1]; j++ {
				y += 1
				ground[x-minX][y-minY] = "#"
				res += 1
			}
		} else {
			for j := 0; j > v[1]; j-- {
				y -= 1
				ground[x-minX][y-minY] = "#"
				res += 1
			}
		}
	}
	for i := 0; i < len(ground); i++ {
		ground[i] = append(ground[i][:1], ground[i]...)
		ground[i][0] = "."
		ground[i] = append(ground[i], ".")
	}
	newLine := []string{}
	for j := 0; j < len(ground[0]); j++ {
		newLine = append(newLine, ".")
	}
	ground = append(ground[:1], ground...)
	ground[0] = newLine
	ground = append(ground, newLine)
	count := -1
	x = 0
	y = 0
	seen := map[[2]int]bool{}
	nexts := [][2]int{[2]int{x, y}}
	for len(nexts) != 0 {
		count += 1
		next := nexts[0]
		nexts = nexts[1:]
		i := next[0]
		j := next[1]
		if i > 0 {
			adj := [2]int{i-1, j}
			_, exists := seen[adj]
			if !exists && ground[adj[0]][adj[1]] == "." {
			seen[adj] = true
				nexts = append(nexts, adj)
			}
		}
		if i < len(ground)-1 {
			adj := [2]int{i+1, j}
			_, exists := seen[adj]
			if !exists && ground[adj[0]][adj[1]] == "." {
				seen[adj] = true
				nexts = append(nexts, adj)
			}
		}
		if j > 0 {
			adj := [2]int{i, j-1}
			_, exists := seen[adj]
			if !exists && ground[adj[0]][adj[1]] == "." {
				seen[adj] = true
				nexts = append(nexts, adj)
			}
		}
		if j < len(ground[0])-1 {
			adj := [2]int{i, j+1}
			_, exists := seen[adj]
			if !exists && ground[adj[0]][adj[1]] == "." {
				seen[adj] = true
				nexts = append(nexts, adj)
			}
		}
	}
	res = len(ground)*len(ground[0]) - count
	return strconv.Itoa(res)
}

func abs(n int) (int) {
	if n < 0 {
		return -n
	}
	return n
}

func Resolve2(lines []string) (string) {
	res := 0
	x := 0
	y := 0
	p := 0
	coords := [][2]int{[2]int{0, 0}}
	for _, line := range lines {
		data := strings.Split(strings.Split(strings.Split(line, " ")[2], ")")[0], "#")[1]
		dir, _ := map[string]string{"0": "R", "1": "D", "2": "L", "3": "U"}[string(data[len(data)-1])]
		n, _ := strconv.ParseInt(data[:len(data)-1], 16, 64)
		v := DiggingVector(dir, int(n))
		x += v[0]
		y += v[1]
		p += abs(v[0]) + abs(v[1])
		if x != 0 || y != 0 {
			res += coords[len(coords)-1][0]*y - coords[len(coords)-1][1]*x
			coords = append(coords, [2]int{x, y})
		}
	}
	res = (abs(res) + p) >> 1
	res += 1 // corner
	return strconv.Itoa(res)
}


