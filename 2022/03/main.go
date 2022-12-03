package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type Rucksack struct {
	contents           []string
	comp1              []string
	comp2              []string
	doubleItem         string
	doubleItemPriority int
}

const (
	CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func (r *Rucksack) findDoubleItemPriority() {
	for i, v := range strings.Split(CHARS, "") {
		if v == r.doubleItem {
			r.doubleItemPriority = i + 1
		}
	}
}

func (r *Rucksack) findDoubleItem() {
	found := false
	for _, l1 := range r.comp1 {
		for _, l2 := range r.comp2 {
			if l1 == l2 {
				found = true
			}
		}

		if found {
			r.doubleItem = l1
			break
		}
	}
}

func (r *Rucksack) New(line string) {
	r.contents = strings.Split(line, "")
	r.comp1 = r.contents[:len(r.contents)/2]
	r.comp2 = r.contents[len(r.contents)/2:]
	r.findDoubleItem()
	r.findDoubleItemPriority()
}

func getInput() []Rucksack {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	allLines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	rucksacks := make([]Rucksack, len(allLines))
	for i, line := range allLines {
		rucksacks[i] = Rucksack{}
		rucksacks[i].New(line)
	}

	return rucksacks
}

func main() {
	rucksacks := getInput()

	doubleItemPrioritySum := 0
	for _, r := range rucksacks {
		doubleItemPrioritySum += r.doubleItemPriority
	}

	log.Print(doubleItemPrioritySum)
}
