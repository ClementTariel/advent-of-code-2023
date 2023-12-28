package main

import (
	"strings"
	"strconv"
)

type module struct {
	modType string
	name string
	output []string
	memory map[string]bool
	flip bool
}

type pulse struct {
	name string
	high bool
}

func CopyModule(mod module) (module) {
	modCopy := module{
		modType: mod.modType,
		name: mod.name,
		output: []string{},
		memory: map[string]bool{},
		flip: mod.flip}
	for _, v := range mod.output {
		modCopy.output = append(modCopy.output, v)
	}
	for k, v := range mod.memory {
		modCopy.memory[k] = v
	}
	return modCopy

}

func CopyModules(modules map[string]module) (map[string]module) {
	modulesCopy := map[string]module{}
	for name, mod := range modules {
		modCopy := CopyModule(mod)
		modulesCopy[name] = modCopy
	}
	return modulesCopy
}

func PulseStep(pulses []pulse, modules map[string]module) (map[string]module ,int, int) {
	lowCount := 0
	highCount := 0

	nextPulses := []pulse{}

	modulesCopy := CopyModules(modules)
	for _, p := range pulses {
		mod, exists := modules[p.name]
		modCopy, existsCopy := modulesCopy[p.name]
		if !exists || !existsCopy {
			panic(p.name)
		}
		var nextHigh bool
		emit := true
		if mod.modType == "%" {
			if !p.high {
				modCopy.flip = !mod.flip
				modulesCopy[p.name] = modCopy
				nextHigh = modCopy.flip
			} else {
				emit = false
			}
		} else if mod.modType == "&" {
			allHigh := false
			for _, v := range mod.memory {
				allHigh = v
				if !allHigh {
					break
				}
			}
			nextHigh = !allHigh
		} else {
			//broadcast
			nextHigh = p.high
		}
		if emit {
			for _, dest := range mod.output {
				//fmt.Println(p.name, "-",nextHigh,"->", dest)
				if nextHigh {
					highCount += 1
				} else {
					lowCount += 1
				}
				nextMod, existsNext := modules[dest]
				if existsNext {
					if nextMod.modType == "&" {
						modulesCopy[dest].memory[p.name] = nextHigh
					}
				nextPulses = append(nextPulses, pulse{name: dest, high: nextHigh})
				}
			}
		}
	}
	l := 0
	h := 0
	if len(nextPulses) > 0 {
		modules, l, h = PulseStep(nextPulses, modulesCopy)
	}
	lowCount += l
	highCount += h

	return modules, lowCount, highCount
}

func Resolve1(lines []string) (string) {
	res := 0
	modules := map[string]module{}
	for _, line := range lines {
		data := strings.Split(line, " -> ")
		name := ""
		modType := ""
		if data[0] == "broadcaster" {
			name = data[0]
			modType = data[0]
		} else {
			name = string(data[0][1:])
			modType = string(data[0][0])
		}
		output := []string{}
		for _, dest := range strings.Split(data[1], ", ") {
			output = append(output, dest)
		}
		memory := map[string]bool{}
		flip := false
		mod := module{modType: modType, name: name, output: output, memory: memory, flip: flip}
		modules[name] = mod
	}
	for _, mod := range modules {
		for _, name := range mod.output {
			nextMod, exists := modules[name]
			if exists && nextMod.modType == "&" {
				modules[name].memory[mod.name] = false
			}
		}
	}
	firstPulse := []pulse{pulse{name: "broadcaster", high: false}}
	lowCount := 0
	highCount := 0
	l := 0
	h := 0
	for k := 0; k < 1000; k++ {
		modules, l, h = PulseStep(firstPulse, modules)
		l += 1 // button
		lowCount += l
		highCount += h
	}
	res = lowCount * highCount
	return strconv.Itoa(res)
}

