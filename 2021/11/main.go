package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_"sort"
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

var totalFlashes int

type Octopus struct {
	row, column, energy int
	hasFlashed bool
}

func (o *Octopus) GetDictKey() string {
	return fmt.Sprintf("%v-%v", o.row, o.column)
}

func normalizeIndex(i int) int {
	if i < 0  {
		return 0
	}

	if i > 9 {
		return 9
	}

	return i
}

func PrintField(octopuses map[string]*Octopus) string {
	builder := strings.Builder{}
	builder.WriteString("\n")
	for i := 0; i<10; i++ {
		for j := 0; j < 10; j++ {
			tmpOctopus := Octopus{row: i, column: j}
			o := octopuses[tmpOctopus.GetDictKey()]
			builder.WriteString(fmt.Sprintf("%v", o.energy))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func (o *Octopus) Flash(octopuses map[string]*Octopus) {
	if o.hasFlashed {
		return
	}

	o.hasFlashed = true
	log.Printf("Flashing %v", o.GetDictKey())
	for i := normalizeIndex(o.row - 1); i <= normalizeIndex(o.row + 1); i++ {
		for j := normalizeIndex(o.column - 1); j <= normalizeIndex(o.column + 1); j++ {
			if i == o.row && j == o.column {
				continue
			}
			tmpOctopus := Octopus{row: i, column: j}
			nextO := octopuses[tmpOctopus.GetDictKey()]
			if nextO.hasFlashed {
				continue
			}
			log.Printf("Bumping energy of %v", nextO.GetDictKey())
			nextO.energy++
			if nextO.energy > 9 {
				log.Printf("==> Recursive flash of %v", nextO.GetDictKey())
				nextO.Flash(octopuses)
			}
		}
	}
	o.energy = 0
	totalFlashes++
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	input := getInput(filename)

	octopuses := make(map[string]*Octopus)
	for i, row := range input {
		for j, column := range row {
			o := Octopus{row: i, column: j, energy: column, hasFlashed: false}
			octopuses[o.GetDictKey()] = &o
		}
	}

	totalFlashes = 0
	step := 0
	for {
		step++
		totalFlashes = 0
		log.Printf("\n\nStep %v", step)

		// for each step we reset the "flashed" flag of all octopuses
		for _, oct := range octopuses {
			oct.hasFlashed = false
		}

		for _, oct := range octopuses {
			oct.energy++
		}

		for _, oct := range octopuses {
			if oct.energy > 9 {
				oct.Flash(octopuses)
			}
		}

		if totalFlashes == 100 {
			log.Printf("Synchronised at step %v", step)
			break
		}
	}

}
