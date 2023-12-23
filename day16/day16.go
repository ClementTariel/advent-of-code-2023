package main

import (
	"strconv"
)

func Resolve1(lines []string) (string) {
	res := 0
	graph := map[[3]int][][3]int{}
	seen := map[[3]int]bool{}
	for i, line := range lines {
		for j, char := range line {
			mirror := string(char)
			if mirror == "/" {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i-1, j, 1}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i, j+1, 0}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i+1, j, 3}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i, j-1, 2}}
			} else if mirror == "\\" {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i+1, j, 3}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i, j-1, 2}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i-1, j, 1}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i, j+1, 0}}
			} else if mirror == "-" {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i, j+1, 0}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i, j-1, 2}, [3]int{i, j+1, 0}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i, j-1, 2}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i, j-1, 2}, [3]int{i, j+1, 0}}
			} else if mirror == "|" {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i-1, j, 1}, [3]int{i+1, j, 3}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i-1, j, 1}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i-1, j, 1}, [3]int{i+1, j, 3}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i+1, j, 3}}
			} else {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i, j+1, 0}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i-1, j, 1}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i, j-1, 2}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i+1, j, 3}}
			}
		}
	}
	nexts := [][3]int{[3]int{0, 0, 0}}
	seen[nexts[0]] = true
	res = 1
	for len(nexts) != 0 {
		next := nexts[0]
		nexts = nexts[1:]
		coords, _ := graph[next]
		for _, coord := range coords {
			_, inBounds := graph[coord]
			if !inBounds {
				continue
			}
			_, exists := seen[coord]
			if exists {
				continue
			}
			energized := false
			for k := 0; k < 4; k++ {
				_,  energized = seen[[3]int{coord[0], coord[1], k}]
				if energized {
					break
				}
			}
			if !energized {
				res += 1
			}
			seen[coord] = true
			nexts = append(nexts, coord)
		}
	}
	return strconv.Itoa(res)
}
func Resolve2(lines []string) (string) {
	res := 0
	graph := map[[3]int][][3]int{}
	for i, line := range lines {
		for j, char := range line {
			mirror := string(char)
			if mirror == "/" {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i-1, j, 1}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i, j+1, 0}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i+1, j, 3}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i, j-1, 2}}
			} else if mirror == "\\" {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i+1, j, 3}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i, j-1, 2}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i-1, j, 1}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i, j+1, 0}}
			} else if mirror == "-" {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i, j+1, 0}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i, j-1, 2}, [3]int{i, j+1, 0}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i, j-1, 2}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i, j-1, 2}, [3]int{i, j+1, 0}}
			} else if mirror == "|" {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i-1, j, 1}, [3]int{i+1, j, 3}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i-1, j, 1}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i-1, j, 1}, [3]int{i+1, j, 3}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i+1, j, 3}}
			} else {
				graph[[3]int{i, j, 0}] = [][3]int{[3]int{i, j+1, 0}}
				graph[[3]int{i, j, 1}] = [][3]int{[3]int{i-1, j, 1}}
				graph[[3]int{i, j, 2}] = [][3]int{[3]int{i, j-1, 2}}
				graph[[3]int{i, j, 3}] = [][3]int{[3]int{i+1, j, 3}}
			}
		}
	}
	starts := [][3]int{}
	for i, _ := range lines {
		starts = append(starts, [3]int{i, 0, 0})
		starts = append(starts, [3]int{i, len(lines[i])-1, 2})
	}
	for j, _ := range lines[0] {
		starts = append(starts, [3]int{0, j, 3})
		starts = append(starts, [3]int{len(lines), j, 1})
	}
	for _, start := range starts {
		seen := map[[3]int]bool{}
		nexts := [][3]int{start}
		seen[nexts[0]] = true
		count := 1
		for len(nexts) != 0 {
			next := nexts[0]
			nexts = nexts[1:]
			coords, _ := graph[next]
			for _, coord := range coords {
				_, inBounds := graph[coord]
				if !inBounds {
					continue
				}
				_, exists := seen[coord]
				if exists {
					continue
				}
				energized := false
				for k := 0; k < 4; k++ {
					_,  energized = seen[[3]int{coord[0], coord[1], k}]
					if energized {
						break
					}
				}
				if !energized {
					count += 1
				}
				seen[coord] = true
				nexts = append(nexts, coord)
			}
		}
		if count > res {
			res = count
		}
	}
	return strconv.Itoa(res)
}
