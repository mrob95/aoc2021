package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 10

type Octopus struct {
	power   int
	flashed bool
}

type Cave [SIZE][SIZE]Octopus

func (cave *Cave) flash(x, y int) {
	// Flash an octopus and bump the power of adjacent octopi, flashing them too if necessary
	cave[y][x].flashed = true

	for ya := -1; ya <= 1; ya++ {
		for xa := -1; xa <= 1; xa++ {
			if y+ya < 0 || y+ya >= SIZE || x+xa < 0 || x+xa >= SIZE || (xa == 0 && ya == 0) {
				continue //bounds
			}
			cave[y+ya][x+xa].power += 1
			if cave[y+ya][x+xa].power > 9 && !cave[y+ya][x+xa].flashed {
				cave.flash(x+xa, y+ya)
			}
		}
	}
}

func (cave *Cave) step() int {
	// simulate
	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			cave[y][x].power += 1
			if cave[y][x].power > 9 && !cave[y][x].flashed {
				cave.flash(x, y)
			}
		}
	}

	// cleanup and count flashes
	flashes := 0
	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			if cave[y][x].flashed {
				cave[y][x].flashed = false
				cave[y][x].power = 0
				flashes += 1
			}
		}
	}
	return flashes
}

func (cave *Cave) display() {
	// Display the current power levels
	for y := 0; y < SIZE; y++ {
		var ar [SIZE]int
		for x := 0; x < SIZE; x++ {
			ar[x] = cave[y][x].power
		}
		fmt.Println(ar)
	}
}

func main() {
	cave := Cave{}
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, power := range line {
			cave[y][x] = Octopus{int(power - '0'), false}
		}
		y++
	}

	part1 := 0
	i := 0
	for {
		flashes := cave.step()
		if i < 100 {
			part1 += flashes
		} else if i == 100 {
			fmt.Println(part1)
		} else {
			if flashes == 100 {
				fmt.Println(i + 1)
				break
			}
		}
		i += 1
	}
}
