package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var readings []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		readings = append(readings, int(i))
	}

	part1 := 0
	for i := 1; i < len(readings); i++ {
		if readings[i] > readings[i-1] {
			part1 += 1
		}
	}
	fmt.Println(part1)

	part2 := 0
	for i := 3; i < len(readings); i++ {
		windowA := readings[i-1] + readings[i-2] + readings[i-3]
		windowB := readings[i] + readings[i-1] + readings[i-2]
		if windowB > windowA {
			part2 += 1
		}
	}
	fmt.Println(part2)
}
