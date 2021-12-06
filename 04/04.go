package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func splitAndParse(s, sep string) []int {
	var result []int
	for _, is := range strings.Split(s, sep) {
		i, _ := strconv.Atoi(is)
		result = append(result, i)
	}
	return result
}

func splitAndParseFields(s string) []int {
	var result []int
	for _, is := range strings.Fields(s) {
		i, _ := strconv.Atoi(is)
		result = append(result, i)
	}
	return result
}

type Number struct {
	value  int
	marked bool
}

type Board struct {
	numbers  [5][5]Number
	finished bool
}

func (b *Board) mark(value int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.numbers[i][j].value == value {
				b.numbers[i][j].marked = true
			}
		}
	}
}

func (b *Board) done() bool {
	for row := 0; row < 5; row++ {
		win := true
		for col := 0; col < 5; col++ {
			if !b.numbers[row][col].marked {
				win = false
				break
			}
		}
		if win {
			b.finished = true
			return true
		}
	}
	for col := 0; col < 5; col++ {
		win := true
		for row := 0; row < 5; row++ {
			if !b.numbers[row][col].marked {
				win = false
				break
			}
		}
		if win {
			b.finished = true
			return true
		}
	}
	return false
}

func (b *Board) sumUnmarked() int {
	result := 0
	for col := 0; col < 5; col++ {
		for row := 0; row < 5; row++ {
			if !b.numbers[row][col].marked {
				result += b.numbers[row][col].value
			}
		}
	}
	return result
}

type Game struct {
	boards []Board
	order  []int
}

func parseGame(r io.Reader) Game {
	game := Game{}
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	orderString := scanner.Text()
	game.order = splitAndParse(orderString, ",")

	for scanner.Scan() {
		board := Board{}
		for row := 0; row < 5; row++ {
			scanner.Scan()
			rowString := scanner.Text()
			for i, val := range splitAndParseFields(rowString) {
				board.numbers[row][i] = Number{val, false}
			}
		}
		game.boards = append(game.boards, board)
	}
	return game
}

func (g *Game) numFinished() int {
	result := 0
	for bi, _ := range g.boards {
		if g.boards[bi].finished {
			result += 1
		}
	}
	return result
}

func main() {
	game := parseGame(os.Stdin)

	p1done, p2done := false, false
	for _, move := range game.order {
		if p1done && p2done {
			break
		}

		for board_i, _ := range game.boards {
			if game.boards[board_i].finished {
				continue
			}

			game.boards[board_i].mark(move)
			if game.boards[board_i].done() {
				if !p1done {
					fmt.Println("Part one ", move*game.boards[board_i].sumUnmarked())
					p1done = true
				}
				if game.numFinished() == len(game.boards) && !p2done {
					fmt.Println("Part two ", move*game.boards[board_i].sumUnmarked())
					p2done = true
				}
			}
		}
	}
}
