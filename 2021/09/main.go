package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getInput() [][]int {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	strLines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	if err != nil {
		log.Fatal(err)
	}

	ret := make([][]int, len(strLines))
	for i, line := range strLines {
		ret[i] = make([]int, len(line))
		for j, column := range line {
			ret[i][j], err = strconv.Atoi(string(column))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return ret
}

func main() {
	input := getInput()

	rowCount := len(input)
	colCount := len(input[0])
	riskLevelSum := 0
	for i, row := range input {
		for j, _ := range row {
			var adjacentCells []int
			if i-1 >= 0 {
				adjacentCells = append(adjacentCells, input[i-1][j])
			}
			if i+1 < rowCount {
				adjacentCells = append(adjacentCells, input[i+1][j])
			}
			if j-1 >= 0 {
				adjacentCells = append(adjacentCells, input[i][j-1])
			}
			if j+1 < colCount {
				adjacentCells = append(adjacentCells, input[i][j+1])
			}

			lowerCount := 0
			for _, adjacentCell := range adjacentCells {
				if input[i][j] < adjacentCell {
					lowerCount++
				}
			}
			if lowerCount == len(adjacentCells) {
				riskLevelSum += 1 + input[i][j]
			}
		}
	}

	log.Print(riskLevelSum)
}
