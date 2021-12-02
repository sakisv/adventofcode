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

	return strings.Split(strings.TrimSpace(string(contents)), "\n")
}

func main() {
	input := getInput()

	depth := 0
	horizontal := 0
	aim := 0
	for _, v := range input {
		split := strings.Split(v, " ")
		direction, count_str := split[0], split[1]

		count, err := strconv.Atoi(count_str)
		if err != nil {
			log.Fatal(err)
		}

		if direction == "forward" {
			horizontal += count
			depth = depth + (count * aim)
			continue
		}

		if direction == "down" {
			aim += count
		}

		if direction == "up" {
			aim -= count
		}

	}
	log.Print(depth * horizontal)
}
