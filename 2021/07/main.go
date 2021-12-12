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

func calculateStepsSum(steps int) int {
	sum := 0
	for i := 0; i <= steps; i++ {
		sum += i
	}

	return sum
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
			steps := int(math.Abs(float64(position - i)))
			fuelConsumption[i] +=  calculateStepsSum(steps) * count
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

	//log.Print(fuelConsumption)
	log.Print(minConsumption)
	log.Print(minPosition)
}
