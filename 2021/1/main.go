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
	int_slice := readFromFile()

	count := 0
	for i := 1; i < len(int_slice); i++ {
		if int_slice[i] > int_slice[i-1] {
			count++
		}
	}

	log.Print(count)
}
