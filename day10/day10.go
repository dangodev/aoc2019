package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type coord struct {
	x float64
	y float64
}

type asteroid struct {
	angle    float64
	distance float64
	x        float64
	y        float64
}

type ByDistance []asteroid

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].distance < a[j].distance }

func makeAsteroidMap(origin coord, lines []string) map[float64][]asteroid {
	asteroidMap := make(map[float64][]asteroid)

	for y, row := range lines {
		for x, col := range row {
			if string(col) == "#" {
				relX := float64(x) - origin.x
				relY := float64(y) - origin.y
				angle := math.Atan2(relY, relX)
				if angle < 0 {
					angle = math.Pi*2 + angle
				}
				if angle == 0 && relX < 0 {
					angle = math.Pi
				}
				angle += math.Pi / 2
				if angle >= math.Pi*2 {
					angle = angle - math.Pi*2
				}
				newAsteroid := asteroid{
					angle:    angle,
					distance: math.Sqrt(math.Pow(relX, 2) + math.Pow(relY, 2)),
					x:        float64(x),
					y:        float64(y),
				}
				if _, ok := asteroidMap[angle]; ok {
					asteroidMap[angle] = append(asteroidMap[angle], newAsteroid)
				} else {
					asteroidMap[angle] = []asteroid{newAsteroid}
				}
			}
		}
	}

	for k := range asteroidMap {
		sort.Sort(ByDistance(asteroidMap[k]))
	}

	return asteroidMap
}

func blastingOrder(allAsteroids map[float64][]asteroid) []asteroid {
	maxLen := 1
	var angles []float64
	for k, v := range allAsteroids {
		if len(v) > maxLen {
			maxLen = len(v)
		}
		angles = append(angles, k)
	}
	sort.Float64s(angles)

	var ordered []asteroid
	for n := 0; n < maxLen; n++ {
		for _, v := range angles {
			if n < len(allAsteroids[v]) {
				ordered = append(ordered, allAsteroids[v][n])
			}
		}
	}

	return ordered
}

func readFile(filename string) []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func main() {
	var originMap map[float64][]asteroid
	visible := 0

	locations := readFile("input.txt")
	for y, row := range locations {
		for x, col := range row {
			if string(col) == "." {
				pos := coord{x: float64(x), y: float64(y)}
				asteroidMap := makeAsteroidMap(pos, locations)
				asteroidCount := len(asteroidMap)

				if asteroidCount > visible {
					originMap = asteroidMap
					visible = asteroidCount
				}
			}
		}
	}

	fmt.Printf("Day 10, Part 1: %v", visible)
	fmt.Println()

	ordered := blastingOrder(originMap)
	fmt.Printf("Day 10, Part 2: %v", ordered[199].x*100+ordered[199].y)
	fmt.Println()
}
