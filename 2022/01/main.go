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

	return strings.Split(strings.TrimSpace(string(contents)), "\n\n")
}

func main() {
	calLinesPerElf := getInput()

	maxCals := 0
	for _, elfLines := range calLinesPerElf {

		elfSum := 0
		for _, calString := range strings.Split(strings.TrimSpace(string(elfLines)), "\n") {
			calInt, err := strconv.Atoi(calString)
			if err != nil {
				log.Fatal("could not convert ", calString, " to int")
			}
			elfSum += calInt
		}

		if elfSum > maxCals {
			maxCals = elfSum
		}
	}
	log.Print(maxCals)
}
