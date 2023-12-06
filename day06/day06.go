package main

import (
	"strconv"
	"strings"
	"math"
)

func RemoveEmpty(s []string) ([]string ) {
	ret := []string{}
	for _, n := range s {
		if n != "" {
			ret = append(ret, n)
		}
	}
	return ret
}

func Resolve1(lines []string) (string) {
	res := 1
	times := RemoveEmpty(strings.Split(strings.Split(lines[0], ":")[1], " "))
	dists := RemoveEmpty(strings.Split(strings.Split(lines[1], ":")[1], " "))
	if len(times) != len(dists) {
		panic("times and dists have different lenght")
	}
	for i, _ := range times {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(dists[i])
		// a := 1
		// b := -t
		// c := d
		t1float := (float64(t) - math.Sqrt(float64(t*t - 4*d)))/2
		t1 := int(t1float)
		if t1 < 0 {
			t1 = 0
		}
		res *= t - 1 - 2*t1
	}
	return strconv.Itoa(res)
}

func Resolve2(lines []string) (string) {
	time := strings.Join(strings.Split(strings.Split(lines[0], ":")[1], " "), "")
	dist := strings.Join(strings.Split(strings.Split(lines[1], ":")[1], " "), "")
	// a := 1
	// b := -t
	// c := d
	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(dist)
	t1float := (float64(t) - math.Sqrt(float64(t*t - 4*d)))/2
	t1 := int(t1float)
	if t1 < 0 {
		t1 = 0
	}
	res := t - 1 - 2*t1
	return strconv.Itoa(res)
}

