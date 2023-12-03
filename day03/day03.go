package main

import (
	"strconv"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func IsNumber(s string) (bool) {
	return s=="0" || s=="1" || s=="2" || s=="3" || s=="4" || s=="5" || s=="6" || s=="7" || s=="8" || s=="9"
}

func IsEnginePart(letter string) (bool) {
	return !IsNumber(letter) && letter != "."
}

func Resolve1(lines []string) (string) {
	res := 0
	for i, line := range lines {
		numBuffer := ""
		for j, char := range line {
			letter := string(char)
			if IsNumber(letter){
				numBuffer += letter
			}
			if (!IsNumber(letter) || j == len(line)-1) && len(numBuffer) > 0 {
				isPartNumber := false
				partNumber, _ := strconv.Atoi(numBuffer)
				offset := 0
				if !IsNumber(letter){
					offset = 1
				}
				if i > 0 {
					for k := Max(0,j-offset-len(numBuffer)); k <= j; k++ {
						if IsEnginePart(string(lines[i-1][k])) {
							isPartNumber = true
							break
						}
					}
				}
				if IsEnginePart(string(lines[i][Max(0,j-offset-len(numBuffer))])) ||
					IsEnginePart(string(lines[i][j])) {
					isPartNumber = true
				}
				if !isPartNumber && i < len(lines)-1 {
					for k := Max(0,j-offset-len(numBuffer)); k <= j; k++ {
						if IsEnginePart(string(lines[i+1][k])) {
							isPartNumber = true
							break
						}
					}
				}
				if isPartNumber {
					res += partNumber
				}
				numBuffer = ""
			}
		}
	}
	return strconv.Itoa(res)
}

func Resolve2(lines []string) (string) {
	res := 0
	for i, line := range lines {
		for j, char := range line {
			letter := string(char)
			if letter == "*" {
				numBuffer := ""
				partCount := 0
				power := 1
				x := i
				y := j
				if i > 0 {
					x = i-1
					for y > 0 && IsNumber(string(lines[x][y-1])){
						y-=1
					}
					numBuffer = ""
					for IsNumber(string(lines[x][y])) && y < len(lines[x])-1 {
						numBuffer += string(lines[x][y])
						y += 1
					}
					if IsNumber(string(lines[x][y])) {
						numBuffer += string(lines[x][y])
					}
					if len(numBuffer) > 0 {
						partCount += 1
						num, _ := strconv.Atoi(numBuffer)
						power *= num
					}
					y = j
					numBuffer = ""
					if !IsNumber(string(lines[x][y])) {
						for y+1<len(line) && IsNumber(string(lines[x][y+1])) {
							numBuffer += string(lines[x][y+1])
							y += 1
						}
						if len(numBuffer) > 0 {
							partCount += 1
							num, _ := strconv.Atoi(numBuffer)
							power *= num
						}
					}
				}
				y = j
				if i < len(lines)-1 {
					x = i+1
					for y > 0 && IsNumber(string(lines[x][y-1])){
						y-=1
					}
					numBuffer = ""
					for IsNumber(string(lines[x][y])) && y < len(lines[x])-1 {
						numBuffer += string(lines[x][y])
						y += 1
					}
					if IsNumber(string(lines[x][y])) {
						numBuffer += string(lines[x][y])
					}
					if len(numBuffer) > 0 {
						partCount += 1
						num, _ := strconv.Atoi(numBuffer)
						power *= num
					}
					y = j
					numBuffer = ""
					if !IsNumber(string(lines[x][y])) {
						for y+1<len(line) && IsNumber(string(lines[x][y+1])) {
							numBuffer += string(lines[x][y+1])
							y += 1
						}
						if len(numBuffer) > 0 {
							partCount += 1
							num, _ := strconv.Atoi(numBuffer)
							power *= num
						}
					}
				}
				x = i
				y = j
				for y > 0 && IsNumber(string(lines[x][y-1])){
					y-=1
				}
				numBuffer = ""
				for IsNumber(string(lines[x][y])) && y < len(lines[x])-1 {
					numBuffer += string(lines[x][y])
					y += 1
				}
				if IsNumber(string(lines[x][y])) {
					numBuffer += string(lines[x][y])
				}
				if len(numBuffer) > 0 {
					partCount += 1
					num, _ := strconv.Atoi(numBuffer)
					power *= num
				}
				y = j
				numBuffer = ""
				for y+1 < len(lines[x]) && IsNumber(string(lines[x][y+1])) {
					numBuffer += string(lines[x][y+1])
					y += 1
				}
				if len(numBuffer) > 0 {
					partCount += 1
					num, _ := strconv.Atoi(numBuffer)
					power *= num
				}
				if partCount == 2 {
					res += power
				}
			}
		}
	}
	return strconv.Itoa(res)
}


