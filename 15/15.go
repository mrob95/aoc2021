package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y int
}

type HeapKey struct {
	p        Point
	distance int
}

type Heap struct {
	keys      []HeapKey
	positions map[Point]int
}

func (h *Heap) add(p Point, d int) {
	h.keys = append(h.keys, HeapKey{p, d})
	pos := len(h.keys) - 1
	h.positions[p] = pos
	h.swim(pos)
}

func (h *Heap) set(p Point, d int) {
	pos, _ := h.positions[p]
	h.keys[pos].distance = d
	h.swim(pos)
}

func (h *Heap) exists(p Point) bool {
	_, ok := h.positions[p]
	return ok
}

func (h *Heap) empty() bool {
	return len(h.keys) == 0
}

func (h *Heap) swap(a, b int) {
	tmp_a := h.keys[a]
	tmp_b := h.keys[b]
	h.keys[a] = tmp_b
	h.keys[b] = tmp_a
	h.positions[tmp_a.p] = b
	h.positions[tmp_b.p] = a
}

func (h *Heap) swim(i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	if h.keys[parent].distance > h.keys[i].distance {
		h.swap(i, parent)
		h.sink(parent)
		h.swim(parent)
	}
}

func (h *Heap) sink(i int) {
	child1, child2 := 2*i+1, 2*i+2
	if child1 < len(h.keys) && h.keys[child1].distance < h.keys[i].distance {
		h.swap(i, child1)
		h.sink(child1)
		h.sink(i)
	} else if child2 < len(h.keys) && h.keys[child2].distance < h.keys[i].distance {
		h.swap(i, child2)
		h.sink(child2)
		h.sink(i)
	}
}

func (h *Heap) pop() HeapKey {
	result := h.keys[0]
	h.swap(0, len(h.keys)-1)
	h.keys = h.keys[0 : len(h.keys)-1]
	delete(h.positions, result.p)
	h.sink(0)
	return result
}

var grid map[Point]int
var width, height int

func adjacentPoints(p Point) []Point {
	return []Point{
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
	}
}

func shortestPath(start, end Point) int {
	distances := make(map[Point]int)
	q := Heap{}
	q.positions = make(map[Point]int)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := Point{x, y}
			if p != start {
				distances[p] = math.MaxInt32
				q.add(p, math.MaxInt32)
			}
		}
	}
	q.add(start, 0)
	distances[start] = 0

	for !q.empty() {
		current := q.pop().p
		for _, p := range adjacentPoints(current) {
			if !q.exists(p) {
				continue
			}
			alternative := distances[current] + grid[p]
			currentLowest := distances[p]
			if alternative < currentLowest {
				distances[p] = alternative
				q.set(p, alternative)
			}
		}
	}
	result := distances[end]
	return result
}

func main() {
	grid = make(map[Point]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			val := int(line[i] - '0')
			grid[Point{i, height}] = val
		}
		width = len(line)
		height += 1
	}

	// Part 1
	start := Point{0, 0}
	end := Point{width - 1, height - 1}
	fmt.Println(shortestPath(start, end))

	// Part 2
	newGrid := make(map[Point]int)
	for p, val := range grid {
		for dx := 0; dx < 5; dx++ {
			for dy := 0; dy < 5; dy++ {
				newP := Point{p.x + dx*width, p.y + dy*height}
				extra := dx + dy
				newGrid[newP] = ((val + extra - 1) % 9) + 1
			}
		}
	}

	grid = newGrid
	width = 5 * width
	height = 5 * height
	end = Point{width - 1, height - 1}
	fmt.Println(shortestPath(start, end))
}
