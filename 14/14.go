package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const NUM_ITERATIONS = 40

type ElementCounts [26]int

var cache [NUM_ITERATIONS + 1]map[string]ElementCounts
var rules map[string]byte

func add_counts(dst, src *ElementCounts) {
	for i, v := range src {
		dst[i] += v
	}
}

func largest_minus_smallest(counts ElementCounts) int {
	fmt.Println(counts)
	largest, smallest := math.MinInt64, math.MaxInt64
	for _, count := range counts {
		if count == 0 { // Not all letters appear
			continue
		}
		if count > largest {
			largest = count
		}
		if count < smallest {
			smallest = count
		}
	}
	return largest - smallest
}

// Dynamic programming type thing
// A pair splits to add a middle element and create two new pairs.
// Bump the count of the middle element and recurse to add the final counts produced
// by each of the new pairs.
func counts_for_pair(pair string, iterations int) ElementCounts {
	var result ElementCounts

	if iterations == 0 {
		return result
	}
	middle, exists := rules[pair]
	if !exists { // No mapping for this pair
		return result
	}

	if cache[iterations] == nil {
		cache[iterations] = make(map[string]ElementCounts)
	}
	value, exists := cache[iterations][pair]
	if exists { // Cache hit - we already know the result for this pair and number of iterations
		return value
	}

	result[int(middle-'A')] += 1
	left_pair := string([]byte{pair[0], middle})
	right_pair := string([]byte{middle, pair[1]})
	left_counts := counts_for_pair(left_pair, iterations-1)
	right_counts := counts_for_pair(right_pair, iterations-1)
	add_counts(&result, &left_counts)
	add_counts(&result, &right_counts)

	cache[iterations][pair] = result
	return result
}

func counts_for_polymer(polymer string, iterations int) ElementCounts {
	var result ElementCounts
	for i := 0; i < len(polymer); i++ {
		result[int(polymer[i]-'A')] += 1
	}
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		pair_counts := counts_for_pair(pair, iterations)
		add_counts(&result, &pair_counts)
	}
	return result
}

func main() {
	rules = make(map[string]byte)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	polymer := line
	scanner.Scan()
	for scanner.Scan() {
		line = scanner.Text()
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1][0]
	}

	part1 := counts_for_polymer(polymer, 10)
	fmt.Println(largest_minus_smallest(part1))

	part2 := counts_for_polymer(polymer, 40)
	fmt.Println(largest_minus_smallest(part2))
}
