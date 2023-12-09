package main

import (
	"strconv"
	"strings"
)

func Resolve1(lines []string) (string) {
	res := 0
	path := map[string]string{}
	moves := ""
	for i, line := range lines {
		if i == 0 {
			moves = line
		} else if len(line) > 0 {
			data := strings.Split(line, " = (")
			key := data[0]
			dests := strings.Split(data[1], ", ")
			destL := dests[0]
			destR := strings.Split(dests[1], ")")[0]
			path[key + "L"] = destL
			path[key + "R"] = destR
		}
	}
	loc := "AAA"
	for loc != "ZZZ" {
		loc = path[loc + string(moves[res%len(moves)])]
		res += 1
	}
	return strconv.Itoa(res)
}

func abs(n int) (int) {
	if n < 0 {
		return -n
	}
	return n
}

func Euclide(a int, b int) (int, int, int) {
	u1 := 1
	u2 := 0
	v1 := 0
	v2 := 1
	r1 := a
	r2 := b
	for r2 != 0 {
		q := r1/r2
		r := r1
		u := u1
		v := v1
		r1 = r2
		u1 = u2
		v1 = v2
		r2 = r - q*r2
		u2 = u - q*u2
		v2 = v - q*v2
	}
	return r1, u1, v1
}

func Resolve2(lines []string) (string) {
	res := 0
	path := map[string]string{}
	moves := ""
	locs := []string{}
	offsets := []int{}
	cycles := []int{}
	ends := [][]int{}
	for i, line := range lines {
		if i == 0 {
			moves = line
		} else if len(line) > 0 {
			data := strings.Split(line, " = (")
			key := data[0]
			dests := strings.Split(data[1], ", ")
			destL := dests[0]
			destR := strings.Split(dests[1], ")")[0]
			path[key + "L"] = destL
			path[key + "R"] = destR
			if string(key[len(key)-1]) == "A" {
				locs = append(locs, key)
				offsets = append(offsets, -1)
				cycles = append(cycles, -1)
				ends = append(ends, []int{})
			}
		}
	}
	tMin := 0
	for i, start := range locs {
		loc := start
		seen := map[string]bool{}
		t := 0
		key := loc+"-"+strconv.Itoa(t%len(moves))
		for !seen[key] {
			if string(loc[len(loc)-1]) == "Z" {
				ends[i] = append(ends[i], t)
			}
			loc = path[loc + string(moves[t%len(moves)])]
			seen[key] = true
			t += 1
			key = loc+"-"+strconv.Itoa(t%len(moves))
		}
		seen[key] = false
		loc = start
		offset := 0
		key = loc+"-"+strconv.Itoa(offset%len(moves))
		for seen[key] {
			loc = path[loc + string(moves[offset%len(moves)])]
			offset += 1
			key = loc+"-"+strconv.Itoa(offset%len(moves))
		}
		offsets[i] = offset
		cycles[i] = t - offset
		if ends[i][0] > tMin {
			tMin = ends[i][0]
		}
	}
	rs := []int{}
	ns := []int{}
	for i, end := range ends {
		ris := []int{}
		nis := []int{}
		ni := cycles[i]
		for _, rij := range end {
			ris = append(ris, rij%ni)
			nis = append(nis, ni)
		}
		nextRs := []int{}
		nextNs := []int{}
		for x, _ := range rs {
			for y, _ := range ris {
				c, u, v := Euclide(ns[x], nis[y])
				nextN := ns[x]*nis[y]/c
				nextR := (rs[x]*v*(nis[y]/c) + ris[y]*u*(ns[x]/c))*c
				for nextR < 0 {
					nextR += nextN
				}
				nextR = nextR%nextN
				nextNs = append(nextNs, nextN)
				nextRs = append(nextRs, nextR)
			}
		}
		rs = nil
		ns = nil
		if i == 0 {
			ns = nis
			rs = ris
		} else {
			rs = nextRs
			ns = nextNs
		}
	}
	for i, _ := range ns {
		if i == 0 {
			res = rs[0]
			for res < tMin {
				res += ns[0]
			}
		} else {
			localRes := rs[i]
			for localRes < tMin {
				localRes += ns[i]
			}
			if localRes < res {
				res = localRes
			}
		}
	}
	return strconv.Itoa(res)
}


