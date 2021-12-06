package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fish := make(map[int]int)
	for i := 0; i < 9; i++ {
		fish[i] = 0
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fishStrings := strings.Split(line, ",")
	for _, s := range fishStrings {
		i, _ := strconv.Atoi(s)
		fish[i] += 1
	}

	days := 256
	for day := 0; day < days; day++ {
		previous := make(map[int]int)
		for i := 0; i < 9; i++ {
			previous[i] = fish[i]
		}

		for i := 0; i < 8; i++ {
			fish[i] = previous[i+1]
		}
		fish[8] = previous[0]
		fish[6] += previous[0]
	}

	total := 0
	for _, v := range fish {
		total += v
	}
	fmt.Println(total)
}
