package main

import (
	_ "fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func getInput(filename string) []string {
	contents, err := ioutil.ReadFile(filename)

	strLines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	if err != nil {
		log.Fatal(err)
	}

	return strLines
}

const (
	CLOSING_PARENTHESES = ")"
	CLOSING_BRACKET = "]"
	CLOSING_CURLY = "}"
	CLOSING_QUOTE = ">"

	OPENING_PARENTHESES = "("
	OPENING_BRACKET = "["
	OPENING_CURLY = "{"
	OPENING_QUOTE = "<"

	UNPROCESSED = "UNPROCESSED"
	CORRUPTED	= "CORRUPTED"
	INCOMPLETE	= "INCOMPLETE"
	GOOD		= "GOOD"
)

var illegalCharacterScore = map[string]int {
	CLOSING_PARENTHESES: 3,
	CLOSING_BRACKET: 57,
	CLOSING_CURLY: 1197,
	CLOSING_QUOTE: 25137,
}

var fixCharacterScore = map[string]int {
	CLOSING_PARENTHESES: 1,
	CLOSING_BRACKET: 2,
	CLOSING_CURLY: 3,
	CLOSING_QUOTE: 4,
}

var openingBrackets = map[string]string {
	OPENING_PARENTHESES: OPENING_PARENTHESES,
	OPENING_BRACKET: OPENING_BRACKET,
	OPENING_CURLY: OPENING_CURLY,
	OPENING_QUOTE:  OPENING_QUOTE,
}

var closingBracketsToOpening = map[string]string {
	CLOSING_PARENTHESES: OPENING_PARENTHESES,
	CLOSING_BRACKET: OPENING_BRACKET,
	CLOSING_CURLY: OPENING_CURLY,
	CLOSING_QUOTE: OPENING_QUOTE,
}

var openingBracketsToClosing = map[string]string {
	OPENING_PARENTHESES:	CLOSING_PARENTHESES,
	OPENING_BRACKET:		CLOSING_BRACKET,
	OPENING_CURLY:			CLOSING_CURLY,
	OPENING_QUOTE:			CLOSING_QUOTE,
}

func (l *Line) Push(newItem string) {
	r := rune(newItem[0])
	l.stack = append(l.stack, r)
}

func (l *Line) Pop() string {
	charCount := len(l.stack)
	if charCount == 0 {
		log.Print("trying to pop from an empty stack. Stack: ", l.stack)
	}

	// get the last element
	ret := l.stack[charCount - 1]

	// remove it from the stack
	l.stack = l.stack[:charCount - 1]
	return string(ret)
}

func LineFromString(line string) Line {
	r := make([]rune, len(line))

	for i, v := range line {
		r[i] = v
	}

	return Line{characters: r, origin: line}
}

type Line struct {
	characters []rune
	stack []rune
	origin string
	status string
	score int
}

func (l *Line) processLine() {
	l.status = UNPROCESSED
	for i:= 0; i < len(l.characters); i++ {
		c := string(l.characters[i])

		if _, ok := openingBrackets[c]; ok {
			l.Push(c)
			continue
		}

		popped := l.Pop()
		if matchingBracket, ok := closingBracketsToOpening[c]; ok {
			if popped != matchingBracket {
				l.status = CORRUPTED
				l.score = illegalCharacterScore[c]
				break
			}
		}
	}

	// if the line is not corrupted and it still has items in the stack, then it's incomplete
	if l.status != CORRUPTED && len(l.stack) > 0 {
		l.status = INCOMPLETE

		stackLength := len(l.stack)
		for i := 0; i < stackLength; i++ {
			popped := l.Pop()

			if matchingBracket, ok := openingBracketsToClosing[popped]; ok {
				l.score = l.score * 5 + fixCharacterScore[matchingBracket]
			} else {
				log.Print("Found a closing bracket in an incomplete line, should be corrupted?")
			}
		}
	}
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	input := getInput(filename)

	scoreSum := 0
	var incompleteLinesScore []int
	for _, line := range input {
		l := LineFromString(line)
		l.processLine()
		if l.status == CORRUPTED {
			scoreSum += l.score
		}

		if l.status == INCOMPLETE {
			incompleteLinesScore = append(incompleteLinesScore, l.score)
		}
	}
	log.Print(scoreSum)

	sort.Ints(incompleteLinesScore)
	log.Print(incompleteLinesScore[len(incompleteLinesScore)/2])
}
