package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	var crabs []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fishStrings := strings.Split(line, ",")
	for _, s := range fishStrings {
		i, _ := strconv.Atoi(s)
		crabs = append(crabs, i)
	}

	smallest, largest := 9999, 0
	for _, v := range crabs {
		if v < smallest {
			smallest = v
		} else if v > largest {
			largest = v
		}
	}

	smallestResult1 := -1
	smallestResult2 := -1
	for target := smallest; target <= largest; target++ {
		result1 := 0
		result2 := 0
		for _, v := range crabs {
			distance := abs(v - target)
			result1 += distance
			result2 += (distance * (distance + 1)) / 2

		}
		if smallestResult1 == -1 || result1 < smallestResult1 {
			smallestResult1 = result1
		}
		if smallestResult2 == -1 || result2 < smallestResult2 {
			smallestResult2 = result2
		}
	}

	fmt.Println(smallestResult1)
	fmt.Println(smallestResult2)
}
