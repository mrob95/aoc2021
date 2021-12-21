package main

import "fmt"

type Player struct {
	position, score int
}

type State struct {
	players [2]Player
	turn    int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numVictories(state State, cache map[State][2]int) [2]int {
	if cachedResult, exists := cache[state]; exists {
		return cachedResult
	}
	var result [2]int
	for roll1 := 1; roll1 <= 3; roll1++ {
		for roll2 := 1; roll2 <= 3; roll2++ {
			for roll3 := 1; roll3 <= 3; roll3++ {
				roll := roll1 + roll2 + roll3
				newState := state
				newPos := (((newState.players[newState.turn].position - 1) + roll) % 10) + 1
				newState.players[newState.turn].position = newPos
				newState.players[newState.turn].score += newPos
				if newState.players[newState.turn].score >= 21 {
					result[newState.turn]++
				} else {
					newState.turn = (newState.turn + 1) % 2
					victories := numVictories(newState, cache)
					result[0] += victories[0]
					result[1] += victories[1]
				}
			}
		}
	}
	cache[state] = result
	return result
}

func main() {
	// state := State{[2]Player{{4, 0}, {8, 0}}, 0}
	state := State{[2]Player{{10, 0}, {2, 0}}, 0}
	cache := make(map[State][2]int)
	fmt.Println(numVictories(state, cache))
}
