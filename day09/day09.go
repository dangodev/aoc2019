package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Intcode []int

const POSITION = 0
const IMMEDIATE = 1
const RELATIVE = 2

const MULTIPLY = 1
const ADD = 2
const INPUT = 3
const OUTPUT = 4
const IF_TRUE = 5
const IF_FALSE = 6
const LT = 7
const EQ = 8
const ADJUST_RB = 9
const EXIT = 99

func command(instruction int) (int, []int) {
	paramCount := map[int]int{
		MULTIPLY:  3,
		ADD:       3,
		INPUT:     1,
		OUTPUT:    1,
		IF_TRUE:   2,
		IF_FALSE:  2,
		LT:        3,
		EQ:        3,
		ADJUST_RB: 1,
		EXIT:      0,
	}

	// if <= 99, return number
	str := strconv.Itoa(instruction)
	modes := make([]int, paramCount[instruction])
	if len(str) < 3 {
		return instruction, modes
	}

	// otherwise, parse modes
	method, _ := strconv.Atoi(str[len(str)-2:])
	modes = make([]int, paramCount[method])
	params := []int{}
	for i := len(str) - 3; i >= 0; i-- {
		mode, _ := strconv.Atoi(string(str[i]))
		params = append(params, mode)
	}
	for i, v := range params {
		modes[i] = v
	}
	return method, modes
}

func multiply(x int, y int) int {
	return x * y
}

func add(x int, y int) int {
	return x + y
}

func ifTrue(x int) bool {
	return x != 0
}

func ifFalse(x int) bool {
	return x == 0
}

func lessThan(x int, y int) bool {
	return x < y
}

func eq(x int, y int) bool {
	return x == y
}

func loop(x int, length int) int {
	return int(math.Mod(float64(x), float64(length)))
}

func run(intcode Intcode, pos int, input []int, inputCount int, output []int, rb int) []int {
	method, modes := command(intcode[pos])
	params := make([]int, len(modes))
	firstParam := pos + 1
	for i, mode := range modes {
		next := firstParam + i
		switch mode {
		case IMMEDIATE:
			params[i] = intcode[next]
		case RELATIVE:
			params[i] = intcode[loop(rb+next, len(intcode))]
		case POSITION:
		default:
			params[i] = intcode[loop(intcode[next], len(intcode))]
		}
	}

	switch method {
	case MULTIPLY:
		intcode[params[2]] = multiply(params[0], params[1])
	case ADD:
		intcode[params[2]] = add(params[0], params[1])
	case INPUT:
		intcode[params[0]] = input[inputCount]
		inputCount++
	case OUTPUT:
		output = append(output, params[0])
	case IF_TRUE:
		if ifTrue(params[0]) {
			run(intcode, pos+params[1], input, inputCount, output, rb)
		}
	case IF_FALSE:
		if ifFalse(params[0]) {
			run(intcode, pos+params[1], input, inputCount, output, rb)
		}
	case LT:
		if lessThan(params[0], params[1]) {
			intcode[params[2]] = 1
		} else {
			intcode[params[2]] = 0
		}
	case EQ:
		if eq(params[0], params[1]) {
			intcode[params[2]] = 1
		} else {
			intcode[params[2]] = 0
		}
	case ADJUST_RB:
		rb += params[0]
	case EXIT:
		return output
	}

	next := firstParam + len(modes)
	if next < len(intcode)-1 {
		return run(intcode, next, input, inputCount, output, rb)
	}
	return output
}

func readIntcode(filename string) Intcode {
	f, _ := os.Open(filename)
	defer f.Close()

	intcode := Intcode{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		codes := strings.Split(scanner.Text(), ",")
		for _, v := range codes {
			c, _ := strconv.Atoi(v)
			intcode = append(intcode, c)
		}
	}
	return intcode
}

func main() {
	program := readIntcode("input.txt")
	pos := 0
	input := []int{}
	inputCount := 0
	output := []int{}
	rb := 0
	boost := run(program, pos, input, inputCount, output, rb)

	fmt.Printf("Day 09, Part 1: %v", boost)
}
