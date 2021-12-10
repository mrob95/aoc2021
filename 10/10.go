package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var p1_score_map = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var p2_score_map = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var matches = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	p1total := 0
	var p2scores []int

	for scanner.Scan() {
		var stack []rune
		corrupt := false

		line := scanner.Text()
		for _, char := range line {
			if match, ok := matches[char]; ok {
				if stack[len(stack)-1] == match {
					stack = stack[:len(stack)-1]
				} else {
					score, _ := p1_score_map[char]
					p1total += score
					corrupt = true
					break
				}
			} else {
				stack = append(stack, char)
			}
		}

		if len(stack) > 0 && !corrupt { // Incomplete
			lineScore := 0
			for i := len(stack) - 1; i >= 0; i-- {
				s, _ := p2_score_map[stack[i]]
				lineScore *= 5
				lineScore += s
			}
			p2scores = append(p2scores, lineScore)
		}
	}

	fmt.Println(p1total) // 1

	sort.Ints(p2scores)
	middle := len(p2scores) / 2
	fmt.Println(p2scores[middle]) // 2
}
