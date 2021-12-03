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
	var bit_counts = [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	input := getInput()
	for _, line := range input {

		for i, char := range line {
			bit, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}

			bit_counts[i] += bit
		}
	}

	input_length := len(input)
	var gamma_rate, epsilon_rate []string
	for _, v := range bit_counts {

		// if the bit_count[i] is more than half of the total input size, then the most common bit was 1
		// else, it was 0
		if v > (input_length / 2) {
			gamma_rate = append(gamma_rate, "1")
			epsilon_rate = append(epsilon_rate, "0")
		} else {
			gamma_rate = append(gamma_rate, "0")
			epsilon_rate = append(epsilon_rate, "1")
		}
	}

	log.Print(bit_counts)
	log.Print(strings.Join(gamma_rate, ""))
	log.Print(strings.Join(epsilon_rate, ""))

	gr_decimal, _ := strconv.ParseInt(strings.Join(gamma_rate, ""), 2, 32)
	er_decimal, _ := strconv.ParseInt(strings.Join(epsilon_rate, ""), 2, 32)
	log.Print(gr_decimal)
	log.Print(er_decimal)
	log.Print(gr_decimal * er_decimal)
}
