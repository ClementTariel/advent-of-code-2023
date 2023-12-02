package main

import (
	"regexp"
	"strconv"
)

func Resolve1(lines []string) (string) {
	res := 0
	for _, line := range lines {
		r2 := regexp.MustCompile(`^.*([0-9])[^0-9]*$`)
		r1 := regexp.MustCompile(`^[^0-9]*([0-9]).*$`)
		matches1 := r1.FindStringSubmatch(line)
		matches2 := r2.FindStringSubmatch(line)
		d1, err := strconv.Atoi(matches1[1])
		if err != nil {
			panic(err)
		}
		d2, err := strconv.Atoi(matches2[1])
		if err != nil {
			panic(err)
		}
		res += 10*d1 + d2
	}
	return strconv.Itoa(res)
}

func UnefficientStringReverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}


func Resolve2(lines []string) (string) {
	letterToInt := map[string]int {	
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
	res := 0
	for _, line := range lines {
		numbers := "one|two|three|four|five|six|seven|eight|nine"
		re1 := regexp.MustCompile(`([1-9]|`+numbers+`)`)
		matches1 := re1.FindAllStringSubmatch(line,-1)
		num1 := matches1[0][1]

		reversedNumbers := UnefficientStringReverse(numbers)
		re2 := regexp.MustCompile(`([1-9]|`+reversedNumbers+`)`)
		matches2 := re2.FindAllStringSubmatch(UnefficientStringReverse(line),-1)
		num2 := UnefficientStringReverse(matches2[0][1])

		d1, err := strconv.Atoi(num1)
		if err != nil {
			d1 = letterToInt[num1]
		}
		d2, err := strconv.Atoi(num2)
		if err != nil {
			d2 = letterToInt[num2]
		}
		res += 10*d1 + d2
	}
	return strconv.Itoa(res)
}
