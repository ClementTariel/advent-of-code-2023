package main

import (
	"strconv"
)

func Resolve1(lines []string) (string) {
	res := 0
	graph := map[[4]int][][4]int{}
	seen := map[[4]int]int{}
	for i, line := range lines {
		for j, _ := range line {
			for k := 0; k < 3; k++ {
				if k < 2 {
					graph[[4]int{i, j, 0, k}] = [][4]int{[4]int{i-1, j, 1, 0}, [4]int{i+1, j, 3, 0}, [4]int{i, j+1, 0, k+1}}
					graph[[4]int{i, j, 1, k}] = [][4]int{[4]int{i, j-1, 2, 0}, [4]int{i, j+1, 0, 0}, [4]int{i-1, j, 1, k+1}}
					graph[[4]int{i, j, 2, k}] = [][4]int{[4]int{i-1, j, 1, 0}, [4]int{i+1, j, 3, 0}, [4]int{i, j-1, 2, k+1}}
					graph[[4]int{i, j, 3, k}] = [][4]int{[4]int{i, j-1, 2, 0}, [4]int{i, j+1, 0, 0}, [4]int{i+1, j, 3, k+1}}
				} else {
					graph[[4]int{i, j, 0, k}] = [][4]int{[4]int{i-1, j, 1, 0}, [4]int{i+1, j, 3, 0}}
					graph[[4]int{i, j, 1, k}] = [][4]int{[4]int{i, j-1, 2, 0}, [4]int{i, j+1, 0, 0}}
					graph[[4]int{i, j, 2, k}] = [][4]int{[4]int{i-1, j, 1, 0}, [4]int{i+1, j, 3, 0}}
					graph[[4]int{i, j, 3, k}] = [][4]int{[4]int{i, j-1, 2, 0}, [4]int{i, j+1, 0, 0}}
				}
			}
		}
	}
	startHeat, _ := strconv.Atoi(string(lines[0][0]))
	nexts := [][5]int{[5]int{0, 0, 0, 0, 0}, [5]int{0, 0, 3, 0, 0}}
	seen[([4]int)(nexts[0][:4])] = 0
	seen[([4]int)(nexts[1][:4])] = 0
	for len(nexts) != 0 {
		next := ([4]int)(nexts[0][:4])
		totalHeatLoss := nexts[0][4]
		nexts = nexts[1:]
		n, _ := strconv.Atoi(string(lines[next[0]][next[1]]))
		coords, _ := graph[next]
		for _, coord := range coords {
			_, inBounds := graph[coord]
			if !inBounds {
				continue
			}
			heatLoss, exists := seen[coord]
			if exists {
				if totalHeatLoss + n >= heatLoss {
					continue
				}
			}
			seen[coord] = totalHeatLoss + n
			if res > 0 && res <= totalHeatLoss + n {
				continue
			}
			if coord[0] == len(lines) - 1 && coord[1] == len(lines[0]) - 1 {
				if res == 0 || totalHeatLoss + n < res {
					res = totalHeatLoss + n
				}
				continue
			}
			nexts = append(nexts, ([5]int)(append(([]int)(coord[:]), totalHeatLoss + n)))
		}
	}
	finalHeatLoss, _ := strconv.Atoi(string(lines[len(lines)-1][len(lines[0])-1]))
	res += finalHeatLoss - startHeat
	return strconv.Itoa(res)
}

