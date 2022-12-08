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
	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	numbers := make([][]int, len(lines))

	for i, line := range lines {
		columns := strings.Split(line, "")
		tmp := make([]int, len(columns))
		for j, v := range columns {
			tmp[j], _ = strconv.Atoi(v)
		}
		numbers[i] = tmp
	}

	return numbers
}

func isTallestTree(v int, input []int) bool {
	for _, item := range input {
		if item >= v {
			return false
		}
	}

	return true
}

func main() {
	input := getInput()

	visibleFromOutside := 0
	for i, line := range input {
		for j, currentTree := range line {
			// this is an edge element, it is visible, we move on
			if i == 0 || j == 0 || i == len(input)-1 || j == len(line)-1 {
				visibleFromOutside += 1
				continue
			}

			// for each tree we check top, down, left and right

			// first with rows
			colBefore := make([]int, i)
			colAfter := make([]int, len(input[i+1:]))

			// before
			for tmpI := 0; tmpI < i; tmpI++ {
				colBefore = append(colBefore, input[tmpI][j])
			}
			if isTallestTree(currentTree, colBefore) {
				visibleFromOutside += 1
				continue
			}

			// after
			for tmpI := i + 1; tmpI < len(input); tmpI++ {
				colAfter = append(colAfter, input[tmpI][j])
			}
			if isTallestTree(currentTree, colAfter) {
				visibleFromOutside += 1
				continue
			}

			// get row items
			rowBefore := make([]int, j)
			rowAfter := make([]int, len(line[j+1:]))

			// before
			for tmpJ := 0; tmpJ < j; tmpJ++ {
				rowBefore = append(rowBefore, input[i][tmpJ])
			}
			if isTallestTree(currentTree, rowBefore) {
				visibleFromOutside += 1
				continue
			}

			// after
			for tmpJ := j + 1; tmpJ < len(line); tmpJ++ {
				rowAfter = append(rowAfter, input[i][tmpJ])
			}
			if isTallestTree(currentTree, rowAfter) {
				visibleFromOutside += 1
				continue
			}
		}
	}

	log.Print(visibleFromOutside)
}
