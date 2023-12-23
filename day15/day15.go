package main

import (
	"strconv"
	"strings"
)

func hash(s string) (int) {
	res := 0
	for _, char := range s {
		res += int(char)
		res *= 17
		res = res%256
	}
	return res
}

func Resolve1(lines []string) (string) {
	res := 0
	line := lines[0]
	for _, seq := range strings.Split(line, ",") {
		res += hash(seq)
	}
	return strconv.Itoa(res)
}

type focalLens struct {
    name string
    focal int
}

func Resolve2(lines []string) (string) {
	res := 0
	line := lines[0]
	boxes := map[int][]focalLens{}
	for _, seq := range strings.Split(line, ",") {
		seq = strings.Join(strings.Split(seq, "="),"/=")
		seq = strings.Join(strings.Split(seq, "-"),"/-")
		data := strings.Split(seq, "/")
		operation := data[1][0]
		boxName := hash(data[0])
		box, exists := boxes[boxName];
		if string(operation) == "-" {
			if !exists {
				boxes[boxName] = []focalLens{}
			} else {
				for i, lens := range box {
					if lens.name == data[0] {
						boxes[boxName] = append(box[:i], box[i+1:]...)
						break
					}
				}
			}
		} else {
			focal, _ := strconv.Atoi(string(data[1][1]))
			if !exists {
				boxes[boxName] = []focalLens{focalLens{name: data[0], focal: focal}}
			} else {
				i := len(box)
				for k, lens := range box {
					if lens.name == data[0] {
						i = k
						break
					}
				}
				if i == len(box) {
					boxes[boxName] = append(box, focalLens{name: data[0], focal: focal})
				} else {
					boxes[boxName][i].focal = focal
				}
			}
		}
	}
	for boxName, box := range boxes {
		for i, lens := range box {
			res += (boxName+1)*(i+1)*lens.focal
		}
	}
	return strconv.Itoa(res)
}
