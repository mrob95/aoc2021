package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Cuboid struct {
	x1, x2, y1, y2, z1, z2 int
}

type Instruction struct {
	on bool
	c  Cuboid
}

func overlap(a1, a2, b1, b2 int) bool {
	return (b1 >= a1 && b1 <= a2) || (a1 >= b1 && a1 <= b2)
}

func intersect(a, b Cuboid) bool {
	return overlap(a.x1, a.x2, b.x1, b.x2) && overlap(a.y1, a.y2, b.y1, b.y2) && overlap(a.z1, a.z2, b.z1, b.z2)
}

func spans(a1, a2, b1, b2 int) [][2]int {
	// All segments within a1, a2
	if b1 <= a1 && b2 >= a2 {
		return [][2]int{{a1, a2}}
	} else if b1 <= a1 && b2 < a2 {
		return [][2]int{{a1, b2}, {b2 + 1, a2}}
	} else if b1 > a1 && b2 < a2 {
		return [][2]int{{a1, b1 - 1}, {b1, b2}, {b2 + 1, a2}}
	} else if b1 > a1 && b2 >= a2 {
		return [][2]int{{a1, b1 - 1}, {b1, a2}}
	}
	panic("Unreachable")
}

func subtract(a, b Cuboid) []Cuboid {
	// Return cuboids that are within a but not within b
	var result []Cuboid
	xss := spans(a.x1, a.x2, b.x1, b.x2)
	yss := spans(a.y1, a.y2, b.y1, b.y2)
	zss := spans(a.z1, a.z2, b.z1, b.z2)
	for _, xs := range xss {
		for _, ys := range yss {
			for _, zs := range zss {
				c := Cuboid{xs[0], xs[1], ys[0], ys[1], zs[0], zs[1]}
				if !intersect(c, b) {
					result = append(result, c)
				}
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func volume(c Cuboid) int {
	return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func volumeSum(cs []Cuboid) int {
	result := 0
	for _, c := range cs {
		result += volume(c)
	}
	return result
}

func main() {
	var instructions []Instruction
	scanner := bufio.NewScanner(os.Stdin)
	re := regexp.MustCompile("-?\\d+")
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		x1, _ := strconv.Atoi(matches[0])
		x2, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		y2, _ := strconv.Atoi(matches[3])
		z1, _ := strconv.Atoi(matches[4])
		z2, _ := strconv.Atoi(matches[5])
		inst := Instruction{line[1] == 'n', Cuboid{x1, x2, y1, y2, z1, z2}}
		instructions = append(instructions, inst)
	}

	var on []Cuboid

	for _, inst := range instructions {
		// Subtract the cuboid in this instruction from all of the currently lit cuboids,
		// then if lit, add it.
		var new []Cuboid
		for _, c := range on {
			if intersect(inst.c, c) {
				subbed := subtract(c, inst.c)
				new = append(new, subbed...)
			} else {
				new = append(new, c)
			}
		}
		if inst.on {
			new = append(new, inst.c)
		}
		on = new
	}

	fmt.Println(volumeSum(on))
}
