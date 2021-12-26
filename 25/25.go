package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var seafloor [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		lineCopy := append([]byte{}, line...)
		seafloor = append(seafloor, lineCopy)
	}

	ymax := len(seafloor)
	xmax := len(seafloor[0])
	step := 0
	moved := true
	for moved {
		moved = false

		eastMoves := [][2]int{}
		for y := 0; y < ymax; y++ {
			for x := 0; x < xmax; x++ {
				if seafloor[y][x] == '>' && seafloor[y][(x+1)%xmax] == '.' {
					eastMoves = append(eastMoves, [2]int{y, x})
					moved = true
				}
			}
		}
		for _, m := range eastMoves {
			seafloor[m[0]][m[1]] = '.'
			seafloor[m[0]][(m[1]+1)%xmax] = '>'
		}

		southMoves := [][2]int{}
		for y := 0; y < ymax; y++ {
			for x := 0; x < xmax; x++ {
				if seafloor[y][x] == 'v' && seafloor[(y+1)%ymax][x] == '.' {
					southMoves = append(southMoves, [2]int{y, x})
					moved = true
				}
			}
		}
		for _, m := range southMoves {
			seafloor[m[0]][m[1]] = '.'
			seafloor[(m[0]+1)%ymax][m[1]] = 'v'
		}
		step++
	}
	fmt.Println(step)
}
