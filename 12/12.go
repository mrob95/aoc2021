package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cave struct {
	id  string
	big bool
}

type CaveMap struct {
	connections map[string][]Cave
}

func (c *CaveMap) add(left, right string) {
	if left == "end" || right == "start" {
		return
	}
	right_big := strings.ToUpper(right) == right
	c.connections[left] = append(c.connections[left], Cave{right, right_big})
}

func (c *CaveMap) numPaths(from string, visited map[string]interface{}, visitedTwice bool, path []string) int {
	// DFS, ignoring small caves we have already visited
	path = append(path, from)
	if from == "end" {
		// fmt.Println(path)
		return 1
	}
	result := 0
	visited[from] = nil
	for _, v := range c.connections[from] {
		if _, ok := visited[v.id]; ok && !v.big { // small cave, already visited
			if !visitedTwice {
				result += c.numPaths(v.id, visited, true, path)
				visited[v.id] = nil // HACK - this gets deleted at the end of the above call, but in this case, where we are visiting for the second time, we want to keep the record of the first visit
			}
			continue
		}
		result += c.numPaths(v.id, visited, visitedTwice, path)
	}
	delete(visited, from)
	return result
}

func main() {
	m := CaveMap{make(map[string][]Cave)}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		m.add(parts[0], parts[1])
		m.add(parts[1], parts[0])
	}

	visited := make(map[string]interface{})
	var path []string

	part1 := m.numPaths("start", visited, true, path)
	fmt.Println(part1)

	part2 := m.numPaths("start", visited, false, path)
	fmt.Println(part2)
}
