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

	return strings.Split(string(contents), "")
}

func countOccurences(characters map[string]int) int {
	sum := 0
	for _, v := range characters {
		sum += v
	}

	return sum
}

func getStartMarker(input []string, windowSize int) int {
	uniqCharactersBuffer := make(map[string]int)

	for i, v := range input {
		// first decrease count or remove the oldest character
		if i-windowSize >= 0 {
			oldestChar := input[i-windowSize]
			if uniqCharactersBuffer[oldestChar] > 1 {
				uniqCharactersBuffer[oldestChar] -= 1
			} else {
				delete(uniqCharactersBuffer, oldestChar)
			}
		}

		// then add the new character
		uniqCharactersBuffer[v] += 1

		// then check if the occurences add up to windowSize
		if len(uniqCharactersBuffer) == windowSize && countOccurences(uniqCharactersBuffer) == windowSize {
			return i + 1
		}
	}

	return -1
}

func main() {
	input := getInput()

	startMarker := getStartMarker(input, 4)
	log.Print(startMarker)
	startMarker = getStartMarker(input, 14)
	log.Print(startMarker)
}
