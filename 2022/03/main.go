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

func findItemPriority(item string) int {
	ret := -1
	for i, v := range strings.Split(CHARS, "") {
		if v == item {
			ret = i + 1
		}
	}

	return ret
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

func findCommonItem(rucksacks []Rucksack) string {
	allItems := make(map[string]int)
	for _, r := range rucksacks {
		rucksackSet := make(map[string]bool)
		for _, item := range r.contents {
			if rucksackSet[item] {
				continue
			}
			rucksackSet[item] = true
			allItems[item] += 1
		}
	}

	commonItem := ""
	for k, v := range allItems {
		if v == 3 {
			commonItem = k
		}
	}

	return commonItem
}

func (r *Rucksack) New(line string) {
	r.contents = strings.Split(line, "")
	r.comp1 = r.contents[:len(r.contents)/2]
	r.comp2 = r.contents[len(r.contents)/2:]
	r.findDoubleItem()
	r.doubleItemPriority = findItemPriority(r.doubleItem)
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

	commonItemSum := 0
	for i := 0; i < len(rucksacks); i = i + 3 {
		rucksacksToCompare := []Rucksack{rucksacks[i], rucksacks[i+1], rucksacks[i+2]}
		commonItem := findCommonItem(rucksacksToCompare)
		commonItemSum += findItemPriority(commonItem)
	}

	log.Print(commonItemSum)
}
