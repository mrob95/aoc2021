package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Fold struct {
	direction string
	location  int
}

func fold(paper map[Point]interface{}, f Fold) {
	for p, _ := range paper {
		if f.direction == "y" && p.y > f.location {
			new := Point{p.x, 2*f.location - p.y}
			paper[new] = nil
			delete(paper, p)
		} else if f.direction == "x" && p.x > f.location {
			new := Point{2*f.location - p.x, p.y}
			paper[new] = nil
			delete(paper, p)
		}
	}
}

func display(paper map[Point]interface{}) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 100; x++ {
			if _, ok := paper[Point{x, y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	paper := make(map[Point]interface{})
	var folds []Fold

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		point := Point{x, y}
		paper[point] = nil
	}

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("[xy]=\\d+")
		match := re.FindString(line)
		parts := strings.Split(match, "=")
		location, _ := strconv.Atoi(parts[1])
		folds = append(folds, Fold{parts[0], location})
	}

	for i, f := range folds {
		if i == 0 {
			fmt.Println(len(paper))
		}
		fold(paper, f)
	}
	display(paper)
}
