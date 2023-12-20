package main

import (
	"strconv"
)

func Resolve1(lines []string) (string) {
	res := 0

	reflexions := []string{}
	for lineCount, reflexionLine := range lines {
		if len(reflexionLine) > 0 {
			reflexions = append(reflexions, reflexionLine)
			if lineCount+1 < len(lines) {
				continue
			}
		}
		horizontal := []int{}
		vertical := []int{}

		// yes I accdentally switched vertical and horizontal but whatever
		for i := 1; i < len(reflexions); i++ {
			vertical = append(vertical, i)
		}
		for j := 1; j < len(reflexions[0]); j++ {
			horizontal = append(horizontal, j)
		}

		for i, line := range reflexions {
			for k := len(horizontal)-1; k >= 0; k-- {
				j0 := horizontal[k]
				for j := 0; j0 > j && j0+j < len(reflexions[i]); j++ {
					if line[j0-j-1] != line[j0+j] {
						horizontal = append(horizontal[:k], horizontal[k+1:]...)
						break
					}
				}
			}
		}
		for j := 0; j < len(reflexions[0]); j++ {
			for k := len(vertical)-1; k >= 0; k-- {
				i0 := vertical[k]
				for i := 0; i0 > i && i0+i < len(reflexions); i++ {
					if reflexions[i0-i-1][j] != reflexions[i0+i][j] {
						vertical = append(vertical[:k], vertical[k+1:]...)
						break
					}
				}
			}
		}
		if len(horizontal) == 1 {
			res += horizontal[0]
		}
		if len(vertical) == 1 {
			res += 100 * vertical[0]
		}
		reflexions = nil
	}
	return strconv.Itoa(res)
}

