package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const UP = "UP"
const DOWN = "DOWN"
const LEFT = "LEFT"
const RIGHT = "RIGHT"

type whileTrue func(char string) bool

type Coord struct {
	x float64
	y float64
}

type Paths map[Coord][]Coord

type Maze [][]string

type Solution map[string]bool

type Warps map[string][]Coord

func getSurrounding(pos Coord, maze Maze) map[string]string {
	x := pos.x
	y := pos.y
	move := map[string]string{}
	if pos.x > 0 {
		move[LEFT] = readTile(Coord{x: x - 1, y: y}, maze)
	}
	if int(pos.x) < len(maze[0])-1 {
		move[RIGHT] = readTile(Coord{x: x + 1, y: y}, maze)
	}
	if pos.y > 0 {
		move[UP] = readTile(Coord{x: x, y: y - 1}, maze)
	}
	if int(pos.y) < len(maze)-1 {
		move[DOWN] = readTile(Coord{x: x, y: y + 1}, maze)
	}
	return move
}

func readTile(pos Coord, maze Maze) string {
	return maze[int(pos.y)][int(pos.x)]
}

func move(start Coord, dir string) Coord {
	newPos := Coord{x: start.x, y: start.y}
	switch dir {
	case LEFT:
		newPos = Coord{x: start.x - 1, y: start.y}
	case RIGHT:
		newPos = Coord{x: start.x + 1, y: start.y}
	case UP:
		newPos = Coord{x: start.x, y: start.y - 1}
	case DOWN:
		newPos = Coord{x: start.x, y: start.y + 1}
	}
	return newPos
}

func isLetter(input string) bool {
	char := regexp.MustCompile(`[A-Z]`)
	return char.MatchString(input)
}

func warpName(start Coord, maze Maze) string {
	var dir string
	surrounding := getSurrounding(start, maze)
	// two warps won’t ever share a letter, so it’s safe to assume any immediate
	// letters belong to the same warp name.
	// the first adjacent tile reveals orientation of name
	if t, ok := surrounding[UP]; ok && isLetter(t) {
		dir = UP
	} else if t, ok := surrounding[DOWN]; ok && isLetter(t) {
		dir = DOWN
	} else if t, ok := surrounding[LEFT]; ok && isLetter(t) {
		dir = LEFT
	} else if t, ok := surrounding[RIGHT]; ok && isLetter(t) {
		dir = RIGHT
	} else {
		return ""
	}

	pos := Coord{x: start.x, y: start.y}
	name := ""
	for len(name) < 2 {
		c := readTile(pos, maze)
		if isLetter(c) {
			if dir == UP || dir == LEFT {
				name = c + name
			} else {
				name += c
			}
		}
		pos = move(pos, dir)
	}

	return name
}

func readMaze(filename string) Maze {
	f, _ := os.Open(filename)
	defer f.Close()

	var maze Maze
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		maze = append(maze, strings.Split(scanner.Text(), ""))
	}
	return maze
}

func warpMap(maze Maze) Warps {
	warps := Warps{}
	for y, row := range maze {
		for x, col := range row {
			if col == "." {
				pos := Coord{x: float64(x), y: float64(y)}
				warp := warpName(pos, maze)
				if len(warp) == 2 {
					if _, ok := warps[warp]; ok {
						// note: this doesn’t check for duplicates, but since maze is only
						// read once it shouldn’t be a concern
						warps[warp] = append(warps[warp], pos)
					} else {
						warps[warp] = []Coord{pos}
					}
				}
			}
		}
	}

	return warps
}

func coordToString(c Coord) string {
	return strconv.Itoa(int(c.x)) + "," + strconv.Itoa(int(c.y))
}

func strToCoord(s string) Coord {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Coord{x: float64(x), y: float64(y)}
}

func alreadyScanned(path string, c Coord) bool {
	scanned := false
	parts := strings.Split(path, " ")
	test := coordToString(c)
	for _, v := range parts {
		if v == test {
			scanned = true
		}
	}
	return scanned
}

func trace(id string, end Coord, paths Solution, maze Maze, warps Warps) Solution {
	parts := strings.Split(id, " ")
	lastPart := parts[len(parts)-1]
	last := strToCoord(lastPart)

	// reached end 🎉
	if last.x == end.x && last.y == end.y {
		paths[id] = true
		return paths
	}

	adjacent := getSurrounding(last, maze)
	for dir, char := range adjacent {
		next := move(Coord{x: last.x, y: last.y}, dir)

		if isLetter(char) { // warp
			warp := warpName(last, maze)

			// ignore starting warp, and other non-warps
			if len(warps[warp]) < 2 {
				continue
			}

			// warps have 2 entries so figure out which to use
			if last.x == warps[warp][0].x && last.y == warps[warp][0].y {
				next = warps[warp][1]
			} else {
				next = warps[warp][0]
			}

			// ignore previously-scanned item
			if alreadyScanned(id, next) {
				continue
			}

			nextID := id + " " + coordToString(next)
			paths[nextID] = false // only add to solution if we can continue
			trace(nextID, end, paths, maze, warps)
		} else if char == "." { // single step
			// ignore previously-scanned item
			if alreadyScanned(id, next) {
				continue
			}

			nextID := id + " " + coordToString(next)
			paths[nextID] = false // only add to solution if we can continue
			trace(nextID, end, paths, maze, warps)
		}
	}

	return paths
}

func wayfind(filename string) [][]Coord {
	maze := readMaze(filename)
	warps := warpMap(maze)

	// get solutions
	end := warps["ZZ"][0]
	id := coordToString(warps["AA"][0])
	paths := Solution{}
	paths[id] = false
	trace(id, end, paths, maze, warps)

	// copy to slice
	routes := [][]Coord{}
	for k, v := range paths {
		if v {
			coords := strings.Split(k, " ")
			path := []Coord{}
			for _, c := range coords {
				path = append(path, strToCoord(c))
			}
			routes = append(routes, path)
		}
	}

	// sort by length
	sort.Slice(routes, func(i, j int) bool {
		return len(routes[i]) < len(routes[j])
	})

	return routes
}

func main() {
	possible := wayfind("input.txt")
	fmt.Printf("Day 20, Part 1: %v", len(possible[0])-1) // subtract 1 because starting x,y was included
	fmt.Println()
}
