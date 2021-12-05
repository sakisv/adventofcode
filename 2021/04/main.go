package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getInput() ([]int, []string) {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(contents))
	split_input := strings.SplitN(input, "\n", 2)
	numbers_drawn, boards := split_input[0], split_input[1]

	return strSliceToIntSlice(strings.Split(numbers_drawn, ",")), strings.Split(boards, "\n\n")
}

func strSliceToIntSlice(input []string) ([]int) {
	var ret []int

	for _, v := range input {
		v_int, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		ret = append(ret, v_int)
	}

	return ret
}

type Square struct {
	number int
	marked bool
}

type Board struct {
	squares [][]Square
	all_numbers map[int]int
}

func NewBoard(input string) Board {
	b := Board{}
	rows := strings.Split(input, "\n")

	b.squares = make([][]Square, len(rows))
	b.all_numbers = make(map[int]int)
	for i, row := range rows {
		numbers := strSliceToIntSlice(strings.Fields(row))

		b.squares[i] = make([]Square, len(numbers))
		for j, number := range numbers {
			b.squares[i][j] = Square{number, false}
			b.all_numbers[number] += 1
		}
	}

	return b
}

func (b *Board) MarkNumber(number int) {
	for i, row := range b.squares {
		for j, square := range row {
			if square.number == number {
				b.squares[i][j].marked = true
			}
		}
	}
}

func (b *Board) HasCompletedRow() bool {
	count := 0
	for _, row := range b.squares {
		count = 0
		for _, square := range row {
			if square.marked {
				count += 1
			}
		}

		if count == len(row) {
			return true
		}
	}
	return false
}

func (b *Board) HasCompletedColumn() bool {
	count := 0
	for col := 0; col < len(b.squares[0]); col ++ {
		count = 0
		for row := 0; row < len(b.squares); row++ {
			if b.squares[row][col].marked {
				count += 1
			}
		}
		if count == len(b.squares[0]) {
			return true
		}
	}
	return false
}

func (b *Board) CalculateScore(winning_number int) int {
	sum := 0
	for _, row := range b.squares {
		for _, square := range row {
			if ! square.marked {
				sum += square.number
			}
		}
	}

	return sum * winning_number
}

func main() {
	numbers_drawn, board_strings := getInput()
	var winning_boards_map = make(map[int]int)
	var winning_boards []int

	var boards []Board
	for _, board_string := range board_strings {
		board := NewBoard(strings.TrimSpace(board_string))

		boards = append(boards, board)
	}

	log.Print(numbers_drawn)
	log.Print(len(board_strings))

	for _, number := range numbers_drawn {
		for board_count, board := range boards {
			if _, exists := winning_boards_map[board_count]; exists {
				continue
			}
			if _, exists := board.all_numbers[number]; ! exists {
				log.Print("Number ", number, " doesn't exist in board ", board_count, " skipping")
				continue
			}

			log.Print("Marking number ", number, " in board ", board_count)
			board.MarkNumber(number)

			if board.HasCompletedColumn() || board.HasCompletedRow() {
				winning_boards_map[board_count] = board.CalculateScore(number)
				winning_boards = append(winning_boards, board_count)
			}
		}
	}

	log.Print(winning_boards_map[winning_boards[len(winning_boards) - 1]])
}
