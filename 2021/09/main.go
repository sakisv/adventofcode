package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

func (c *Cell) getMapKey() string {
	return fmt.Sprintf("%v-%v-%v", c.row, c.column, c.value)
}

func (c *Cell) visitNeighbors(field [][]int, visitedNeighbors map[string]bool) {
	if c.value == 9 {
		return
	}

	if isVisited, exists := visitedNeighbors[c.getMapKey()]; exists && isVisited {
		return
	}

	visitedNeighbors[c.getMapKey()] = true

	if c.row - 1 >= 0 {
		nextCell := Cell{c.row - 1, c.column, field[c.row - 1][c.column]}
		nextCell.visitNeighbors(field, visitedNeighbors)
	}

	if c.row + 1 < len(field) {
		nextCell := Cell{c.row + 1, c.column, field[c.row + 1][c.column]}
		nextCell.visitNeighbors(field, visitedNeighbors)
	}

	if c.column - 1 >= 0 {
		nextCell := Cell{c.row, c.column - 1, field[c.row][c.column - 1]}
		nextCell.visitNeighbors(field, visitedNeighbors)
	}
	if c.column + 1 < len(field[c.row]) {
		nextCell := Cell{c.row, c.column + 1, field[c.row][c.column + 1]}
		nextCell.visitNeighbors(field, visitedNeighbors)
	}

}

func (c *Cell) findBasinSize(field [][]int) int {
	visitedCells := make(map[string]bool)
	c.visitNeighbors(field, visitedCells)

	basinSize := len(visitedCells)

	return basinSize
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	input := getInput(filename)

	riskLevelSum := 0
	var basinSizes []int
	for row := 0; row < len(input); row++ {
		for column := 0; column < len(input[row]); column++ {
			c := Cell{row, column, input[row][column]}
			if c.isLowerThanNeighbors(input) {
				riskLevelSum += c.value + 1
				basinSizes = append(basinSizes, c.findBasinSize(input))
			}
		}
	}

	sort.Ints(basinSizes)
	l := len(basinSizes)
	log.Print(riskLevelSum)
	log.Print(basinSizes[l - 1] * basinSizes[l - 2] * basinSizes[l - 3])
}
