package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type Passport struct {
	details string
}

func (p *Passport) New(lines string) {
	if strings.Contains(lines, "\n") {
		lines = strings.ReplaceAll(lines, "\n", " ")
	}
	p.details = lines
}

func (p *Passport) isValid() bool {
	sortedFieldNames := strings.Join(getSortedFieldNames(p.details), "")

	vp1 := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	vp2 := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	sort.Strings(vp1)
	sort.Strings(vp2)
	validPassport1 := strings.Join(vp1, "")
	validPassport2 := strings.Join(vp2, "")

	if sortedFieldNames == validPassport1 || sortedFieldNames == validPassport2 {
		return true
	}
	return false
}

func getSortedFieldNames(passportDetails string) []string {
	fields := strings.Split(passportDetails, " ")

	fieldNames := make([]string, len(fields))
	for i, v := range fields {
		fieldNames[i] = strings.Split(v, ":")[0]
	}
	sort.Strings(fieldNames)
	return fieldNames
}

func getInput() []Passport {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	passportsLines := strings.Split(strings.TrimSpace(string(contents)), "\n\n")
	passports := make([]Passport, 0)

	for _, line := range passportsLines {
		p := Passport{}
		p.New(line)
		passports = append(passports, p)
	}

	return passports
}

func main() {
	input := getInput()

	validPassports := 0
	for _, v := range input {
		if v.isValid() {
			validPassports++
		}
	}

	log.Print(validPassports)
}
