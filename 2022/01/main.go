package main

import (
	"io/ioutil"
	"log"
	"sort"
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

	calsPerElf := make([]int, len(calLinesPerElf))
	for i, elfLines := range calLinesPerElf {

		elfSum := 0
		for _, calString := range strings.Split(strings.TrimSpace(string(elfLines)), "\n") {
			calInt, err := strconv.Atoi(calString)
			if err != nil {
				log.Fatal("could not convert ", calString, " to int")
			}
			elfSum += calInt
		}

		calsPerElf[i] = elfSum
	}

	sort.Ints(calsPerElf)
	topThreeTotalCals := 0
	for _, v := range calsPerElf[len(calsPerElf)-3 : len(calLinesPerElf)] {
		log.Print(v)
		topThreeTotalCals += v
	}
	log.Print(topThreeTotalCals)

}
