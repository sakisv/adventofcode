package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Password struct {
	minCount        int
	maxCount        int
	password        string
	targetCharacter string
}

func (p *Password) New(line string) {
	spaceSplit := strings.Split(line, " ")

	p.password = spaceSplit[2]

	character := spaceSplit[1]
	p.targetCharacter = strings.Split(character, ":")[0]

	counts := spaceSplit[0]
	minMaxSplit := strings.Split(counts, "-")
	p.minCount, _ = strconv.Atoi(minMaxSplit[0])
	p.maxCount, _ = strconv.Atoi(minMaxSplit[1])
}

func (p Password) isValid() bool {
	charCount := strings.Count(p.password, p.targetCharacter)

	if charCount > p.maxCount {
		return false
	}

	if charCount < p.minCount {
		return false
	}

	return true
}

func getInput() []Password {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")

	passwords := make([]Password, len(lines))
	for _, v := range lines {
		p := Password{}
		p.New(v)
		passwords = append(passwords, p)
	}
	return passwords
}

func main() {
	input := getInput()

	validPasswords := 0
	for _, v := range input {
		if v.isValid() {
			validPasswords++
		}
	}

	log.Print(validPasswords)
}
