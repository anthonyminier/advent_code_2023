package main

import (
	"regexp"
	"strings"
)

func day8(input string) int {
	instructions := ""
	reg := regexp.MustCompile(`^([1-9A-Z]{3}) = \(([1-9A-Z]{3}), ([1-9A-Z]{3})\)$`)
	nbLine := 0
	nodes := map[string]map[string]string{}
	fromInput(input, func(line string) {
		defer func() { nbLine++ }()
		if nbLine == 0 {
			instructions = strings.TrimSpace(line)
			return
		}
		if !reg.MatchString(line) {
			return
		}
		matches := reg.FindAllStringSubmatch(line, -1)
		nodes[matches[0][1]] = map[string]string{`L`: matches[0][2], `R`: matches[0][3]}
	})
	step := 0
	node := "AAA"

	for node != "ZZZ" {
		for _, instruction := range instructions {
			node = nodes[node][string(instruction)]
			step++
		}
	}
	return step
}

func day8_2(input string) int {
	instructions := ""
	reg := regexp.MustCompile(`^([1-9A-Z]{3}) = \(([1-9A-Z]{3}), ([1-9A-Z]{3})\)$`)
	nbLine := 0
	mapNodes := map[string]map[string]string{}
	fromInput(input, func(line string) {
		defer func() { nbLine++ }()
		if nbLine == 0 {
			instructions = strings.TrimSpace(line)
			return
		}
		if !reg.MatchString(line) {
			return
		}
		matches := reg.FindAllStringSubmatch(line, -1)
		mapNodes[matches[0][1]] = map[string]string{`L`: matches[0][2], `R`: matches[0][3]}
	})
	step := 0
	endCondition := func(nodes map[string]map[string]string) bool {
		for node := range nodes {
			if !strings.HasSuffix(node, "Z") {
				return false
			}
		}
		return true
	}
	nodes := map[string]map[string]string{}
	for node, paths := range mapNodes {
		if strings.HasSuffix(node, "A") {
			nodes[node] = paths
		}
	}
	for !endCondition(nodes) {
		for _, instruction := range instructions {
			newNode := map[string]map[string]string{}
			for _, paths := range nodes {
				node := paths[string(instruction)]
				newNode[node] = mapNodes[node]
			}
			step++
			nodes = newNode
			if endCondition(nodes) {
				break
			}
		}
	}
	return step
}
