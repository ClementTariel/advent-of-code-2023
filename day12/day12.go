package main

import (
	"strconv"
	"strings"
)

func Choose(n int, k int) (int) {
	res := 1
	if n-k < k {
		k = n-k
	}
	if k < 1 || n < 1 {
		return 1
	}
	for i := 0; i < k; i++ {
		res *= n-i
	}
	for i := 0; i < k; i++ {
		res /= (i+1)
	}
	return res
}


func GreedyCount(
	springs []string,
	counts []int,
	upcomingDamagedCount int,
	upcomingUnknownCount int,
	totalCount int,
	currentCount int)(int) {
	if upcomingDamagedCount + currentCount > totalCount {
		return 0
	}
	if upcomingDamagedCount + upcomingUnknownCount + currentCount < totalCount {
		return 0
	}
	if len(springs) == 0 {
		if upcomingDamagedCount == 0 && (len(counts) == 0 || (len(counts) == 1 && counts[0] == currentCount)) {
			return 1
		}
		return 0
	}
	if len(counts) > 0  && currentCount > counts[0] {
		return 0
	}
	i := 0
	groupSize := 0
	spring := string(springs[0])
	for i < len(springs) && string(springs[i]) == spring {
		groupSize += 1
		i += 1
	}
	if spring == "#" {
		upcomingDamagedCount -= groupSize
		if groupSize + currentCount == counts[0] {
			totalCount -= groupSize + currentCount
			offset := 0
			if i < len(springs) && springs[i] == "?" {
				upcomingUnknownCount -= 1
			}
			if i + offset < len(springs) {
				offset += 1
			}
			return GreedyCount(springs[i+offset:],
				counts[1:],
				upcomingDamagedCount,
				upcomingUnknownCount,
				totalCount,
				0)
		} else {
			if groupSize + currentCount > counts[0] {
				return 0
			}
			return GreedyCount(
				springs[i:],
				counts,
				upcomingDamagedCount,
				upcomingUnknownCount,
				totalCount,
				currentCount + groupSize)
		}
	} else if spring == "." {
		if currentCount > 0 && counts[0] == currentCount {
			totalCount -= currentCount
			return GreedyCount(springs[i:], counts[1:], upcomingDamagedCount, upcomingUnknownCount, totalCount, 0)
		} else if currentCount > 0 && counts[0] != currentCount {
			return 0
		}
		return GreedyCount(springs[i:], counts, upcomingDamagedCount, upcomingUnknownCount, totalCount, 0)
	} else {
		if currentCount > 0 {
			completion := counts[0] - currentCount
			if completion < groupSize {
				i += completion - groupSize + 1
				upcomingUnknownCount -= completion + 1
				totalCount -= counts[0]
				return GreedyCount(springs[i:], counts[1:], upcomingDamagedCount, upcomingUnknownCount, totalCount, 0)
			} else {
				return GreedyCount(
					springs[i:],
					counts,
					upcomingDamagedCount,
					upcomingUnknownCount - groupSize,
					totalCount,
					currentCount + groupSize)
			}
		}
		upcomingUnknownCount -= groupSize
		nextGroupsTotalSize := 0
		nextIsEmpty := i == len(springs) || springs[i] == "."
		offset := 0
		nextPartialDamaged := 0
		if !nextIsEmpty {
			offset = 1
			j := i
			for j < len(springs) && springs[j] == "#" {
				j += 1
				nextPartialDamaged += 1
			}
		}
		possibilities := 0
		if nextIsEmpty {
			possibilities = GreedyCount(
				springs[i:],
				counts,
				upcomingDamagedCount,
				upcomingUnknownCount,
				totalCount,
				0)
		} else {

			for j:= 0; j <= groupSize; j++ {
				possibilities += GreedyCount(
					springs[i+nextPartialDamaged:],
					counts,
					upcomingDamagedCount - nextPartialDamaged,
					upcomingUnknownCount,
					totalCount,
					j + nextPartialDamaged)
			}
		}
		for k, nextGroupSize := range counts {
			nextGroupsTotalSize += nextGroupSize
			if nextGroupsTotalSize + k + offset > groupSize {
				break
			}
			totalCount -= nextGroupSize
			if nextIsEmpty {
				newPossibilities := GreedyCount(
					springs[i:],
					counts[k+1:],
					upcomingDamagedCount,
					upcomingUnknownCount,
					totalCount,
					0)
				n := groupSize + 1 - nextGroupsTotalSize
				possibilities += Choose(n, k+1) * newPossibilities
			} else {
				for j := 0; j < groupSize - nextGroupsTotalSize - k; j++ {
					newPossibilities := GreedyCount(
						springs[i+nextPartialDamaged:],
						counts[k+1:],
						upcomingDamagedCount - nextPartialDamaged,
						upcomingUnknownCount,
						totalCount,
						j + nextPartialDamaged)
					n := groupSize - nextGroupsTotalSize - j
					possibilities += Choose(n, k+1) * newPossibilities
				}
			}
		}
		return possibilities
	}
}

func Resolve1(lines []string) (string) {
	res := 0
	for _, line := range lines {
		data := strings.Split(line, " ")
		springsAsStr := data[0]
		upcomingDamagedCount := 0
		upcomingUnknownCount := 0
		springs := []string{}
		for _, c := range springsAsStr {
			l := string(c)
			springs = append(springs, l)
			if l == "#" {
				upcomingDamagedCount += 1
			} else if l == "?" {
				upcomingUnknownCount += 1
			}
		}
		totalCount := 0
		damagedCounts := []int{}
		for _, nAsStr := range strings.Split(data[1], ",") {
			n, _ :=  strconv.Atoi(nAsStr)
			damagedCounts = append(damagedCounts, n)
			totalCount += n
		}
		possibilities := GreedyCount(springs, damagedCounts, upcomingDamagedCount, upcomingUnknownCount, totalCount, 0)
		res += possibilities
	}
	return strconv.Itoa(res)
}
