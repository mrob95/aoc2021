package main

import "fmt"

type Target struct {
	xmin, xmax, ymin, ymax int
}

func simulate(ux, uy int, t Target) (int, bool) {
	highest := 0
	x, y := 0, 0

	for x <= t.xmax && y >= t.ymin {
		x += ux
		y += uy

		if y > highest {
			highest = y
		}
		if x >= t.xmin && x <= t.xmax && y >= t.ymin && y <= t.ymax {
			return highest, true
		}

		if ux < 0 {
			ux += 1
		} else if ux > 0 {
			ux -= 1
		}
		uy -= 1
	}
	return 0, false
}

func main() {

	// target := Target{20, 30}
	target := Target{185, 221, -122, -74}

	highest := 0
	numHits := 0

	// initial velocity can't be less than ymin or greater than xmax or we would overshoot immediately
	// 10k seems to be big enough...
	for uy := target.ymin; uy < 10000; uy++ {
		for ux := 0; ux <= target.xmax; ux++ {
			high, hit := simulate(ux, uy, target)
			if hit {
				numHits++
			}
			if hit && high > highest {
				highest = high
			}
		}
	}
	fmt.Println(highest)
	fmt.Println(numHits)

}
