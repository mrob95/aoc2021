package main

import (
	"fmt"
	"math"
)

// 11 corridor positions, then 16 rooms
type State struct {
	positions [27]rune
	energy    int
}

func newState(in string) State {
	var result State
	for i := 0; i < 11; i++ {
		result.positions[i] = '.'
	}
	for i, c := range in {
		result.positions[11+i] = c
	}
	return result
}

func finished(s State) bool {
	for i, v := range []rune{'A', 'A', 'A', 'A', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'C', 'D', 'D', 'D', 'D'} {
		if s.positions[11+i] != v {
			return false
		}
	}
	return true
}

func possibleMoves(s State, i int) []int {
	c := s.positions[i]
	targetRoom := int(c - 'A')
	if i < 11 { // in corridor, can move into final position?
		top := 11 + 4*targetRoom
		entrance := (targetRoom * 2) + 2

		if entrance > i {
			for j := i + 1; j <= entrance; j++ {
				if s.positions[j] != '.' {
					return []int{} // blocked
				}
			}
		} else {
			for j := i - 1; j >= entrance; j-- {
				if s.positions[j] != '.' {
					return []int{} // blocked
				}
			}
		}

		for j := top + 3; j >= top; j-- {
			if s.positions[j] == '.' {
				return []int{j}
			} else if s.positions[j] != c {
				return []int{} // room not ready yet
			}
		}
	} else { // in room, where can we move?
		roomPos := i - 11
		room := roomPos / 4
		top := 11 + (room * 4)
		for j := i - 1; j >= top; j-- {
			if s.positions[j] != '.' {
				return []int{} // blocked above
			}
		}

		if room == targetRoom {
			done := true
			for j := i + 1; j < top+4; j++ {
				if s.positions[j] != c {
					done = false
					break
				}
			}
			if done { // in target room and no others below
				return []int{}
			}
		}

		var result []int
		entrance := (room * 2) + 2
		for j := entrance; j >= 0; j-- { // try left
			if s.positions[j] != '.' { // blocked
				break
			}
			if j > 1 && (j-2)%2 == 0 { // is an entrance
				continue
			}
			result = append(result, j)
		}
		for j := entrance; j < 11; j++ { // try right
			if s.positions[j] != '.' { // blocked
				break
			}
			if j < 9 && (j-2)%2 == 0 { // is an entrance
				continue
			}
			result = append(result, j)
		}
		return result
	}
	panic("Unreachable")
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func distanceForMove(from, to int) int {
	if from < 11 { // must be going into a room
		roomPos := to - 11
		room := roomPos / 4
		entrance := (room * 2) + 2
		distanceToEntrance := diff(entrance, from)
		return distanceToEntrance + roomPos%4 + 1
	} else {
		if to < 11 {
			roomPos := from - 11
			room := roomPos / 4
			entrance := (room * 2) + 2
			distanceFromEntrance := diff(entrance, to)
			distanceToEntrance := roomPos%4 + 1
			return distanceToEntrance + distanceFromEntrance
		} else {
			fromRoomPos := from - 11
			fromRoom := fromRoomPos / 4
			toRoomPos := to - 11
			toRoom := toRoomPos / 4

			roomJumps := diff(fromRoom, toRoom)
			distance := (roomJumps * 2)
			return (fromRoomPos%4 + 1) + distance + (toRoomPos%4 + 1)
		}
	}
}

func minimumEnergy(s State, cache map[State]int) int {
	if finished(s) {
		return s.energy
	}
	if v, exists := cache[s]; exists {
		return v
	}

	result := math.MaxInt32
	for start, c := range s.positions {
		if c != '.' {
			for _, finish := range possibleMoves(s, start) {
				distance := distanceForMove(start, finish)
				energy := distance
				for i := 0; i < int(c-'A'); i++ {
					energy *= 10
				}

				newState := s
				newState.energy += energy
				newState.positions[start] = '.'
				newState.positions[finish] = c
				minimum := minimumEnergy(newState, cache)
				if minimum < result {
					result = minimum
				}
			}
		}
	}
	cache[s] = result
	return result
}

func main() {
	// exampleState := newState("BDDACCBDBBACDACA")
	state := newState("DDDBACBCDBABCACA")

	cache := make(map[State]int)
	fmt.Println(minimumEnergy(state, cache))

}
