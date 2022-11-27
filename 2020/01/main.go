package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFromFile() []int {
	var input []int
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range strings.Split(strings.TrimSpace(string(contents)), "\n") {
		intv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, intv)
	}

	return input
}

func main() {
	input := readFromFile()

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(input); k++ {
				if i == k || j == k {
					continue
				}
				if input[i]+input[j]+input[k] == 2020 {
					log.Print(input[i] * input[j] * input[k])
					break
				}
			}
		}
	}
}
