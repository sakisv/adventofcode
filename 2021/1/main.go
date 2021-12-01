package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFromFile() []int {
	var int_slice []int
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range strings.Split(strings.TrimSpace(string(contents)), "\n") {
		intv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		int_slice = append(int_slice, intv)
	}

	return int_slice
}

func main() {
	input := readFromFile()

	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
	}

	log.Print(count)
}
