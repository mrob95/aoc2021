package main

import (
	"fmt"
	"math"
)

// .._DB._AC._DB._CA..
// 11 corridor positions, then 8 rooms
type State struct {
	positions [19]rune
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
	for i, v := range []rune{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D'} {
		if s.positions[11+i] != v {
			return false
		}
	}
	return true
}

func possibleMoves(s State, i int) []int {
	c := s.positions[i]
	if i < 11 { // in corridor, can move into final position?
		targetRoom := int(c - 'A')
		top := 11 + 2*targetRoom
		bottom := top + 1
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
		if s.positions[top] == '.' && s.positions[bottom] == '.' {
			return []int{bottom}
		} else if s.positions[top] == '.' && s.positions[bottom] == c {
			return []int{top}
		} else {
			return []int{}
		}
	} else { // in room, where can we move?
		roomPos := i - 11
		room := roomPos / 2
		atBottom := roomPos%2 == 1
		if atBottom && room == int(c-'A') {
			return []int{} // already in position
		} else if s.positions[11+room*2+1] == c && room == int(c-'A') {
			return []int{} // already in position
		} else if atBottom && s.positions[11+room*2] != '.' {
			return []int{} // blocked
		} else { // at
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
	}
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
		room := roomPos / 2
		entrance := (room * 2) + 2
		distanceToEntrance := diff(entrance, from)
		return distanceToEntrance + roomPos%2 + 1
	} else {
		if to < 11 {
			roomPos := from - 11
			room := roomPos / 2
			entrance := (room * 2) + 2
			distanceFromEntrance := diff(entrance, to)
			distanceToEntrance := roomPos%2 + 1
			return distanceToEntrance + distanceFromEntrance
		} else {
			fromRoomPos := from - 11
			fromRoom := fromRoomPos / 2
			toRoomPos := to - 11
			toRoom := toRoomPos / 2

			roomJumps := diff(fromRoom, toRoom)
			distance := (roomJumps * 2)
			return (fromRoomPos%2 + 1) + distance + (toRoomPos%2 + 1)
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
	// state := newState("BACDBCDA")
	state := newState("DBACDBCA")

	cache := make(map[State]int)
	fmt.Println(minimumEnergy(state, cache))
}
