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

func boolToString(flag bool) string {
	ret := "1"
	if ! flag {
		ret = "0"
	}

	return ret
}

func getRating(input []string, index int, mostCommon bool) string {
	input_length := len(input)
	if input_length == 1 {
		return input[0]
	}

	bit_count := 0
	for _, line := range input {
		bit, _ := strconv.Atoi(string(line[index]))
		bit_count += bit
	}

	// get a "0" or "1" as a string
	desired_bit := boolToString(mostCommon)
	if bit_count < input_length / 2 {
		desired_bit = boolToString(!mostCommon)
	}

	var next_input []string
	for _, line := range input {
		if string(line[index]) == desired_bit {
			next_input = append(next_input, line)
		}
	}

	return getRating(next_input, index + 1, mostCommon)
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

	log.Print("Part 1")
	log.Print(bit_counts)
	log.Print(strings.Join(gamma_rate, ""))
	log.Print(strings.Join(epsilon_rate, ""))

	gr_decimal, _ := strconv.ParseInt(strings.Join(gamma_rate, ""), 2, 16)
	er_decimal, _ := strconv.ParseInt(strings.Join(epsilon_rate, ""), 2, 16)
	log.Print(gr_decimal)
	log.Print(er_decimal)
	log.Print(gr_decimal * er_decimal)

	log.Print("Part 2")
	oxygen_generator_rating := getRating(input, 0, true)
	co2_scrubber_rating := getRating(input, 0, false)

	oxygen_generator_rating_decimal, _ := strconv.ParseInt(oxygen_generator_rating, 2, 16)
	co2_scrubber_rating_decimal, _ := strconv.ParseInt(co2_scrubber_rating, 2, 16)

	log.Print(oxygen_generator_rating_decimal)
	log.Print(co2_scrubber_rating_decimal)
	log.Print(oxygen_generator_rating_decimal * co2_scrubber_rating_decimal)
}
