package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"os/exec"
	"errors"
	"bufio"
	"plugin"
)

func getDays() ([]string, error) {
	dir, err := os.Open("./")
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	
	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var days []string
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() && strings.HasPrefix(fileName, "day") {
			days = append(days, fileName)
		}
	}
	return days, nil
}

func getLastDay() (int, string, error){
	days, err := getDays()
	if err != nil {
		return 0, "", err
	}
	max := 0
	dayName := ""
	for _, day := range days {
		if len(day) > 3 && day[0:3] == "day" {	
			num, _ := strconv.Atoi(day[3:])
			if num > max {
				max = num
				dayName = day
			}
		}
	}
	return max, dayName, nil
}

func main() {
	fmt.Println("Start")


	lastDay, dayName, _ := getLastDay()
	fmt.Printf("day %d\n",lastDay)
	fmt.Printf("file name : %s\n",dayName)
	
	sessionToken := ""
	sessionFile, err := os.Open("session.txt")
	if err != nil {
		panic(err)
	}
	defer sessionFile.Close()
	sessionScanner := bufio.NewScanner(sessionFile)
	for sessionScanner.Scan() {
		line := sessionScanner.Text()
		if len(line) > 0 {
			sessionToken = line
		}
	}
	
	if len(sessionToken) > 0 {
		if _, err := os.Stat(dayName + "/input.txt"); errors.Is(err, os.ErrNotExist) {
		
			cmd := exec.Command("curl","https://adventofcode.com/2023/day/"+strconv.Itoa(lastDay)+"/input", "--cookie", "session="+sessionToken, "-o", dayName+"/input.txt")
			fmt.Printf("%s\n", cmd.String())
			stdout, err := cmd.Output()

			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("%s\n", stdout)
		
		} else {
			fmt.Println("Input already downloaded")
		}
	} else {
		fmt.Println("$AoCSession env var was not defined")
	}

	var testLines []string
	test, err := os.Open(dayName + "/example.txt")
	if err != nil {
		fmt.Println("example input not found (Could not find example.txt)")
	} else {
		defer test.Close()
		scanner := bufio.NewScanner(test)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 0 {
				testLines = append(testLines, line)
			}
		}
		fmt.Printf("%d lines in example\n", len(testLines))
	}

	file, err := os.Open(dayName + "/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	fmt.Printf("%d lines in input\n", len(lines))
	
	p, err := plugin.Open(dayName+"/"+dayName+".so")
	if err != nil {
		panic(err)
	}
	
	f1Exists := true
	f1, err := p.Lookup("Resolve1")
	if err != nil {
		f1Exists = false
	}

	f2Exists := true
	f2, err := p.Lookup("Resolve2")
	if err != nil {
		f2Exists = false
	}

	fmt.Println("\n==========\n")
	if f1Exists {
		test1 := ""
		if len(testLines) > 0 {
			test1 = f1.(func(lines []string)(string))(testLines)
			fmt.Printf("example 1 : '%s'\n", test1)
		} else {
			fmt.Println("not tested with example")
		}
		result1 := f1.(func(lines []string)(string))(lines)
		fmt.Printf("result 1 : '%s'\n", result1)
	} else {
		fmt.Println("Resolve1 not found")
	}
	fmt.Println("\n==========\n")
	if f2Exists {
		test2 := ""
		if len(testLines) > 0 {
			test2 = f2.(func(lines []string)(string))(testLines)
			fmt.Printf("example 2 : '%s'\n", test2)
		} else {
			fmt.Println("not tested with example")
		}
		result2 := f2.(func(lines []string)(string))(lines)
		fmt.Printf("result 2 : '%s'\n", result2)
	} else {
		fmt.Println("Resolve2 not found")
	}
	fmt.Println("\n==========\n")
	fmt.Println("End")
}

