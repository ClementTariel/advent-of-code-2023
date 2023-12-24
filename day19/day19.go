package main

import (
	"strings"
	"strconv"
)

type step struct {
	part string
	lower bool
	value int
	next string
}

type workflow struct {
	steps []step
	redirect string
}

func accept(p map[string]int, workflows map[string]workflow, name string) (bool) {
	if name == "A" {
		return true
	}
	if name == "R" {
		return false
	}
	w, _ := workflows[name]
	for _, step := range w.steps {
		if (step.lower && p[step.part] < step.value) || (!step.lower && p[step.part] > step.value) {
			return accept(p, workflows, step.next)
		}
	}
	return accept(p, workflows, w.redirect)
}

func Resolve1(lines []string) (string) {
	res := 0
	workflows := map[string](workflow){}
	workflowsAllRead := false
	for _, line := range lines {
		if len(line) == 0 {
			workflowsAllRead = true
			continue
		}
		if !workflowsAllRead {
			name := strings.Split(line, "{")[0]
			datas := strings.Split(strings.Split(strings.Split(line, "{")[1], "}")[0], ",")
			steps := []step{}
			for _, data := range datas[:len(datas)-1] {
				value, _ := strconv.Atoi(strings.Split(data, ":")[0][2:])
				next := strings.Split(data, ":")[1]
				steps = append(steps, step{
					part: string(data[0]),
					lower: string(data[1]) == "<",
					value : value,
					next: next})
			}
			workflows[name] = workflow{steps: steps, redirect: datas[len(datas)-1]}
		} else {
			data := strings.Split(strings.Split(strings.Split(line, "{")[1], "}")[0], ",")
			x, _ := strconv.Atoi(strings.Split(data[0], "=")[1])
			m, _ := strconv.Atoi(strings.Split(data[1], "=")[1])
			a, _ := strconv.Atoi(strings.Split(data[2], "=")[1])
			s, _ := strconv.Atoi(strings.Split(data[3], "=")[1])
			p := map[string]int{"x": x, "m": m, "a": a, "s": s}
			if accept(p, workflows, "in") {
				res += x + m + a + s
			}
		}
	}
	return strconv.Itoa(res)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func count(p map[string][][2]int, workflows map[string]workflow, name string) (int) {
	if name == "A" {
		n := 1
		for _, v := range p {
			size := 0
			for _, interval := range v {
				size += 1 + interval[1] - interval[0]
			}
			n *= size
		}
		return n
	}
	if name == "R" {
		return 0
	}
	w, _ := workflows[name]
	n := 0
	for _, step := range w.steps {
		newP := map[string][][2]int{}
		newP["x"] = p["x"]
		newP["m"] = p["m"]
		newP["a"] = p["a"]
		newP["s"] = p["s"]
		newIntervals := [][2]int{}
		intervals := [][2]int{}
		for _, interval := range p[step.part] {
			if step.lower {
				if interval[0] < step.value {
					newIntervals = append(newIntervals, [2]int{interval[0], min(interval[1], step.value-1)})
				}
				if interval[1] >= step.value {
					intervals = append(intervals, [2]int{max(interval[0], step.value), interval[1]})
				}
			} else {
				if interval[1] > step.value {
					newIntervals = append(newIntervals, [2]int{max(interval[0], step.value+1), interval[1]})
				}
				if interval[0] <= step.value {
					intervals = append(intervals, [2]int{interval[0], min(interval[1], step.value)})
				}
			}
		}
		p[step.part] = intervals
		newP[step.part] = newIntervals
		n += count(newP, workflows, step.next)
	}
	return n + count(p, workflows, w.redirect)
}

func Resolve2(lines []string) (string) {
	res := 0
	workflows := map[string](workflow){}
	workflowsAllRead := false
	for _, line := range lines {
		if len(line) == 0 {
			workflowsAllRead = true
			continue
		}
		if !workflowsAllRead {
			name := strings.Split(line, "{")[0]
			datas := strings.Split(strings.Split(strings.Split(line, "{")[1], "}")[0], ",")
			steps := []step{}
			for _, data := range datas[:len(datas)-1] {
				value, _ := strconv.Atoi(strings.Split(data, ":")[0][2:])
				next := strings.Split(data, ":")[1]
				steps = append(steps, step{
					part: string(data[0]),
					lower: string(data[1]) == "<",
					value : value,
					next: next})
			}
			workflows[name] = workflow{steps: steps, redirect: datas[len(datas)-1]}
		}
	}
	p := map[string][][2]int{
		"x": [][2]int{[2]int{1, 4000}},
		"m": [][2]int{[2]int{1, 4000}},
		"a": [][2]int{[2]int{1, 4000}},
		"s": [][2]int{[2]int{1, 4000}}}
	res = count(p, workflows, "in")
	return strconv.Itoa(res)
}
