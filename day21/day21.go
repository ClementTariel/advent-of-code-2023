package main

import (
	"fmt"
	"strconv"
)

func printGarden(lines []string, dist map[[2]int]int, steps int) {
	for i, line := range lines {
		for j, c := range line {
			d, _ := dist[[2]int{i, j}]
			if d <= steps && (d+2)%2 == 0 {
				fmt.Printf("O")
			} else {
				fmt.Printf(string(c))
			}
		}
		fmt.Println()
	}
}

func Resolve1(lines []string) (string) {
	res := 0
	border := [][2]int{}
	dist := map[[2]int]int{}
	steps := 64
	for i, line := range lines {
		for j, c := range line {
			if string(c) == "#" {
				dist[[2]int{i, j}] = -1
			} else {
				dist[[2]int{i, j}] = steps+1
			}
			if string(c) == "S" {
				border = append(border, [2]int{i, j})
				dist[[2]int{i, j}] = 0
			}

		}
	}
	for d := 1; d <= steps; d++ {
		nextBorder := [][2]int{}
		for _, b := range border {
			for k := 0; k < 4; k++ {
				i := b[0] + (k%2)*(2-k)
				j := b[1] + ((k+1)%2)*(1-k)
				n, exists := dist[[2]int{i, j}]
				if exists && n > d {
					dist[[2]int{i, j}] = d
					nextBorder = append(nextBorder, [2]int{i, j})
				}
			}
		}
		border = nextBorder
	}
	for _, d := range dist {
		if d <= steps && (d+2)%2 == 0 {
			res += 1
		}
	}
	return strconv.Itoa(res)
}


