package main

import (
	"strconv"
	"strings"
)

func Resolve1(lines []string) (string) {
	colorToCap := map[string]int {	
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	res := 0
	for _, line := range lines {
		data := strings.Split(line, ": ")
		gameIdAsStr := strings.Split(data[0], " ")[1]
		gameId, _ := strconv.Atoi(gameIdAsStr)
		sets := strings.Split(data[1], "; ")
		setIsValid := true
		for _, set := range sets {
			for _, draw := range strings.Split(set, ", ") {
				color := strings.Split(draw, " ")[1]
				numberAsStr := strings.Split(draw, " ")[0]
				number, _ := strconv.Atoi(numberAsStr)
				if number > colorToCap[color] {
					setIsValid = false
					break
				}
			}
			if !setIsValid {
				break
			}
		}
		if setIsValid {
			res += gameId
		}
	}
	return strconv.Itoa(res)
}


func Resolve2(lines []string) (string) {
	res := 0
	for _, line := range lines {
		data := strings.Split(line, ": ")
		sets := strings.Split(data[1], "; ")
		r := 0
		g := 0
		b := 0
		colorToCount := map[string]*int {
			"red": &r,
			"green": &g,
			"blue": &b,
		}
		for _, set := range sets {
			for _, draw := range strings.Split(set, ", ") {
				color := strings.Split(draw, " ")[1]
				numberAsStr := strings.Split(draw, " ")[0]
				number, _ := strconv.Atoi(numberAsStr)
				if number > *(colorToCount[color]) {
					*(colorToCount[color]) = number
				}
			}
		}
		power := 1
		power *= *(colorToCount["red"])
		power *= *(colorToCount["green"])
		power *= *(colorToCount["blue"])
		res += power
	}
	return strconv.Itoa(res)
}
