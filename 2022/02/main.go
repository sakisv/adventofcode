package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	WIN  = 6
	DRAW = 3
	LOSE = 0

	ROCK     = "rock"
	PAPER    = "paper"
	SCISSORS = "scissors"
)

type Shape struct {
	letter      string
	description string
	score       int
}

func (s *Shape) New(letter string) {
	s.letter = letter
	if letter == "A" {
		s.description = ROCK
		s.score = 1
	} else if letter == "B" {
		s.description = PAPER
		s.score = 2
	} else if letter == "C" {
		s.description = SCISSORS
		s.score = 3
	}
}

func getWinForDescription(description string) (string, int) {
	if description == ROCK {
		return PAPER, 2
	}
	if description == PAPER {
		return SCISSORS, 3
	}
	return ROCK, 1
}

func getLossForDescription(description string) (string, int) {
	if description == ROCK {
		return SCISSORS, 3
	}
	if description == PAPER {
		return ROCK, 1
	}
	return PAPER, 2
}

func (s *Shape) NewFromOpponent(letter string, opponent Shape) {
	s.letter = letter
	if letter == "Z" {
		s.description, s.score = getWinForDescription(opponent.description)
	}

	if letter == "Y" {
		s.description = opponent.description
		s.score = opponent.score
	}

	if letter == "X" {
		s.description, s.score = getLossForDescription(opponent.description)
	}
}

type Game struct {
	player1 Shape
	player2 Shape
}

func (g *Game) New(line string) {
	players := strings.Split(line, " ")
	s1 := Shape{}
	s2 := Shape{}
	s1.New(players[0])
	s2.NewFromOpponent(players[1], s1)
	g.player1 = s1
	g.player2 = s2
}

func (g *Game) result() int {
	result := 0
	if g.player1.description == ROCK {
		if g.player2.description == PAPER {
			result = WIN
		}
		if g.player2.description == ROCK {
			result = DRAW
		}
		if g.player2.description == SCISSORS {
			result = LOSE
		}
	}

	if g.player1.description == PAPER {
		if g.player2.description == PAPER {
			result = DRAW
		}
		if g.player2.description == ROCK {
			result = LOSE
		}
		if g.player2.description == SCISSORS {
			result = WIN
		}
	}

	if g.player1.description == SCISSORS {
		if g.player2.description == PAPER {
			result = LOSE
		}
		if g.player2.description == ROCK {
			result = WIN
		}
		if g.player2.description == SCISSORS {
			result = DRAW
		}
	}

	// log.Print(g.player1.description, " VS ", g.player2.description, ": RESULT: ", result)
	return result + g.player2.score
}

func getInput() []Game {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	allLines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	games := make([]Game, len(allLines))
	for i, line := range allLines {
		g := Game{}
		g.New(line)
		games[i] = g
	}

	return games
}

func main() {
	games := getInput()

	totalScore := 0
	for _, g := range games {
		totalScore += g.result()
	}

	log.Print(totalScore)
}
