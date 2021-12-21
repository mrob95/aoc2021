package main

import "fmt"

type Player struct {
	current, score int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// players := [2]Player{{4, 0}, {8, 0}}
	players := [2]Player{{10, 0}, {2, 0}}
	turn := 0
	dice := 1

	for players[0].score < 1000 && players[1].score < 1000 {
		roll := 0
		for i := 0; i < 3; i++ {
			roll += dice
			dice = (dice % 100) + 1
		}
		turn_idx := turn % 2
		players[turn_idx].current = (((players[turn_idx].current - 1) + roll) % 10) + 1
		players[turn_idx].score += players[turn_idx].current
		turn++
	}

	rolls := turn * 3
	fmt.Println(min(players[0].score, players[1].score) * rolls)
}
