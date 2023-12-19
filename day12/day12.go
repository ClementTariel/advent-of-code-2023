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
	if k < 0 {
		return 0
	}
	for i := 1; i <= k; i++ {
		res *= n-(k-i)
		res /= i
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
		if len(counts) == 0 || (len(counts) == 1 && counts[0] == currentCount) {
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
		return GreedyCount(springs[i:],
			counts,
			upcomingDamagedCount - groupSize,
			upcomingUnknownCount,
			totalCount,
			currentCount + groupSize)
	} else if spring == "." {
		if currentCount > 0 && counts[0] != currentCount {
			return 0
		} else if currentCount > 0 {
			return GreedyCount(springs[i:],
				counts[1:],
				upcomingDamagedCount,
				upcomingUnknownCount,
				totalCount - currentCount,
				0)
		}
		return GreedyCount(springs[i:], counts, upcomingDamagedCount, upcomingUnknownCount, totalCount, 0)
	} else {
		if currentCount > 0 {
			completion := counts[0] - currentCount
			if completion < groupSize {
				i -= groupSize
				i += completion + 1
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
		nextIsDamaged := i < len(springs) && springs[i] == "#"
		nextIsEmpty := i == len(springs) || springs[i] == "."
		emptyOffset := 0
		if i < len(springs) {
			emptyOffset = 1
		}
		nextPartialDamaged := 0
		if nextIsDamaged {
			j := i
			for j < len(springs) && springs[j] == "#" {
				j += 1
				nextPartialDamaged += 1
				upcomingDamagedCount -= 1
			}
		}
		possibilities := 0
		if nextIsEmpty {
			possibilities += GreedyCount(
				springs[i+emptyOffset:],
				counts,
				upcomingDamagedCount,
				upcomingUnknownCount,
				totalCount,
				0)
		}
		if nextIsDamaged {
			for j:= 0; j <= groupSize; j++ {
				possibilities += GreedyCount(
					springs[i+nextPartialDamaged:],
					counts,
					upcomingDamagedCount,
					upcomingUnknownCount,
					totalCount,
					j + nextPartialDamaged)
			}
		}
		nextGroupsTotalSize := 0
		offset := 0
		if nextIsDamaged {
			offset = 1
		}
		for k, nextGroupSize := range counts {
			nextGroupsTotalSize += nextGroupSize
			if nextGroupsTotalSize + k > groupSize - offset {
				break
			}
			n := groupSize + 1 - nextGroupsTotalSize - offset
			if nextIsEmpty {
				newPossibilities := GreedyCount(
					springs[i+emptyOffset:],
					counts[k+1:],
					upcomingDamagedCount,
					upcomingUnknownCount,
					totalCount - nextGroupsTotalSize,
					0)
				possibilities += Choose(n, k+1) * newPossibilities
			}
			if nextIsDamaged {
				for j := 0; j < n; j++ {
					newPossibilities := GreedyCount(
						springs[i+nextPartialDamaged:],
						counts[k+1:],
						upcomingDamagedCount,
						upcomingUnknownCount,
						totalCount - nextGroupsTotalSize,
						j + nextPartialDamaged)
					possibilities += Choose(n-j, k+1) * newPossibilities
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

func Resolve2(lines []string) (string) {
	res := 0
	for _, line := range lines {
		data := strings.Split(line, " ")
		springsAsStr := data[0]
		upcomingDamagedCount := 0
		upcomingUnknownCount := 0
		springs := []string{}
		for k := 0; k < 5; k++ {
			if k > 0 {
				springs = append(springs, "?")
				upcomingUnknownCount += 1
			}
			for _, c := range springsAsStr {
				l := string(c)
				springs = append(springs, l)
				if l == "#" {
					upcomingDamagedCount += 1
				} else if l == "?" {
					upcomingUnknownCount += 1
				}
			}
		}
		totalCount := 0
		damagedCounts := []int{}
		for k := 0; k < 5; k++ {
			for _, nAsStr := range strings.Split(data[1], ",") {
				n, _ :=  strconv.Atoi(nAsStr)
				damagedCounts = append(damagedCounts, n)
				totalCount += n
			}
		}
		possibilities := GreedyCount(springs, damagedCounts, upcomingDamagedCount, upcomingUnknownCount, totalCount, 0)
		res += possibilities
	}
	return strconv.Itoa(res)
}
