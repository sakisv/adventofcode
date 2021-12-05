package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getInput() []string {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(contents)), "\n")
}

type Point struct {
	x, y int
}

type Line struct {
	start Point
	end Point
}

func NewPoint(input string) Point {
	p := Point{}

	split := strings.Split(input, ",")
	p.x, _ = strconv.Atoi(split[0])
	p.y, _ = strconv.Atoi(split[1])

	return p
}

func (p *Point) ToString() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func getVerticalLine(x int, start, end Point) []Point {
	var points []Point
	minY := 0
	maxY := 0

	if start.y > end.y {
		minY = end.y
		maxY = start.y
	} else {
		minY = start.y
		maxY = end.y
	}

	for i := minY; i <= maxY; i++ {
		points = append(points, Point{x, i})
	}

	return points
}

func getHorizontalLine(y int, start, end Point) []Point {
	var points []Point
	minX := 0
	maxX := 0

	if start.x > end.x {
		minX = end.x
		maxX = start.x
	} else {
		minX = start.x
		maxX = end.x
	}

	for i := minX; i <= maxX; i++ {
		points = append(points, Point{i, y})
	}

	return points
}

func NewLine(input string) Line {
	l := Line{}

	split := strings.Split(input, " -> ")
	start, end := split[0], split[1]

	l.start = NewPoint(start)
	l.end = NewPoint(end)

	return l
}

func (l *Line) GetCoveringPoints() []Point {
	var points []Point

	if l.start.x == l.end.x {
		return getVerticalLine(l.start.x, l.start, l.end)
	}

	if l.start.y == l.end.y {
		return getHorizontalLine(l.start.y, l.start, l.end)
	}

	return points
}

func main() {
	input := getInput()
	var pointMap = make(map[string]int)

	var lines []Line
	count := 0
	for _, line := range input {
		l := NewLine(line)
		lines = append(lines, NewLine(line))
		coveringPoints := l.GetCoveringPoints()

		for _, point := range coveringPoints {
			pointMap[point.ToString()] += 1
		}
	}

	count = 0
	for _, v := range pointMap {
		if v >= 2 {
			count += 1
		}
	}

	log.Print(count)
}
