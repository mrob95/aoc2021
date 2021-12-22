package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x, y, z int
}

type Instruction struct {
	on                     bool
	x1, x2, y1, y2, z1, z2 int
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
		inst := Instruction{line[1] == 'n', x1, x2, y1, y2, z1, z2}
		instructions = append(instructions, inst)
	}

	fmt.Println(instructions)

	grid := make(map[Point]interface{})

	for _, inst := range instructions {
		for x := max(inst.x1, -50); x <= min(inst.x2, 50); x++ {
			for y := max(inst.y1, -50); y <= min(inst.y2, 50); y++ {
				for z := max(inst.z1, -50); z <= min(inst.z2, 50); z++ {
					p := Point{x, y, z}
					if inst.on {
						grid[p] = nil
					} else {
						delete(grid, p)
					}
				}
			}
		}
	}
	fmt.Println(len(grid))

}
