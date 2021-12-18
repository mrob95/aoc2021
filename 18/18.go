package main

import (
	"bufio"
	"fmt"
	"os"
)

type Number struct {
	digits, depths []int
}

func parsePair(s string) Number {
	result := Number{}
	depth := 0
	for _, c := range s {
		switch c {
		case '[':
			depth++
		case ',':
			continue
		case ']':
			depth--
		default:
			result.digits = append(result.digits, int(c-'0'))
			result.depths = append(result.depths, depth)
		}
	}
	return result
}

func (n *Number) explode() bool {
	for i, d := range n.depths {
		if d == 5 {
			// Increase left
			if i-1 >= 0 {
				n.digits[i-1] += n.digits[i]
			}
			// Increase right
			if i+2 < len(n.digits) {
				n.digits[i+2] += n.digits[i+1]
			}
			// Boom
			n.digits = append(n.digits[:i], n.digits[i+1:]...)
			n.digits[i] = 0
			n.depths = append(n.depths[:i], n.depths[i+1:]...)
			n.depths[i] = 4
			return true
		}
	}
	return false
}

func (n *Number) split() bool {
	for i, v := range n.digits {
		if v > 9 {
			left, right := v/2, v/2
			if v%2 != 0 {
				right += 1
			}
			n.digits = append(n.digits[:i+1], n.digits[i:]...)
			n.digits[i] = left
			n.digits[i+1] = right
			n.depths = append(n.depths[:i+1], n.depths[i:]...)
			n.depths[i]++
			n.depths[i+1]++
			return true
		}
	}
	return false
}

func (n *Number) reduce() {
	for {
		if n.explode() {
			continue
		}
		if n.split() {
			continue
		}
		return
	}
}

func (n *Number) add(right Number) {
	n.digits = append(n.digits, right.digits...)
	n.depths = append(n.depths, right.depths...)
	for i := 0; i < len(n.depths); i++ {
		n.depths[i]++
	}
	n.reduce()
}

func (n *Number) magnitude() int {
	// Note that this destroys the input
	for len(n.digits) > 1 {
		for i := 0; i < len(n.digits)-1; i++ {
			if n.depths[i] == n.depths[i+1] {
				depth := n.depths[i]
				left, right := n.digits[i], n.digits[i+1]
				new := 3*left + 2*right
				n.digits = append(n.digits[:i], n.digits[i+1:]...)
				n.digits[i] = new
				n.depths = append(n.depths[:i], n.depths[i+1:]...)
				n.depths[i] = depth - 1
				break
			}
		}
	}
	return n.digits[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// P1
	n := parsePair(lines[0])
	for i := 1; i < len(lines); i++ {
		n.add(parsePair(lines[i]))
	}
	fmt.Println(n.magnitude())

	// P2
	highest := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			l := parsePair(lines[i])
			r := parsePair(lines[j])
			l.add(r)
			mag := l.magnitude()
			if mag > highest {
				highest = mag
			}
		}
	}
	fmt.Println(highest)
}
