package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x float64
	y float64
	z float64
}

type moon struct {
	pos      coord
	velocity coord
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func moonPos(m moon) string {
	str := strconv.FormatFloat(m.pos.x, 'f', -1, 64)
	str += ","
	str += strconv.FormatFloat(m.pos.y, 'f', -1, 64)
	str += ","
	str += strconv.FormatFloat(m.pos.z, 'f', -1, 64)
	str += "/"
	str += strconv.FormatFloat(m.velocity.x, 'f', -1, 64)
	str += ","
	str += strconv.FormatFloat(m.velocity.y, 'f', -1, 64)
	str += ","
	str += strconv.FormatFloat(m.velocity.z, 'f', -1, 64)
	return str
}

func universeToString(moons []moon) string {
	var strs []string
	for _, m := range moons {
		strs = append(strs, moonPos(m))
	}
	return strings.Join(strs, " ")
}

func pairs(max int) [][2]int {
	var combinations [][2]int
	for a := 0; a < max-1; a++ {
		for b := a + 1; b < max; b++ {
			combinations = append(combinations, [2]int{a, b})
		}
	}
	return combinations
}

func step(moonInput []moon) []moon {
	moons := make([]moon, len(moonInput))
	for n, v := range moonInput {
		moons[n] = v
	}

	for _, pair := range pairs(len(moons)) {
		a := pair[0]
		b := pair[1]

		// apply x, y, z
		deltaX := float64(0)
		deltaY := float64(0)
		deltaZ := float64(0)
		if moons[a].pos.x > moons[b].pos.x {
			deltaX = -1
		} else if moons[a].pos.x < moons[b].pos.x {
			deltaX = 1
		}
		if moons[a].pos.y > moons[b].pos.y {
			deltaY = -1
		} else if moons[a].pos.y < moons[b].pos.y {
			deltaY = 1
		}
		if moons[a].pos.z > moons[b].pos.z {
			deltaZ = -1
		} else if moons[a].pos.z < moons[b].pos.z {
			deltaZ = 1
		}
		moons[a].velocity.x += deltaX
		moons[b].velocity.x -= deltaX
		moons[a].velocity.y += deltaY
		moons[b].velocity.y -= deltaY
		moons[a].velocity.z += deltaZ
		moons[b].velocity.z -= deltaZ
	}

	for n := range moons {
		moons[n].pos.x += moons[n].velocity.x
		moons[n].pos.y += moons[n].velocity.y
		moons[n].pos.z += moons[n].velocity.z
	}

	return moons
}

func energy(m moon) float64 {
	return (math.Abs(m.pos.x) + math.Abs(m.pos.y) + math.Abs(m.pos.z)) * (math.Abs(m.velocity.x) + math.Abs(m.velocity.y) + math.Abs(m.velocity.z))
}

func posToString(moons []moon) (string, string, string) {
	xPos := []string{}
	xVel := []string{}
	yPos := []string{}
	yVel := []string{}
	zPos := []string{}
	zVel := []string{}
	for _, v := range moons {
		xPos = append(xPos, strconv.FormatFloat(v.pos.x, 'f', -1, 64))
		xVel = append(xVel, strconv.FormatFloat(v.velocity.x, 'f', -1, 64))
		yPos = append(yPos, strconv.FormatFloat(v.pos.y, 'f', -1, 64))
		yVel = append(yVel, strconv.FormatFloat(v.velocity.y, 'f', -1, 64))
		zPos = append(zPos, strconv.FormatFloat(v.pos.z, 'f', -1, 64))
		zVel = append(zVel, strconv.FormatFloat(v.velocity.z, 'f', -1, 64))
	}
	x := strings.Join(xPos, ",")
	x += "/"
	x += strings.Join(xVel, ",")
	y := strings.Join(yPos, ",")
	y += "/"
	y += strings.Join(yVel, ",")
	z := strings.Join(zPos, ",")
	z += "/"
	z += strings.Join(zVel, ",")
	return x, y, z
}

func calcRepeat(moons []moon) int {
	next := moons

	xPeriod := 0
	xHistory := map[string]int{}
	yPeriod := 0
	yHistory := map[string]int{}
	zPeriod := 0
	zHistory := map[string]int{}

	n := 0
	for xPeriod == 0 || yPeriod == 0 || zPeriod == 0 {
		x, y, z := posToString(next)
		if xPeriod == 0 {
			if _, ok := xHistory[x]; ok {
				xPeriod = n
			} else {
				xHistory[x] = n
			}
		}
		if yPeriod == 0 {
			if _, ok := yHistory[y]; ok {
				yPeriod = n
			} else {
				yHistory[y] = n
			}
		}
		if zPeriod == 0 {
			if _, ok := zHistory[z]; ok {
				zPeriod = n
			} else {
				zHistory[z] = n
			}
		}
		next = step(next)
		n++
	}

	return LCM(xPeriod, yPeriod, zPeriod)
}

func main() {
	input := []moon{
		{
			pos:      coord{x: -3, y: 15, z: -11},
			velocity: coord{x: 0, y: 0, z: 0},
		},
		{
			pos:      coord{x: 3, y: 13, z: -19},
			velocity: coord{x: 0, y: 0, z: 0},
		},
		{
			pos:      coord{x: -13, y: 18, z: -2},
			velocity: coord{x: 0, y: 0, z: 0},
		},
		{
			pos:      coord{x: 6, y: 0, z: -1},
			velocity: coord{x: 0, y: 0, z: 0},
		},
	}

	nextStep := input
	for n := 0; n < 1000; n++ {
		nextStep = step(nextStep)
	}

	totalEnergy := float64(0)
	for _, m := range nextStep {
		totalEnergy += energy(m)
	}
	fmt.Printf("Day 12, Part 1: %v", totalEnergy)
	fmt.Println()

	fmt.Printf("Day 12, Part 2: %v", calcRepeat(input))
	fmt.Println()
}
