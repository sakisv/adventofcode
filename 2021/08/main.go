package main

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func getLines() []string {
	var ret []string
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(strings.TrimSpace(string(contents)), "\n") {
		if err != nil {
			log.Fatal(err)
		}

		ret = append(ret, line)
	}
	return ret
}

func splitLine(line string) (string, string) {
	s := strings.Split(line, " | ")

	return s[0], s[1]
}

func extractOneFourSeven(patterns []string) (string, string, string) {
	var one, four, seven string

	for _, s := range patterns {
		if len(s) == 2 {
			one = s
		}
		if len(s) == 3 {
			seven = s
		}
		if len(s) == 4 {
			four = s
		}
	}

	return one, four, seven
}

// 0: a, b, c, e, f, g		-> 6 segments
// 1: c, f					-> 2 segments
// 2: a, c, d, e, g			-> 5 segments
// 3: a, c, d, f, g			-> 5 segments
// 4: b, c, d, f			-> 4 segments
// 5: a, b, d, f, g			-> 5 segments
// 6: a, b, d, e, f, g		-> 6 segments
// 7: a, c, f				-> 3 segments
// 8: a, b, c, d, e, f, g	-> 7 segments
// 9: a, b, c, d, f, g		-> 6 segments

type Pattern struct {
	top                     string
	topLeft, topRight       string
	middle                  string
	bottomLeft, bottomRight string
	bottom                  string
}

func (p *Pattern) NumberFromString(input string) int {
	segments := make(map[string]int, 8)

	zero := make(map[string]int, 1)
	one := map[string]int{"tr": 1, "br": 1}
	two := map[string]int{"t": 1, "tr": 1, "m": 1, "bl": 1, "b": 1}
	three := map[string]int{"t": 1, "tr": 1, "m": 1, "br": 1, "b": 1}
	four := map[string]int{"tl": 1, "tr": 1, "m": 1, "br": 1}
	five := map[string]int{"t": 1, "tl": 1, "m": 1, "br": 1, "b": 1}
	six := map[string]int{"t": 1, "tl": 1, "m": 1, "bl": 1, "br": 1, "b": 1}
	seven := map[string]int{"t": 1, "tr": 1, "br": 1}
	eight := map[string]int{"t": 1, "tl": 1, "tr": 1, "m": 1, "bl": 1, "br": 1, "b": 1}
	nine := map[string]int{"t": 1, "tl": 1, "tr": 1, "m": 1, "br": 1, "b": 1}

	for _, k := range input {
		l := string(k)
		if l == p.top {
			segments["t"]++
			continue
		}
		if l == p.topLeft {
			segments["tl"]++
			continue
		}
		if l == p.topRight {
			segments["tr"]++
			continue
		}
		if l == p.middle {
			segments["m"]++
			continue
		}
		if l == p.bottomLeft {
			segments["bl"]++
			continue
		}
		if l == p.bottomRight {
			segments["br"]++
			continue
		}
		if l == p.bottom {
			segments["b"]++
			continue
		}
	}

	numbers := []map[string]int{zero, one, two, three, four, five, six, seven, eight, nine}
	for index, item := range numbers {
		if len(item) == len(segments) {
			fullMatch := true
			for k, v := range item {
				if segments[k] != v {
					fullMatch = false
				}
			}

			if fullMatch {
				return index
			}
		}
	}
	return 0
}

func NewPattern(patterns string) Pattern {
	letterCount := make(map[string]int, 7)
	reverseLetterCount := make(map[int][]string, 7)

	// count the letters and inverse them
	for i := 0; i < len(patterns); i++ {
		if string(patterns[i]) == " " {
			continue
		}
		letterCount[string(patterns[i])]++
	}

	for k, v := range letterCount {
		reverseLetterCount[v] = append(reverseLetterCount[v], k)
	}

	// we have 7 segments:
	// top, top-right, top-left, middle, bottom-right, bottom-left, bottom
	// which correspond to the "original" letters:
	// a, b, c, d, e, f, g
	//
	// We can count how many times each letter is used and deduce which segment it is
	// e.g: bottom-right is used 9 times
	// which means that we can straight away find some segments
	//
	// Specifically the counts are:
	// t and tr -> 8, m and b -> 7, tl -> 6, bl -> 4, br -> 9

	tl := reverseLetterCount[6][0]
	bl := reverseLetterCount[4][0]
	br := reverseLetterCount[9][0]

	// for the remaining ones we cross check with the easy numbers, i.e. 1, 4 and 7
	one, four, _ := extractOneFourSeven(strings.Fields(patterns))

	// we know bottom-right, we use `one` to extract the top-right
	tr := ""
	if br == string(one[0]) {
		tr = string(one[1])
	} else {
		tr = string(one[0])
	}

	// now that we know top-right we also know the "top" since they have the same count:
	t := ""
	if tr == reverseLetterCount[8][0] {
		t = reverseLetterCount[8][1]
	} else {
		t = reverseLetterCount[8][0]
	}

	// finally we can deduce the middle and bottom using `four`
	m := ""
	for i := 0; i < len(four); i++ {
		letter := string(four[i])
		if letter != tl && letter != tr && letter != br {
			m = letter
			break
		}
	}

	b := ""
	if m == reverseLetterCount[7][0] {
		b = reverseLetterCount[7][1]
	} else {
		b = reverseLetterCount[7][0]
	}

	return Pattern{top: t, topLeft: tl, topRight: tr, middle: m, bottomLeft: bl, bottomRight: br, bottom: b}
}

func main() {
	lines := getLines()

	outputSum := 0
	for _, line := range lines {
		patterns, outputValues := splitLine(line)

		p := NewPattern(patterns)
		exp := 3
		for _, v := range strings.Fields(outputValues) {
			outputSum += int(math.Pow10(exp)) * p.NumberFromString(v)
			exp -= 1
		}
	}

	log.Print(outputSum)
}
