package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getInput(filename string) [][]int {
	contents, err := ioutil.ReadFile(filename)
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

func getScenicScore(v int, input []int) int {
	//log.Print("====> scenic score input: ", input)
	treesItCanSee := 0
	for _, tree := range input {
		treesItCanSee += 1
		if tree >= v {
			break
		}
	}

	return treesItCanSee
}

func main() {
	input := getInput("input.txt")

	visibleFromOutside := 0
	maxScenicScore := 0
	for i, line := range input {
		for j, currentTree := range line {
			//log.Print("Tree: ", currentTree, "<", i, j, ">")
			treeScenicScore := 1
			// this is an edge element, it is visible, we move on
			if i == 0 || j == 0 || i == len(input)-1 || j == len(line)-1 {
				visibleFromOutside += 1
				continue
			}

			// for each tree we check top, down, left and right

			// first with rows
			colBefore := make([]int, 0)
			colAfter := make([]int, 0)

			// before
			// for the trees that come _before_ ours we need to add them to the
			// list in reverse
			for tmpI := i - 1; tmpI >= 0; tmpI-- {
				colBefore = append(colBefore, input[tmpI][j])
			}
			treesItCanSee := getScenicScore(currentTree, colBefore)
			//log.Print("=> Trees it can see up: ", treesItCanSee)
			treeScenicScore *= treesItCanSee

			// after
			for tmpI := i + 1; tmpI < len(input); tmpI++ {
				colAfter = append(colAfter, input[tmpI][j])
			}
			treesItCanSee = getScenicScore(currentTree, colAfter)
			//log.Print("=> Trees it can see down: ", treesItCanSee)
			treeScenicScore *= treesItCanSee

			// get row items
			rowBefore := make([]int, 0)
			rowAfter := make([]int, 0)

			// before
			// for the trees that come _before_ ours we need to add them to the
			// list in reverse
			for tmpJ := j - 1; tmpJ >= 0; tmpJ-- {
				rowBefore = append(rowBefore, line[tmpJ])
			}
			treesItCanSee = getScenicScore(currentTree, rowBefore)
			//log.Print("=> Trees it can see left: ", treesItCanSee)
			treeScenicScore *= treesItCanSee

			// after
			for tmpJ := j + 1; tmpJ < len(line); tmpJ++ {
				rowAfter = append(rowAfter, line[tmpJ])
			}
			treesItCanSee = getScenicScore(currentTree, rowAfter)
			//log.Print("=> Trees it can see right: ", treesItCanSee)
			treeScenicScore *= treesItCanSee

			//log.Print("==> Scenic score: ", treeScenicScore)

			if treeScenicScore > maxScenicScore {
				maxScenicScore = treeScenicScore
			}
		}
	}

	log.Print(maxScenicScore)
}
