package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getModes(input int) (oppcode int, modes []int) {
	if input > 99 {
		s := strconv.Itoa(input)
		input, _ = strconv.Atoi(s[len(s)-2:])
		params := s[0 : len(s)-2]
		var paramModes []int
		for i := len(params) - 1; i >= 0; i-- {
			mode, _ := strconv.Atoi(params[i:i])
			paramModes = append(paramModes, mode)
		}
		return input, paramModes
	}
	return input, []int{}
}

func run(intcode []int, inputs []int) []int {
	immediate := 1
	relative := 2

	output := []int{}

	instruction := 0
	pos := 0
	relativeBase := 0

	paramCount := map[int]int{
		1:  3, // multiply
		2:  3, // add
		3:  1, // input
		4:  1, // output
		5:  2, // jump-if-true
		6:  2, // jump-if-false
		7:  3, // less-than
		8:  3, // equals
		9:  1, // relative base
		99: 0, // exit
	}

	trial := make([]int, len(intcode))
	for i := 0; i < len(intcode); i++ {
		trial[i] = intcode[i]
	}

	for pos < len(trial) {
		input := trial[pos]
		oppcode, modes := getModes(input)
		params := make([]int, paramCount[oppcode])
		for param := 0; param < paramCount[oppcode]; param++ {
			nextParam := param + 1
			if param < len(modes) {
				switch modes[param] {
				case immediate:
					params[param] = trial[pos+nextParam]
					break
				case relative:
					params[param] = trial[nextParam+relativeBase]
					break
				}
			}
			params[param] = trial[trial[pos+nextParam]]
		}

		switch oppcode {
		case 1: // add
			trial[trial[params[2]]] = params[0] + params[1]
		case 2: // multiply
			trial[trial[params[2]]] = params[0] * params[1]
		case 3: // input
			trial[trial[params[1]]] = inputs[instruction]
			instruction++
		case 4: // output
			output = append(output, params[0])
		case 5: // jump-if-true
			if params[0] != 0 {
				pos = params[1]
				break
			}
		case 6: // jump-if-false
			if params[0] == 0 {
				pos = params[1]
				break
			}
		case 7: // less-than
			if params[0] < params[1] {
				trial[trial[params[2]]] = 1
			} else {
				trial[trial[params[2]]] = 0
			}
		case 8: // equals
			if params[0] == params[1] {
				trial[trial[params[2]]] = 1
			} else {
				trial[trial[params[2]]] = 0
			}
		case 9: // adjust relative base
			relativeBase += params[0]
		case 99: // exit
			return output
		}
		pos += 1 + paramCount[oppcode]
	}

	return output
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var codes []string
	for scanner.Scan() {
		codes = strings.Split(scanner.Text(), ",")
	}
	program := make([]int, len(codes))
	for i, v := range codes {
		program[i], _ = strconv.Atoi(v)
	}

	// part 1
	boost := run(program, []int{1})
	fmt.Printf("Day 09, Part 1: %v", boost)
}
