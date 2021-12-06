package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func parseLines(r io.Reader) []Line {
	var lines []Line
	re := regexp.MustCompile("\\d+")
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		x1, _ := strconv.Atoi(matches[0])
		y1, _ := strconv.Atoi(matches[1])
		x2, _ := strconv.Atoi(matches[2])
		y2, _ := strconv.Atoi(matches[3])

		// Always pointing right
		if x1 <= x2 {
			lines = append(lines, Line{Point{x1, y1}, Point{x2, y2}})
		} else {
			lines = append(lines, Line{Point{x2, y2}, Point{x1, y1}})
		}
	}
	return lines
}

func increment(counts map[Point]int, p Point) {
	if _, ok := counts[p]; ok {
		counts[p] += 1
	} else {
		counts[p] = 1
	}
}

func overlaps(counts map[Point]int) int {
	result := 0
	for _, hits := range counts {
		if hits > 1 {
			result += 1
		}
	}
	return result
}

func intermediatePoints(p1, p2 Point, diagonal bool) []Point {
	var result []Point
	if p1.x == p2.x { // vertical
		if p1.y <= p2.y {
			for i := p1.y; i <= p2.y; i++ {
				result = append(result, Point{p1.x, i})
			}
		} else {
			for i := p2.y; i <= p1.y; i++ {
				result = append(result, Point{p1.x, i})
			}
		}
	} else if p1.y == p2.y { // horizontal
		for i := p1.x; i <= p2.x; i++ {
			result = append(result, Point{i, p1.y})
		}
	} else if diagonal && p1.y <= p2.y { // down right
		for i := 0; i <= p2.x-p1.x; i++ {
			result = append(result, Point{p1.x + i, p1.y + i})
		}
	} else if diagonal && p1.y > p2.y { // up right
		for i := 0; i <= p2.x-p1.x; i++ {
			result = append(result, Point{p1.x + i, p1.y - i})
		}
	}
	return result
}

func main() {
	lines := parseLines(os.Stdin)

	counts1 := make(map[Point]int)
	counts2 := make(map[Point]int)
	for _, line := range lines {
		for _, p := range intermediatePoints(line.start, line.end, false) {
			increment(counts1, p)
		}
		for _, p := range intermediatePoints(line.start, line.end, true) {
			increment(counts2, p)
		}
	}

	fmt.Println(overlaps(counts1))
	fmt.Println(overlaps(counts2))
}
