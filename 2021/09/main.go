package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) [][]int {
	contents, err := ioutil.ReadFile(filename)

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

type Cell struct {
	row, column, value int
}

func (c *Cell) isLowerThanNeighbors(field [][]int) bool {
	if c.value == 9 {
		return false
	}

	if c.row - 1 >= 0 && field[c.row - 1][c.column] < c.value {
		return false
	}
	if c.row + 1 < len(field) && field[c.row + 1][c.column] < c.value {
		return false
	}

	if c.column - 1 >= 0 && field[c.row][c.column - 1] < c.value {
		return false
	}
	if c.column + 1 < len(field[0]) && field[c.row][c.column + 1] < c.value {
		return false
	}

	return true
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	input := getInput(filename)

	riskLevelSum := 0
	for row := 0; row < len(input); row++ {
		for column := 0; column < len(input[row]); column++ {
			c := Cell{row, column, input[row][column]}
			if c.isLowerThanNeighbors(input) {
				riskLevelSum += c.value + 1
			}
		}
	}

	log.Print(riskLevelSum)
}
