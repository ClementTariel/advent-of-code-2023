package main

import (
	"strconv"
	"strings"
)

func Resolve1(lines []string) (string) {
	res := 0
	seedsMap := [][]int{[]int{},[]int{}}
	for i, line := range lines {
		if i == 0 {
			numbers := strings.Split(strings.Split(line, ": ")[1], " ")
			for _, nAsStr := range numbers {
				n, _ := strconv.Atoi(nAsStr)
				seedsMap[0] = append(seedsMap[0], n)
				seedsMap[1] = append(seedsMap[1], n)
			}
		} else if len(line) > 0 {
			if strings.HasSuffix(line, ":") {
				for j, _ := range seedsMap[0] {
					seedsMap[0][j] = seedsMap[1][j]
				}
				continue
			}
			numbers := strings.Split(line, " ")
			dest, _ := strconv.Atoi(numbers[0])
			src, _ := strconv.Atoi(numbers[1])
			size, _ := strconv.Atoi(numbers[2])
			for j, seed := range seedsMap[0] {
				if seed >= src && seed < src + size {
					seedsMap[1][j] = dest + seed - src
					continue
				}
			}
		}
	}
	res = seedsMap[1][0]
	for _, dist := range seedsMap[1] {
		if dist < res {
			res = dist
		}
	}
	return strconv.Itoa(res)
}

func Resolve2(lines []string) (string) {
	res := 0
	seedsMap := [][]int{[]int{},[]int{}}
	for i, line := range lines {
		if i == 0 {
			numbers := strings.Split(strings.Split(line, ": ")[1], " ")
			for _, nAsStr := range numbers {
				n, _ := strconv.Atoi(nAsStr)
				seedsMap[0] = append(seedsMap[0], n)
				seedsMap[1] = append(seedsMap[1], n)
			}
		} else if len(line) > 0 {
			if strings.HasSuffix(line, ":") {
				seedsMap[0] = nil
				for _, seed := range seedsMap[1] {
					seedsMap[0] = append(seedsMap[0], seed)
				}
				continue
			}
			numbers := strings.Split(line, " ")
			dest, _ := strconv.Atoi(numbers[0])
			src, _ := strconv.Atoi(numbers[1])
			size, _ := strconv.Atoi(numbers[2])
			j := 0
			for j < len(seedsMap[0]) {
				if seedsMap[0][j] >= 0 {
					seed := seedsMap[0][j]
					seedSize := seedsMap[0][j+1]
					if seed < src && seed + seedSize > src{
						seedsMap[0][j+1] = src - seed
						seedsMap[1][j+1] = src - seed
						seedsMap[0] = append(seedsMap[0], src)
						seedsMap[0] = append(seedsMap[0], seedSize - seedsMap[1][j+1])
						seedsMap[1] = append(seedsMap[1], src)
						seedsMap[1] = append(seedsMap[1], seedSize - seedsMap[1][j+1])
					} else if seed >= src && seed < src + size {
						seedsMap[0][j] = -1 //done
						seedsMap[1][j] = dest + seed - src
						if seed + seedSize > src + size {
							seedsMap[1][j+1] = src + size - seed
							seedsMap[0] = append(seedsMap[0], seed + seedsMap[1][j+1])
							seedsMap[0] = append(seedsMap[0], seedSize - seedsMap[1][j+1])
							seedsMap[1] = append(seedsMap[1], seed + seedsMap[1][j+1])
							seedsMap[1] = append(seedsMap[1], seedSize - seedsMap[1][j+1])
						}
					}
				}
				j += 2
			}
		}
	}
	res = seedsMap[1][0]
	for j, dist := range seedsMap[1] {
		if j%2 == 0 && dist < res {
			res = dist
		}
	}
	return strconv.Itoa(res)
}

