package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getInput() []string {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	allLines := strings.Split(strings.TrimSpace(string(contents)), "\n")

	return allLines
}

type Assignment struct {
	start int
	end   int
}

func (a *Assignment) New(startEnd string) {
	split := strings.Split(startEnd, "-")
	a.start, _ = strconv.Atoi(split[0])
	a.end, _ = strconv.Atoi(split[1])
}

func AssignmentsFullyOverlap(a, b Assignment) bool {
	// first within second?
	if a.start >= b.start && a.end <= b.end {
		return true
	}

	// second within first?
	if b.start >= a.start && b.end <= a.end {
		return true
	}

	return false
}

func main() {
	lines := getInput()

	overlappingPairs := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		a1 := Assignment{}
		a1.New(pairs[0])
		a2 := Assignment{}
		a2.New(pairs[1])

		if AssignmentsFullyOverlap(a1, a2) {
			overlappingPairs += 1
		}
	}

	log.Print(overlappingPairs)
}
