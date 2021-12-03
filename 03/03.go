package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mostCommonBitInPosition(report []int, pos int) int {
	acc := 0
	for _, reading := range report {
		acc += (reading >> pos) & 1
	}
	if 2*acc >= len(report) {
		return 1
	} else {
		return 0
	}
}

func mostCommonBitInPositionMasked(report []int, pos int, excluded []bool) int {
	acc := 0
	numIncluded := 0
	for i, reading := range report {
		if !excluded[i] {
			acc += (reading >> pos) & 1
			numIncluded += 1
		}
	}
	if 2*acc >= numIncluded {
		return 1
	} else {
		return 0
	}
}

func numFalse(l []bool) (int, int) {
	// the number of false values, plus the index of the last false
	acc := 0
	last := -1
	for i, b := range l {
		if !b {
			last = i
			acc += 1
		}
	}
	return acc, last
}

func main() {
	var report []int
	var length int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		length = len(line)
		val, _ := strconv.ParseInt(line, 2, 0)
		report = append(report, int(val))
	}

	// Part 1
	gamma := 0
	epsilon := 0
	for i := 0; i < length; i++ {
		if mostCommonBitInPosition(report, i) == 1 {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}
	fmt.Println(gamma * epsilon)

	// Part 2
	oxExcluded, co2Excluded := make([]bool, len(report)), make([]bool, len(report))
	ox, co2 := 0, 0
	for i := length - 1; i >= 0; i-- {
		oxCommon := mostCommonBitInPositionMasked(report, i, oxExcluded)
		co2Common := mostCommonBitInPositionMasked(report, i, co2Excluded)
		for j, reading := range report {
			if (reading >> i & 1) != oxCommon { // bit in position i does not match the most common
				oxExcluded[j] = true
			}
			if (reading >> i & 1) == co2Common {
				co2Excluded[j] = true
			}
		}
		oxRemaining, oxLastIndex := numFalse(oxExcluded)
		if oxRemaining == 1 {
			ox = report[oxLastIndex]
		}
		co2Remaining, co2LastIndex := numFalse(co2Excluded)
		if co2Remaining == 1 {
			co2 = report[co2LastIndex]
		}
	}
	fmt.Println(ox * co2)
}
