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
	nexts := [][5]int{[5]int{0, 0, 0, 0, 0}, [5]int{0, 0, 3, 0, 0}}
	seen[([4]int)(nexts[0][:4])] = 0
	seen[([4]int)(nexts[1][:4])] = 0
	for len(nexts) != 0 {
		next := ([4]int)(nexts[0][:4])
		totalHeatLoss := nexts[0][4]
		nexts = nexts[1:]
		coords, _ := graph[next]
		for _, coord := range coords {
			_, inBounds := graph[coord]
			if !inBounds {
				continue
			}
			n, _ := strconv.Atoi(string(lines[coord[0]][coord[1]]))
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
	return strconv.Itoa(res)
}

func h(d int, i int, j int) (int) {
	return d - i - j
}

func Resolve2(lines []string) (string) {
	res := 0
	graph := map[[4]int][][4]int{}
	seen := map[[4]int]int{}
	for i, line := range lines {
		for j, _ := range line {
			for k := 0; k < 10; k++ {
				if k >= 3 && k < 9 {
					graph[[4]int{i, j, 0, k}] = [][4]int{[4]int{i-1, j, 1, 0}, [4]int{i+1, j, 3, 0}, [4]int{i, j+1, 0, k+1}}
					graph[[4]int{i, j, 1, k}] = [][4]int{[4]int{i, j-1, 2, 0}, [4]int{i, j+1, 0, 0}, [4]int{i-1, j, 1, k+1}}
					graph[[4]int{i, j, 2, k}] = [][4]int{[4]int{i-1, j, 1, 0}, [4]int{i+1, j, 3, 0}, [4]int{i, j-1, 2, k+1}}
					graph[[4]int{i, j, 3, k}] = [][4]int{[4]int{i, j-1, 2, 0}, [4]int{i, j+1, 0, 0}, [4]int{i+1, j, 3, k+1}}
				} else if k < 3 {
					graph[[4]int{i, j, 0, k}] = [][4]int{[4]int{i, j+1, 0, k+1}}
					graph[[4]int{i, j, 1, k}] = [][4]int{[4]int{i-1, j, 1, k+1}}
					graph[[4]int{i, j, 2, k}] = [][4]int{[4]int{i, j-1, 2, k+1}}
					graph[[4]int{i, j, 3, k}] = [][4]int{[4]int{i+1, j, 3, k+1}}
				} else {
					graph[[4]int{i, j, 0, k}] = [][4]int{[4]int{i-1, j, 1, 0}, [4]int{i+1, j, 3, 0}}
					graph[[4]int{i, j, 1, k}] = [][4]int{[4]int{i, j-1, 2, 0}, [4]int{i, j+1, 0, 0}}
					graph[[4]int{i, j, 2, k}] = [][4]int{[4]int{i-1, j, 1, 0}, [4]int{i+1, j, 3, 0}}
					graph[[4]int{i, j, 3, k}] = [][4]int{[4]int{i, j-1, 2, 0}, [4]int{i, j+1, 0, 0}}
				}
			}
		}
	}
	nexts := [][5]int{[5]int{0, 0, 0, 0, 0}, [5]int{0, 0, 3, 0, 0}}
	seen[([4]int)(nexts[0][:4])] = 0
	seen[([4]int)(nexts[1][:4])] = 0
	for len(nexts) != 0 {
		next := ([4]int)(nexts[0][:4])
		totalHeatLoss := nexts[0][4]
		nexts = nexts[1:]
		coords, _ := graph[next]
		for _, coord := range coords {
			_, inBounds := graph[coord]
			if !inBounds {
				continue
			}
			n, _ := strconv.Atoi(string(lines[coord[0]][coord[1]]))
			heatLoss, exists := seen[coord]
			if exists && totalHeatLoss + n >= heatLoss {
				continue
			}
			seen[coord] = totalHeatLoss + n
			if res > 0 && res <= totalHeatLoss + n {
				continue
			}
			if coord[0] == len(lines) - 1 && coord[1] == len(lines[0]) - 1 && coord[3] >= 3 {
				if res == 0 || totalHeatLoss + n < res {
					res = totalHeatLoss + n
				}
				continue
			}
			i := 0
			for i < len(nexts) && h(nexts[i][4], nexts[i][0], nexts[i][1]) < h(totalHeatLoss + n, coord[0], coord[1]) {
				i += 1
			}
			if i < len(nexts) {
				nexts = append(nexts[:i+1], nexts[i:]...)
				nexts[i] = ([5]int)(append(([]int)(coord[:]), totalHeatLoss + n))
			} else {
				nexts = append(nexts, ([5]int)(append(([]int)(coord[:]), totalHeatLoss + n)))
			}
		}
	}
	return strconv.Itoa(res)
}

