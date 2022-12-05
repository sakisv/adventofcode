package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Stack struct {
	contents []string
	index    int
	number   int
}

type Crates struct {
	stacks map[int]Stack
}

func (s *Stack) Push(c string) {
	s.contents = append(s.contents, c)
}

func (s *Stack) Pop() string {
	if len(s.contents) == 0 {
		log.Fatal("trying to pop from an empty stack. Stack: ", s)
	}

	popped := s.contents[len(s.contents)-1]
	s.contents = s.contents[:len(s.contents)-1]

	return popped
}

func (c *Crates) New(stacksInput []string) {
	// letters are in the same index as numbers, so all we have to do is
	// find the index of the numbers and work upwards from there

	c.stacks = make(map[int]Stack)
	// read the last line first
	for i, v := range strings.Split(stacksInput[len(stacksInput)-1], "") {
		if v == " " {
			continue
		}
		v_int, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("failed to convert ", v, " to integer")
		}
		c.stacks[v_int] = Stack{
			contents: make([]string, 0),
			index:    i,
			number:   v_int,
		}
	}

	// continue reading from second to last line
	for i := len(stacksInput) - 2; i >= 0; i-- {
		lineChars := strings.Split(stacksInput[i], "")

		for k, v := range c.stacks {
			if lineChars[v.index] == " " {
				continue
			}
			v.Push(lineChars[v.index])
			c.stacks[k] = v
		}
	}
}

func getInput() ([]string, []string) {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputSplit := strings.Split(string(contents), "\n\n")

	stacks := strings.Split(inputSplit[0], "\n")
	moves := strings.Split(inputSplit[1], "\n")

	return stacks, moves
}

func main() {
	stacks, _ := getInput()
	crates := Crates{}
	crates.New(stacks)
}
