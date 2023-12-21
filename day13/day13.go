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
		vertical := []int{}
		horizontal := []int{}

		for i := 1; i < len(reflexions); i++ {
			horizontal = append(horizontal, i)
		}
		for j := 1; j < len(reflexions[0]); j++ {
			vertical = append(vertical, j)
		}

		for i, line := range reflexions {
			for k := len(vertical)-1; k >= 0; k-- {
				j0 := vertical[k]
				for j := 0; j0 > j && j0+j < len(reflexions[i]); j++ {
					if line[j0-j-1] != line[j0+j] {
						vertical = append(vertical[:k], vertical[k+1:]...)
						break
					}
				}
			}
		}
		for j := 0; j < len(reflexions[0]); j++ {
			for k := len(horizontal)-1; k >= 0; k-- {
				i0 := horizontal[k]
				for i := 0; i0 > i && i0+i < len(reflexions); i++ {
					if reflexions[i0-i-1][j] != reflexions[i0+i][j] {
						horizontal = append(horizontal[:k], horizontal[k+1:]...)
						break
					}
				}
			}
		}
		if len(vertical) == 1 {
			res += vertical[0]
		}
		if len(horizontal) == 1 {
			res += 100 * horizontal[0]
		}
		reflexions = nil
	}
	return strconv.Itoa(res)
}

func Resolve2(lines []string) (string) {
	res := 0

	reflexions := []string{}
	for lineCount, reflexionLine := range lines {
		if len(reflexionLine) > 0 {
			reflexions = append(reflexions, reflexionLine)
			if lineCount+1 < len(lines) {
				continue
			}
		}
		vertical := []int{}
		horizontal := []int{}
		vErrors := []int{0}
		hErrors := []int{0}

		for i := 1; i < len(reflexions); i++ {
			horizontal = append(horizontal, i)
			hErrors = append(hErrors, 0)
		}
		for j := 1; j < len(reflexions[0]); j++ {
			vertical = append(vertical, j)
			vErrors = append(vErrors, 0)
		}

		for i, line := range reflexions {
			for k := len(vertical)-1; k >= 0; k-- {
				j0 := vertical[k]
				for j := 0; j0 > j && j0+j < len(reflexions[i]); j++ {
					if line[j0-j-1] != line[j0+j] {
						vErrors[j0] += 1
					}
					if vErrors[j0] > 1 {
						vertical = append(vertical[:k], vertical[k+1:]...)
						break
					}
				}
			}
		}
		for j := 0; j < len(reflexions[0]); j++ {
			for k := len(horizontal)-1; k >= 0; k-- {
				i0 := horizontal[k]
				for i := 0; i0 > i && i0+i < len(reflexions); i++ {
					if reflexions[i0-i-1][j] != reflexions[i0+i][j] {
						hErrors[i0] += 1
					}
					if hErrors[i0] > 1 {
						horizontal = append(horizontal[:k], horizontal[k+1:]...)
						break
					}
				}
			}
		}
		for i, e := range hErrors {
			if e == 1 {
				res += 100 * i
				break
			}
		}
		for j, e := range vErrors {
			if e == 1 {
				res += j
				break
			}
		}
		reflexions = nil
	}
	return strconv.Itoa(res)
}

