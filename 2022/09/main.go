package main

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getFilename() string {
	filename := "input.txt"
	if len(os.Args[1:]) == 1 {
		filename = os.Args[1]
	}

	return filename
}

func getInput(filename string) []string {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")

	return lines
}

type Position struct {
	x, y int
}

func (p *Position) toString() string {
	return string(p.x) + "-" + string(p.y)
}

func uniqPositions(positions []Position) map[string]bool {
	uniqPositions := make(map[string]bool)
	for _, p := range positions {
		uniqPositions[p.toString()] = true
	}
	return uniqPositions
}

func simulateHead(input []string) []Position {
	//start from 0,0
	positions := make([]Position, 0)
	curPosition := Position{0, 0}
	positions = append(positions, curPosition)
	for _, line := range input {
		split := strings.Fields(line)
		direction := split[0]
		count, _ := strconv.Atoi(split[1])

		for i := 0; i < count; i++ {
			if direction == "U" {
				curPosition = Position{curPosition.x, curPosition.y + 1}
			}
			if direction == "R" {
				curPosition = Position{curPosition.x + 1, curPosition.y}
			}
			if direction == "D" {
				curPosition = Position{curPosition.x, curPosition.y - 1}
			}
			if direction == "L" {
				curPosition = Position{curPosition.x - 1, curPosition.y}
			}
			positions = append(positions, curPosition)
		}

	}

	return positions
}

func sameAxisOppositeDirection(d1, d2 string) bool {
	if (d1 == "U" && d2 == "D") || (d1 == "D" && d2 == "U") {
		return true
	}

	if (d1 == "R" && d2 == "L") || (d1 == "L" && d2 == "R") {
		return true
	}

	return false
}

func differentAxisSameDirection(d1, d2 string) bool {
	// same direction means "increasing" the coordinate, i.e. up or right

	if (d1 == "U" && d2 == "R") || (d1 == "R" && d2 == "U") {
		return true
	}

	if (d1 == "L" && d2 == "D") || (d1 == "D" && d2 == "L") {
		return true
	}

	return false
}

func differentAxisOppositeDirection(d1, d2 string) bool {
	// opposite direction means from increasing to decreasing or vice versa

	if (d1 == "U" && d2 == "L") || (d1 == "L" && d2 == "U") {
		return true
	}

	if (d1 == "D" && d2 == "R") || (d1 == "R" && d2 == "D") {
		return true
	}

	return false
}

func positionsAreFarAway(p1, p2 Position, farAway int) bool {
	if int(math.Abs(float64(p2.x-p1.x))) >= farAway {
		return true
	}

	if int(math.Abs(float64(p2.y-p1.y))) >= farAway {
		return true
	}

	return false
}

func simulateTail(input []string) []Position {
	//start from 0,0
	positions := make([]Position, 0)
	curPosition := Position{0, 0}
	positions = append(positions, curPosition)
	curDirection := ""

	for _, line := range input {
		split := strings.Fields(line)
		nextDirection := split[0]
		count, _ := strconv.Atoi(split[1])

		if sameAxisOppositeDirection(curDirection, nextDirection) {
			// reduce one for overlapping and another for going to the opposite
			// direction
			count -= 2
		} else if differentAxisSameDirection(curDirection, nextDirection) {
			// ignore the first move because we're already remain "in contact"
			// (diagonally)
			count -= 1
		} else if differentAxisOppositeDirection(curDirection, nextDirection) {
			count -= 3
		}

		for i := 0; i < count; i++ {
			if nextDirection == "U" {
				curPosition = Position{curPosition.x, curPosition.y + 1}
			}
			if nextDirection == "R" {
				curPosition = Position{curPosition.x + 1, curPosition.y}
			}
			if nextDirection == "D" {
				curPosition = Position{curPosition.x, curPosition.y - 1}
			}
			if nextDirection == "L" {
				curPosition = Position{curPosition.x - 1, curPosition.y}
			}
			positions = append(positions, curPosition)
			curDirection = nextDirection
		}

	}

	return positions
}

func main() {
	input := getInput(getFilename())

	// get the positions the head will move
	headPositions := simulateHead(input)
	uniqHeadPositions := uniqPositions(headPositions)
	log.Print(len(headPositions))
	log.Print(len(uniqHeadPositions))

	// the tail does 1 fewer move than the head
	// e.g. if the head goes right 4 positions, the tail only goes 3
	//
	// this is also true if the head changes direction vertically
	// e.g. the head goes up/down 1 position, the tail is still close enough
	//
	// if the head moves in the same axies, but to the opposite direction then
	// the tail follows if the head moves more than 2 moves:
	// the first will overlap with tail and the second will still be near enough
	tailPositions := simulateTail(input)
	uniqTailPositions := uniqPositions(tailPositions)
	log.Print(len(tailPositions))
	log.Print(len(uniqTailPositions))
}
