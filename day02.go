package main

import (
	"fmt"
	"strconv"
)

type Intcode []int

func run(intcode Intcode) int {
	trial := intcode
Loop:
	for i := 0; i < len(intcode); i += 4 {
		switch trial[i] {
		// add
		case 1:
			trial[trial[i+3]] = trial[trial[i+1]] + trial[trial[i+2]]
		// multiply
		case 2:
			trial[trial[i+3]] = trial[trial[i+1]] * trial[trial[i+2]]
		// stop execution
		case 99:
			break Loop
		default:
			panic("unknown instruction “" + strconv.Itoa(trial[i]) + "”")
		}
	}
	return trial[0]
}

func main() {
	intcode := Intcode{
		1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 19, 5, 23, 1, 23, 6, 27, 2, 9, 27, 31, 1, 5, 31, 35, 1, 35, 10, 39, 1, 39, 10, 43, 2, 43, 9, 47, 1, 6, 47, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 59, 10, 63, 1, 9, 63, 67, 1, 9, 67, 71, 2, 71, 6, 75, 1, 5, 75, 79, 1, 5, 79, 83, 1, 9, 83, 87, 2, 87, 10, 91, 2, 10, 91, 95, 1, 95, 9, 99, 2, 99, 9, 103, 2, 10, 103, 107, 2, 9, 107, 111, 1, 111, 5, 115, 1, 115, 2, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0,
	}

	// part 1
	run1 := make([]int, len(intcode))
	for i, v := range intcode {
		run1[i] = v
	}

	// alter intcode before starting as per instructions
	run1[1] = 12
	run1[2] = 2

	fmt.Printf("Part 1: %v", run(run1))
	fmt.Println()

	// part 2
	magicNumber := 19690720

Loop:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			trial := make([]int, len(intcode))
			for i, v := range intcode {
				trial[i] = v
			}
			trial[1] = noun
			trial[2] = verb

			result := run(trial)

			if result == magicNumber {
				fmt.Printf("Part 2: %v", 100*noun+verb)
				fmt.Println()
				break Loop
			}
		}
	}
}

/* NOTES (not in the README b/c spoilers)

Things I learned
- Got bitten by array pointer assignment—same as JavaScript. Figured out the for-loop array copy the hard way.
- Learned about labels (Loop:)! That’s so much nicer than having to do something hacky to break out of nested loops.

Questions
- Is there a more efficient way to reset the intcode on every trial?
*/
