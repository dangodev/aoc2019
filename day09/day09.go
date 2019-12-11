package main

import "strconv"

type Intcode []int

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

func run(intcode Intcode, inputs []int) []int {
	immediate := 1
	relative := 2

	output := []int{}

	instruction := 0
	pos := 0
	relativeBase := 0

	trial := make(Intcode, len(intcode))
	for i := 0; i < len(intcode); i++ {
		trial[i] = intcode[i]
	}

	for pos < len(trial) {
		input := trial[pos]
		oppcode, modes := getModes(input)

		var p1 int
		var p2 int
		switch modes[0] {
		case immediate:
			p1 = trial[pos+1]
			break
		case relative:
			p1 = trial[pos+relativeBase]
			break
		default:
			p1 = trial[trial[pos+1]]
			break
		}
		switch modes[1] {
		case immediate:
			p2 = trial[pos+2]
			break
		case relative:
			p2 = trial[pos+relativeBase]
			break
		default:
			p2 = trial[trial[pos+2]]
			break
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
		case 9: // adjust relative base
			relativeBase += p1
			pos += 1 + 1
			break
		case 99: // exit
			return output
		default:
			panic("help I’m dying")
		}
	}

	return output
}

func main() {

}
