package main

import (
	"fmt"
	"strconv"
)

type Intcode []int

func run(intcode Intcode, inputs []int) []int {
	output := []int{}
	immediate := 1
	instruction := 0
	pos := 0

	trial := make(Intcode, len(intcode))
	for i := 0; i < len(intcode); i++ {
		trial[i] = intcode[i]
	}

	for pos < len(trial) {
		mode1 := 0
		mode2 := 0
		oppcode := trial[pos]
		if oppcode > 99 {
			s := strconv.Itoa(oppcode)
			oppcode, _ = strconv.Atoi(s[len(s)-2:])
			mode1, _ = strconv.Atoi(string(s[len(s)-3]))
			if len(s) > 3 {
				mode2, _ = strconv.Atoi(string(s[len(s)-4]))
			}
		}

		// parameters, modes
		var p1 int
		var p2 int
		if pos+1 < len(trial) {
			p1 = trial[pos+1]
			if mode1 != immediate && p1 < len(trial) {
				p1 = trial[p1]
			}
		}
		if pos+2 < len(trial) {
			p2 = trial[pos+2]
			if mode2 != immediate && p2 < len(trial) {
				p2 = trial[p2]
			}
		}

		switch oppcode {
		case 1: // add (3)
			trial[trial[pos+3]] = p1 + p2
			pos += 3 + 1 // move forward 3 parameters + 1 oppcode position
			break
		case 2: // multiply (3)
			trial[trial[pos+3]] = p1 * p2
			pos += 3 + 1
			break
		case 3: // input (1)
			if instruction >= len(inputs) {
				return output
			}
			trial[trial[pos+1]] = inputs[instruction]
			instruction++
			pos += 1 + 1 // move forward 1 parameter + 1 oppcode position, etc.
			break
		case 4: // output (1)
			output = append(output, p1)
			pos += 1 + 1
			break
		case 5: // jump-if-true (2)
			if p1 != 0 {
				pos = p2
			} else {
				pos += 2 + 1
			}
			break
		case 6: // jump-if-false (2)
			if p1 == 0 {
				pos = p2
			} else {
				pos += 2 + 1
			}
			break
		case 7: // less-than (3)
			if p1 < p2 {
				trial[trial[pos+3]] = 1
			} else {
				trial[trial[pos+3]] = 0
			}
			pos += 3 + 1
			break
		case 8: // equals (3)
			if p1 == p2 {
				trial[trial[pos+3]] = 1
			} else {
				trial[trial[pos+3]] = 0
			}
			pos += 3 + 1
			break
		case 99: // exit
			return output
		default:
			panic("help I’m dying")
		}
	}

	return output
}

func thrusters(phases [5]int, intcode Intcode, feedbackLoop bool) []int {
	i := 0
	output := []int{0, phases[0]}
	for {
		input := []int{phases[i%5]}
		input = append(input, output...)
		output = run(intcode, input)
		if (feedbackLoop == false && i == 4) || (feedbackLoop == true && output[0] == 99){
			break
		}
		i++
	}
	return output
}

func iterations(start int) [][5]int {
	var iterations [][5]int

	for a := start; a <= start+4; a++ {
		for b := start; b <= start+4; b++ {
			for c := start; c <= start+4; c++ {
				for d := start; d <= start+4; d++ {
					for e := start; e <= start+4; e++ {
						if a != b && a != c && a != d && a != e && b != c && b != d && b != e && c != d && c != e && d != e {
							iterations = append(iterations, [5]int{a, b, c, d, e})
						}
					}
				}
			}
		}
	}

	return iterations
}

func main() {
	intcode := Intcode{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 42, 51, 76, 93, 110, 191, 272, 353, 434, 99999, 3, 9, 1002, 9, 2, 9, 1001, 9, 3, 9, 1002, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 101, 5, 9, 9, 1002, 9, 3, 9, 1001, 9, 4, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 101, 3, 9, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 101, 5, 9, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}

	// part 1
	max := 0
	phaseSettings := iterations(0)
	for _, phase := range phaseSettings {
		output := thrusters(phase, intcode, false)
		if output[0] > max {
			max = output[0]
		}
	}

	fmt.Printf("Day 07, Part 1: %v", max)
	fmt.Println()

	// part 2
	nextMax := 0
	nextPhaseSettings := iterations(5)
	for _, phase := range nextPhaseSettings {
		output := thrusters(phase, intcode, true)
		if output[1] > nextMax {
			nextMax = output[1]
		}
	}

	fmt.Printf("Day 07, Part 2: %v", nextMax)
	fmt.Println()
}
