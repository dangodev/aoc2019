package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func validPass(input int, boundLower int, boundUpper int) bool {
	if input < boundLower || input > boundUpper {
		return false
	}

	// logic
	adjacentDigits := false // easier to validate
	alwaysIncrease := true  // easier to invalidate

	inputS := strconv.Itoa(input)
	for i, d := range inputS {
		if i == 0 {
			continue
		}

		current, _ := strconv.Atoi(string(d))
		prev, _ := strconv.Atoi(string(inputS[i-1]))
		if prev == current {
			adjacentDigits = true
		}
		if current < prev {
			alwaysIncrease = false
			break
		}
	}

	return adjacentDigits && alwaysIncrease
}

func main() {
	boundLower := 265275
	boundUpper := 781584

	// part 1
	validPasses := []int{}
	for i := boundLower; i <= boundUpper; i++ {
		if validPass(i, boundLower, boundUpper) {
			validPasses = append(validPasses, i)
		}
	}
	fmt.Printf("Day 04, Part 1: %v", len(validPasses))
	fmt.Println()

	// part 2
	pairsOnly := []int{}
	re := regexp.MustCompile(`00+|11+|22+|33+|44+|55+|66+|77+|88+|99+`) // this is the dumbest RegEx I’ve ever written
	for _, pass := range validPasses {
		matches := re.FindAllString(strconv.Itoa(pass), -1)
		validMatch := false
		for _, v := range matches {
			if len(v) == 2 {
				validMatch = true
				break // all we need is 1 pair for this to be valid
			}
		}
		if validMatch {
			pairsOnly = append(pairsOnly, pass)
		}
	}

	fmt.Printf("Day 04, Part 2: %v", len(pairsOnly))
	fmt.Println()
}

/* NOTES (not in the README b/c spoilers)

Things I learned
- RegEx is hard in Go

Questions
- Repeating character problems are surprisingly tricky
*/
