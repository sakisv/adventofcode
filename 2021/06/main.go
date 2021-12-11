package main

import (
	"io/ioutil"
	"log"
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

	numberOfDays := 256

	fishDayCounts := make(map[int]int, 9)
	for _, fishDay := range input {
		fishDayCounts[fishDay]++
	}

	log.Print(fishDayCounts)
	for d := 0; d < numberOfDays; d++ {
		d0Pointer := d % 9		// this will give us a loop going between 0-8
		d6Pointer := (d0Pointer + 7) % 9

		fishDayCounts[d6Pointer] += fishDayCounts[d0Pointer]
		//log.Print(fishDayCounts)
	}

	totalFish := 0
	for _, v := range fishDayCounts {
		totalFish += v
	}

	log.Print(totalFish)
}
