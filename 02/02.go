package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	horizontalPosition, depth, aim int
}

func main() {
	pos1 := Position{0, 0, 0} // Part 1
	pos2 := Position{0, 0, 0} // Part 2

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		value, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			pos1.horizontalPosition += int(value)
			pos2.horizontalPosition += int(value)
			pos2.depth += int(value) * pos2.aim
		case "down":
			pos1.depth += int(value)
			pos2.aim += int(value)
		case "up":
			pos1.depth -= int(value)
			pos2.aim -= int(value)
		}
	}
	fmt.Printf("Part 1: %d\n", pos1.horizontalPosition*pos1.depth)
	fmt.Printf("Part 2: %d\n", pos2.horizontalPosition*pos2.depth)
}
