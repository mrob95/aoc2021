package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

var algorithm [512]bool

type Image struct {
	lit          map[Point]interface{}
	min_x, min_y int
	max_x, max_y int
}

func (i *Image) add(p Point) {
	i.lit[p] = nil
	if p.x < i.min_x {
		i.min_x = p.x
	}
	if p.x > i.max_x {
		i.max_x = p.x
	}
	if p.y < i.min_y {
		i.min_y = p.y
	}
	if p.y > i.max_y {
		i.max_y = p.y
	}
}

func (i *Image) numLit() int {
	return len(i.lit)
}

func (i *Image) display() {
	for y := i.min_y; y <= i.max_y; y++ {
		for x := i.min_x; x <= i.max_x; x++ {
			p := Point{x, y}
			if _, lit := i.lit[p]; lit {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (image *Image) apply(outersLit bool) Image {
	new := Image{}
	new.lit = make(map[Point]interface{})
	new.min_x = image.min_x - 1
	new.min_y = image.min_y - 1
	new.max_x = image.max_x + 1
	new.max_y = image.max_y + 1

	for y := new.min_y; y <= new.max_y; y++ {
		for x := new.min_x; x <= new.max_x; x++ {
			i := 8
			algo_idx := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					p := Point{x + dx, y + dy}
					atEdge := p.x <= new.min_x || p.x >= new.max_x || p.y <= new.min_y || p.y >= new.max_x
					if _, lit := image.lit[p]; lit {
						algo_idx |= (1 << i)
					}
					if outersLit && atEdge {
						// Edge of the region, but these are lit every other application
						algo_idx |= (1 << i)
					}
					i--
				}
			}
			newStatus := algorithm[algo_idx]
			if newStatus {
				new.add(Point{x, y})
			}
		}
	}
	return new
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	algorithmString := scanner.Text()
	for i, c := range algorithmString {
		algorithm[i] = c == '#'
	}

	scanner.Scan()
	image := Image{}
	image.lit = make(map[Point]interface{})
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, v := range line {
			if v == '#' {
				image.add(Point{x, y})
			}
		}
		y++
	}

	next := image
	for i := 0; i < 50; i++ {
		// If algorithm[0] == '#' then the edge of the region will flip flop from lit to unlit
		next = next.apply((i%2 == 1) && algorithm[0])
	}
	next.display()
	fmt.Println(next.numLit())
}
