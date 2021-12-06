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

	numberOfDays := 80

	for day := 0; day<numberOfDays; day++ {

		newFishCount := 0
		for i, lanternfish := range input {
			if lanternfish > 0 {
				input[i] -= 1
				continue
			}

			if lanternfish == 0 {
				input[i] = 6
				newFishCount++
				continue
			}
		}

		for i := 0; i<newFishCount; i++ {
			input = append(input, 8)
		}
		log.Print("Day ", day+1, " total fish: ", len(input))
	}

	log.Print(len(input))
}
