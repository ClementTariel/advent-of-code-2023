package main

import (
	"strconv"
	"strings"
)

func choose(n int, k int) (int) {
	res := 1
	if n-k < k {
		k = n-k
	}
	for i := 0; i < k; i++ {
		res *= n-i
	}
	for i := 0; i < k; i++ {
		res /= (i+1)
	}
	return res
}

func f(coeffs []int, t int) (int) {
	res := 0
	for k, coeff := range coeffs {
		if k <= t {
			res += coeff * choose(t, k)
		}
	}
	return res
}

func Resolve1(lines []string) (string) {
	res := 0
	for _, line := range lines {
		valuesAsStr := strings.Split(line, " ")
		coeffs := []int{}
		diffs := []int{}
		k := 0
		coeff, _ := strconv.Atoi(valuesAsStr[k])
		coeffs = append(coeffs, coeff)
		for _, vAsStr := range valuesAsStr {
			v, _ := strconv.Atoi(vAsStr)
			diffs = append(diffs, v)
		}
		allZeros := false
		for !allZeros && len(diffs) > 1 {
			allZeros = true
			for j := 1; j < len(diffs); j++ {
				diffs[j-1] = diffs[j] - diffs[j-1]
				if diffs[j-1] != 0 {
					allZeros = false
				}
			}
			diffs = diffs[:len(diffs)-1]
			coeffs = append(coeffs, diffs[0])
		}
		prediction := f(coeffs,len(valuesAsStr))
		res += prediction
	}
	return strconv.Itoa(res)
}

func Resolve2(lines []string) (string) {
	res := 0
	for _, line := range lines {
		valuesAsStr := strings.Split(line, " ")
		coeffs := []int{}
		diffs := []int{}
		k := 0
		coeff, _ := strconv.Atoi(valuesAsStr[k])
		coeffs = append(coeffs, coeff)
		for _, vAsStr := range valuesAsStr {
			v, _ := strconv.Atoi(vAsStr)
			diffs = append(diffs, v)
		}
		allZeros := false
		for !allZeros && len(diffs) > 1 {
			allZeros = true
			for j := 1; j < len(diffs); j++ {
				diffs[j-1] = diffs[j] - diffs[j-1]
				if diffs[j-1] != 0 {
					allZeros = false
				}
			}
			diffs = diffs[:len(diffs)-1]
			coeffs = append(coeffs, diffs[0])
		}
		prediction := 0
		for j := len(coeffs)-1; j > 0; j-- {
			prediction = coeffs[j-1] - prediction
		}
		res += prediction
	}
	return strconv.Itoa(res)
}

