package main

import (
	"strconv"
	"strings"
	"sort"
)


func Score(hand string) (int) {
	counts := [5]int{0,0,0,0,0}
	for i, s1 := range hand {
		count := 0
		for _, s2 := range hand {
			if s1 == s2 {
				count += 1
			}
		}
		counts[i] = count
	}
	value := 0
	if counts[0] == 5 {
		value = 6
	} else {
		maxI := 0
		for i, c := range counts {
			if c > counts[maxI] {
				maxI = i
			}
		}
		max2 := 0
		for i, c := range counts {
			if c > max2 && hand[i] != hand[maxI] {
				max2 = c
			}
		}
		if counts[maxI] == 4 {
			value = 5
		} else if counts[maxI] == 3 && max2 == 2 {
			value = 4
		} else if counts[maxI] == 3 {
			value = 3
		} else if counts[maxI] == 2 && max2 == 2 {
			value = 2
		} else if counts[maxI] == 2 {
			value = 1
		}
	}
	for i, _ := range hand {
		d := string(hand[i])
		if d == "T" {
			d = "10"
		} else if d == "J" {
			d = "11"
		} else if d == "Q" {
			d = "12"
		} else if d == "K" {
			d = "13"
		} else if d == "A" {
			d = "14"
		}
		value *= 15
		k, _ := strconv.Atoi(d)
		value += k
	}
	return value
}

func Resolve1(lines []string) (string) {
	res := 0

	values := []int{}
	scores := []int{}
	n := len(lines)
	for i, line := range lines {
		data := strings.Split(line, " ")
		hand := data[0]
		value, _ := strconv.Atoi(data[1])
		values = append(values, value)
		score := Score(hand)
		score *= n
		score += i
		scores = append(scores, score)
	}
	sort.Ints(scores)
	for i, score := range scores {
		res += values[score%n] * (i+1)
	}
	return strconv.Itoa(res)
}

func JokerScore(hand string) (int) {
	counts := [5]int{0,0,0,0,0}
	jokerCount := 0
	for i, s1 := range hand {
		if string(s1) == "J" {
			jokerCount += 1
		}
		count := 0
		for _, s2 := range hand {
			if s1 == s2 {
				count += 1
			}
		}
		counts[i] = count
	}
	value := 0
	if counts[0] == 5 {
		value = 6
	} else {
		maxI := 0
		for i, c := range counts {
			maxCount := counts[maxI]
			countI := c
			if string(hand[maxI]) == "J" {
				maxCount = 0
			}
			if string(hand[i]) == "J" {
				countI = 0
			}
			if countI > maxCount {
				maxI = i
			}
		}
		max2 := 0
		for i, c := range counts {
			countI := c
			if string(hand[i]) == "J" {
				countI = 0
			}
			if countI > max2 && hand[i] != hand[maxI] {
				max2 = c
			}
		}
		if counts[maxI] + jokerCount == 5 {
			value = 6
		} else if counts[maxI] + jokerCount == 4 {
			value = 5
		} else if counts[maxI] + jokerCount == 3 && max2 == 2 {
			value = 4
		} else if counts[maxI] + jokerCount == 3 {
			value = 3
		} else if counts[maxI] + jokerCount == 2 && max2 == 2 {
			value = 2
		} else if counts[maxI] + jokerCount == 2 {
			value = 1
		}
	}
	for i, _ := range hand {
		d := string(hand[i])
		if d == "T" {
			d = "10"
		} else if d == "J" {
			d = "1"
		} else if d == "Q" {
			d = "12"
		} else if d == "K" {
			d = "13"
		} else if d == "A" {
			d = "14"
		}
		value *= 15
		k, _ := strconv.Atoi(d)
		value += k
	}
	return value
}

func Resolve2(lines []string) (string) {
	res := 0

	values := []int{}
	scores := []int{}
	n := len(lines)
	for i, line := range lines {
		data := strings.Split(line, " ")
		hand := data[0]
		value, _ := strconv.Atoi(data[1])
		values = append(values, value)
		score := JokerScore(hand)
		score *= n
		score += i
		scores = append(scores, score)
	}
	sort.Ints(scores)
	for i, score := range scores {
		res += values[score%n] * (i+1)
	}
	return strconv.Itoa(res)
}
