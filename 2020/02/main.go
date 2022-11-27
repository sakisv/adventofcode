package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Password struct {
	position1       int
	position2       int
	password        []rune
	targetCharacter []rune
}

func (p *Password) New(line string) {
	spaceSplit := strings.Split(line, " ")

	p.password = []rune(spaceSplit[2])

	character := spaceSplit[1]
	p.targetCharacter = []rune(strings.Split(character, ":")[0])

	positions := spaceSplit[0]
	positionSplit := strings.Split(positions, "-")
	p.position1, _ = strconv.Atoi(positionSplit[0])
	p.position2, _ = strconv.Atoi(positionSplit[1])
}

func (p *Password) isValid() bool {
	pos1Exists := p.password[p.position1-1] == p.targetCharacter[0]
	pos2Exists := p.password[p.position2-1] == p.targetCharacter[0]

	return pos1Exists != pos2Exists
}

func getInput() []Password {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")

	passwords := make([]Password, len(lines))
	for i, v := range lines {
		p := Password{}
		p.New(v)
		passwords[i] = p
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
