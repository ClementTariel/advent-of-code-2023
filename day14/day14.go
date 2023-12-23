package main

import (
	"fmt"
	"strconv"
	"bytes"
)

func Resolve1(lines []string) (string) {
	res := 0

	columnCount := []int{}
	starts := []int{}
	for i := 0; i < len(lines[0]); i++ {
		columnCount = append(columnCount, 0)
		starts = append(starts, 0)
	}
	for i, line:= range lines {
		for j, c := range line {
			if string(c) == "O" {
				res += len(lines) - starts[j] - columnCount[j]
				columnCount[j] += 1
			}
			if string(c) == "#" {
				starts[j] = i+1
				columnCount[j] = 0
			}
		}
	}
	return strconv.Itoa(res)
}

func ExtractMap(lines []string) ([][]int, [][]int) {
	xMap := [][]int{}
	yMap := [][]int{}
	for i := 0; i < len(lines); i++ {
		xMap = append(xMap, []int{})
	}
	for j := 0; j < len(lines[0]); j++ {
		yMap = append(yMap, []int{})
	}
	for i, line := range lines {
		for j, c := range line {
			if string(c) == "#" {
				xMap[i] = append(xMap[i], j)
				yMap[j] = append(yMap[j], i)
			}
		}
	}
	for i, _ := range xMap {
		xMap[i] = append(xMap[i], len(lines[i]))
	}
	for j, _ := range yMap {
		yMap[j] = append(yMap[j], len(lines))
	}
	return xMap, yMap
}

func PrintState(xMap [][]int, yMap [][]int, state [][]int, xy bool, reverse bool) {
	piMap := &xMap
	if xy {
		piMap = &yMap
	}
	fmt.Println(state)
	for i, line := range state {
		for j, col := range line {
			if reverse {
				start := -1
				if j > 0 {
					start = (*piMap)[i][j-1]
				}
				for k := start + col + 1; k < (*piMap)[i][j]; k++ {
					fmt.Printf(".")
				}
				for k := 0; k < col; k++ {
					fmt.Printf("O")
				}
			} else {
				for k := 0; k < col; k++ {
					fmt.Printf("O")
				}
				start := -1
				if j > 0 {
					start = (*piMap)[i][j-1]
				}
				for k := start + col + 1; k < (*piMap)[i][j]; k++ {
					fmt.Printf(".")
				}
			}
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}

func NorthLoadOfEastState(xMap [][]int, state [][]int) (int) {
	load := 0
	for i, line := range state {
		for _, stackSize := range line {
			load += (len(state)-i)*stackSize
		}
	}
	return load
}

func NorthFromEast(xMap [][]int, yMap [][]int, state [][]int) ([][]int) {
	newState := [][]int{}
	for j, _ := range yMap {
		newState = append(newState, []int{})
		for i := 0; i < len(yMap[j]); i++ {
			newState[j] = append(newState[j], 0)
		}
	}
	for i, line := range state {
		for j, _ := range line {
			stackSize := line[len(line)-1-j]
			start := xMap[i][len(xMap[i])-1-j] - 1
			for k := 0; k < stackSize; k++ {
				col := start - k
				newLine := 0
				for l, northLine := range yMap[col] {
					if northLine < i {
						newLine = l+1
					} else {
						break
					}
				}
				newState[col][newLine] += 1
			}
		}
	}
	return newState
}

func WestFromNorth(xMap [][]int, yMap [][]int, state [][]int) ([][]int) {
	newState := [][]int{}
	for i, _ := range xMap {
		newState = append(newState, []int{})
		for j := 0; j < len(xMap[i]); j++ {
			newState[i] = append(newState[i], 0)
		}
	}
	for j, col := range state {
		for i, stackSize := range col {
			start := 0
			if i > 0 {
				start = yMap[j][i-1] + 1
			}
			for k := 0; k < stackSize; k++ {
				line := start + k
				newCol := 0
				for l := 0; l < len(xMap[line]); l++ {
					westCol := xMap[line][l]
					if westCol < j {
						newCol = l+1
					} else {
						break
					}
				}
				newState[line][newCol] += 1
			}
		}
	}
	return newState
}

func SouthFromWest(xMap [][]int, yMap [][]int, state [][]int) ([][]int) {
	newState := [][]int{}
	for j, _ := range yMap {
		newState = append(newState, []int{})
		for i := 0; i < len(yMap[j]); i++ {
			newState[
			j] = append(newState[j], 0)
		}
	}
	for i, line := range state {
		for j, stackSize := range line {
			start := 0
			if j > 0 {
				start = xMap[i][j-1] + 1
			}
			for k := 0; k < stackSize; k++ {
				col := start + k
				newLine := 0
				for l, southLine := range yMap[col] {
					if southLine < i {
						newLine = l+1
					} else {
						break
					}
				}
				newState[col][newLine] += 1
			}
		}
	}
	return newState
}

func EastFromSouth(xMap [][]int, yMap [][]int, state [][]int) ([][]int) {
	newState := [][]int{}
	for i, _ := range xMap {
		newState = append(newState, []int{})
		for j := 0; j < len(xMap[i]); j++ {
			newState[i] = append(newState[i], 0)
		}
	}
	for j, col := range state {
		for i, _ := range col {
			stackSize := col[len(col)-1-i]
			start := yMap[j][len(yMap[j])-1-i] - 1
			for k := 0; k < stackSize; k++ {
				line := start - k
				newCol := 0
				for l := 0; l < len(xMap[line]); l++ {
					westCol := xMap[line][l]
					if westCol < j {
						newCol = l+1
					} else {
						break
					}
				}
				newState[line][newCol] += 1
			}
		}
	}
	return newState
}

func CycleFromEast(xMap [][]int, yMap [][]int, state [][]int) ([][]int) {
	return EastFromSouth(xMap, yMap, SouthFromWest(xMap, yMap, WestFromNorth(xMap, yMap, NorthFromEast(xMap, yMap, state))))
}

func StateToString(state[][]int) (string){
	var buffer bytes.Buffer
	for _, line := range state {
		for _, n := range line {
			buffer.WriteString(strconv.Itoa(n))
		}
	}
	return buffer.String()
}

func Resolve2(lines []string) (string) {
	res := 0

	xMap, yMap := ExtractMap(lines)

	state := [][]int{}
	stateSeen := map[string]int{}

	columnCount := []int{}
	starts := []int{}
	for j := 0; j < len(lines[0]); j++ {
		columnCount = append(columnCount, 0)
		starts = append(starts, 0)
		state = append(state, []int{0})
	}
	for i, line:= range lines {
		for j, c := range line {
			if string(c) == "O" {
				columnCount[j] += 1
				state[j][len(state[j])-1] += 1
			}
			if string(c) == "#" {
				starts[j] = i+1
				columnCount[j] = 0
				state[j] = append(state[j], 0)
			}
		}
	}
	newState := EastFromSouth(xMap, yMap, SouthFromWest(xMap, yMap, WestFromNorth(xMap, yMap, state)))
	newStateAsString := StateToString(newState)
	_, seen := stateSeen[newStateAsString]
	count := 1
	for !seen {
		count += 1
		stateSeen[newStateAsString] = count
		newState = CycleFromEast(xMap, yMap, newState)
		newStateAsString = StateToString(newState)
		_, seen = stateSeen[newStateAsString]

	}
	lastSeen := stateSeen[newStateAsString]
	cycleLen := count + 1 - lastSeen
	totalLen := 1000000000
	delta := (totalLen - count)%cycleLen
	for k := 0; k < delta; k++ {
		newState = CycleFromEast(xMap, yMap, newState)
	}
	res = NorthLoadOfEastState(xMap, newState)
	return strconv.Itoa(res)
}

