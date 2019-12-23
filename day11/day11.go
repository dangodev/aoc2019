package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func readIntcode(filename string) [][2]int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var commands [][2]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		instructions := strings.Split(line, ",")
		for n := 0; n < len(instructions); n += 2 {
			color, _ := strconv.Atoi(instructions[n])
			turn, _ := strconv.Atoi(instructions[n+1])
			commands = append(commands, [2]int{color, turn})
		}
	}
	return commands
}

func move(turn int, start coord, bearing float64) (coord, float64) {
	nextCoord := coord{x: start.x, y: start.y}
	nextBearing := bearing
	if turn == 0 {
		nextBearing += 90
	} else if turn == 1 {
		nextBearing -= 90
	}
	if nextBearing < 0 {
		nextBearing = 360 - nextBearing
	}
	nextBearing = math.Mod(nextBearing, 360)
	switch nextBearing {
	case 0:
		nextCoord.x++
		break
	case 90:
		nextCoord.y++
		break
	case 180:
		nextCoord.x--
		break
	case 270:
		nextCoord.y--
		break
	}
	return nextCoord, nextBearing
}

func main() {
	intcode := readIntcode("input.txt")
	grid := make(map[coord]int)

	pos := coord{x: 0, y: 0}
	bearing := float64(90)
	for _, v := range intcode {
		color := v[0]
		turn := v[1]
		grid[pos] = color
		pos, bearing = move(turn, pos, bearing)
	}

	fmt.Println(grid)

	fmt.Printf("Day 11, Part 1: %v", len(grid))
	fmt.Println()
}
