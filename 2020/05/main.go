package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func getInput() []string {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")

	return lines
}

func recurseRows(directions []string, min, max int) int {
	//log.Print(directions, min, max)
	if len(directions) == 1 {
		if directions[0] == "F" {
			return min
		}
		return max
	}

	seatRange := max - min
	halfRange := seatRange / 2
	if directions[0] == "F" {
		return recurseRows(directions[1:], min, max-halfRange-1)
	}

	if directions[0] == "B" {
		return recurseRows(directions[1:], min+halfRange+1, max)
	}

	return -1
}

func recurseCols(directions []string, min, max int) int {
	//log.Print(directions, min, max)
	if len(directions) == 1 {
		if directions[0] == "L" {
			return min
		}
		return max
	}

	seatRange := max - min
	halfRange := seatRange / 2
	if directions[0] == "L" {
		return recurseCols(directions[1:], min, max-halfRange-1)
	}

	if directions[0] == "R" {
		return recurseCols(directions[1:], min+halfRange+1, max)
	}

	return -1
}

func getRow(v []string) int {
	return recurseRows(v, 0, 127)
}

func getCol(v []string) int {
	return recurseCols(v, 0, 7)
}

func main() {
	seats := getInput()

	allSeatIds := make([]int, len(seats))
	maxSeatId := 0
	for i, s := range seats {
		row := getRow(strings.Split(s[:7], ""))
		col := getCol(strings.Split(s[7:], ""))
		seatId := (row * 8) + col
		allSeatIds[i] = seatId

		// log.Print(row, col)
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	log.Print(maxSeatId)

	sort.Ints(allSeatIds)
	for i := 1; i < len(allSeatIds); i++ {
		currentSeatId := allSeatIds[i-1] + 1
		if currentSeatId != allSeatIds[i] {
			log.Print("Missing seat is ", currentSeatId)
		}
	}
}
