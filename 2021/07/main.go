package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func getInput() []int {
	var ret []int
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range strings.Split(strings.TrimSpace(string(contents)), ",") {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		ret = append(ret, val)
	}
	return ret
}

func main() {
	input := getInput()
	positionCount := make(map[int]int, len(input))

	for _, p := range input {
		positionCount[p]++
	}

	fuelConsumption := make(map[int]int, len(input))
	for i := 0; i < len(input); i++ {
		for position, count := range positionCount {
			fuelConsumption[i] += int(math.Abs(float64(position - i))) * count
		}
	}

	minPosition := 0
	minConsumption := fuelConsumption[minPosition]
	for position, consumption := range fuelConsumption {
		if consumption < minConsumption {
			minConsumption = consumption
			minPosition = position
		}
	}

	log.Print(minConsumption)
	log.Print(minPosition)
	log.Print(fuelConsumption)
}
