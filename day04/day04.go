package main

import (
	"strconv"
	"slices"
	"strings"
)

func RemoveEmpty(s []string) ( []string ) {
	ret := []string{}
	for _, n := range s {
		if n != "" {
			ret = append(ret, n)
		}
	}
	return ret
}

func Resolve1(lines []string) (string) {
	res := 0
	for _, line := range lines {
		data := strings.Split(line, ": ")
		allNumbers := strings.Split(data[1], "|")
		winningNumbers := RemoveEmpty(strings.Split(allNumbers[0], " "))
		elfNumbers := RemoveEmpty(strings.Split(allNumbers[1], " "))
		winningCount := 0
		for _, n := range elfNumbers {
			if slices.Contains(winningNumbers, n) {
				winningCount += 1
			}
		}
		if winningCount > 0 {
			res += 1 << (winningCount - 1)
		}
	}
	return strconv.Itoa(res)
}

func Resolve2(lines []string) (string) {
	res := 0
	var cardCount = make([]int,len(lines))
	for k, line := range lines {
		cardCount[k] += 1
		data := strings.Split(line, ": ")
		allNumbers := strings.Split(data[1], "|")
		winningNumbers := RemoveEmpty(strings.Split(allNumbers[0], " "))
		elfNumbers := RemoveEmpty(strings.Split(allNumbers[1], " "))
		winningCount := 0
		for _, n := range elfNumbers {
			if slices.Contains(winningNumbers, n) {
				winningCount += 1
			}
		}
		for i := k+1; i <= k+winningCount && i<len(lines); i++ {
			cardCount[i] += cardCount[k]
		}
		res += cardCount[k]
	}
	return strconv.Itoa(res)
}


