package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func getInput() []string {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(contents)), "\n")
}

func slopeCalculator(input []string, down, right int) int {
	treeCount := 0
	tree := '#'

	for i := down; i < len(input); i += down {
		// when we reach end of line, then we start over, as the
		// map repeats itself
		j := 0
		if down <= right {
			j = (i * right) % len(input[i])
		} else {
			j = (i / down) % len(input[i])
		}

		lineRune := []rune(input[i])

		if lineRune[j] == tree {
			treeCount++
		}
	}

	log.Print(right, down, ": ", treeCount)
	return treeCount
}

func main() {
	input := getInput()

	product := 1

	// positions: 1-1, 2-2, 3-3, 4-4, etc
	product *= slopeCalculator(input, 1, 1)

	// positions: 1-3, 2-6, 3-9, 4-12, etc
	product *= slopeCalculator(input, 1, 3)

	// positions: 1-5, 2-10, 3-15, 4-20, etc
	product *= slopeCalculator(input, 1, 5)

	// positions: 1-7, 2-14, 3-21, 4-28, etc
	product *= slopeCalculator(input, 1, 7)

	// positions: 2-1, 4-2, 6-3, 8-4, etc
	product *= slopeCalculator(input, 2, 1)

	log.Print(product)
}
