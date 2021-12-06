package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fish [9]int

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
		spawns := fish[0]
		for i := 0; i < 8; i++ {
			fish[i] = fish[i+1]
		}
		fish[8] = spawns
		fish[6] += spawns
	}

	total := 0
	for _, v := range fish {
		total += v
	}
	fmt.Println(total)
}
