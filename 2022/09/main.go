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
	xString := strconv.Itoa(p.x)
	yString := strconv.Itoa(p.y)
	return xString + "-" + yString
}

func uniqPositions(positions []Position) map[string]bool {
	uniqPositions := make(map[string]bool)
	for _, p := range positions {
		uniqPositions[p.toString()] = true
	}
	return uniqPositions
}

func runMoves(input []string) ([]Position, []Position) {
	//start from 0,0
	curPosition := Position{0, 0}

	headPositions := make([]Position, 0)
	tailPositions := make([]Position, 0)
	headPositions = append(headPositions, curPosition)
	tailPositions = append(tailPositions, curPosition)

	direction := ""
	for _, line := range input {
		split := strings.Fields(line)
		direction = split[0]
		count, _ := strconv.Atoi(split[1])

		for i := 0; i < int(math.Abs(float64(count))); i++ {
			newPosition := Position{}
			if direction == "U" {
				newPosition = Position{curPosition.x, curPosition.y + 1}
			}
			if direction == "R" {
				newPosition = Position{curPosition.x + 1, curPosition.y}
			}
			if direction == "D" {
				newPosition = Position{curPosition.x, curPosition.y - 1}
			}
			if direction == "L" {
				newPosition = Position{curPosition.x - 1, curPosition.y}
			}
			log.Print("Moving from ", curPosition, " to ", newPosition)
			headPositions = append(headPositions, newPosition)

			// if the tail has moved too far away then append the previous position of the head
			// for example:
			// t(0,0), h(0,0) -> t(0,0), h(1, 0) -> t(0,0), h(2,0)
			// in the third moment we add the previous position of the head to
			// the tail positions, (1,0)
			//
			// if head was to move right 1, up 2, from the same starting position (0,0)
			// t(0,0), h(1,0) -> t(0,0), h(1,1) -> t(0,0), h(1,2)
			// again at this point we add the previous position of the head
			// (1,1) to the tail's positions
			lastTailPosition := tailPositions[len(tailPositions)-1]
			log.Print("Last tail position: ", lastTailPosition)
			if positionsAreFarAway(lastTailPosition, newPosition, 2) {
				secondTolastHeadPosition := headPositions[len(headPositions)-2]
				log.Print("==> Tail is far away, moving to ", secondTolastHeadPosition)
				tailPositions = append(tailPositions, secondTolastHeadPosition)
			}
			curPosition = newPosition
		}

	}

	return headPositions, tailPositions
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

func main() {
	input := getInput(getFilename())

	// get the positions the head will move
	headPositions, tailPositions := runMoves(input)
	uniqHeadPositions := uniqPositions(headPositions)
	uniqTailPositions := uniqPositions(tailPositions)
	log.Print(len(headPositions))
	log.Print(len(uniqHeadPositions))
	log.Print(len(tailPositions))
	log.Print(len(uniqTailPositions))
}
