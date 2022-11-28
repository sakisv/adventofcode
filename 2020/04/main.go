package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Passport struct {
	fields  map[string]string
	details string
}

func (p *Passport) New(lines string) {
	if strings.Contains(lines, "\n") {
		lines = strings.ReplaceAll(lines, "\n", " ")
	}
	p.details = lines
	p.fields = make(map[string]string)

	for _, field := range getFields(p.details) {
		fieldSplit := strings.Split(field, ":")
		fieldName := fieldSplit[0]
		fieldValue := fieldSplit[1]

		p.fields[fieldName] = fieldValue
	}
}

func (p *Passport) hasRequiredFields() bool {
	sortedFieldNames := strings.Join(getSortedFieldNames(p.details), "")

	vp1 := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	vp2 := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	sort.Strings(vp1)
	sort.Strings(vp2)
	validPassport1 := strings.Join(vp1, "")
	validPassport2 := strings.Join(vp2, "")

	if sortedFieldNames != validPassport1 && sortedFieldNames != validPassport2 {
		return false
	}

	return true
}

func (p *Passport) validNumber(v string, min, max int) bool {
	number, err := strconv.Atoi(v)

	if err != nil {
		return false
	}
	if number < min {
		return false
	}
	if number > max {
		return false
	}
	return true
}

func (p *Passport) validHeight(v string) bool {
	if strings.HasSuffix(v, "in") {
		return p.validNumber(strings.Split(v, "in")[0], 59, 76)
	}
	if strings.HasSuffix(v, "cm") {
		return p.validNumber(strings.Split(v, "cm")[0], 150, 193)
	}
	return false
}

func (p *Passport) validString(v string, pattern string) bool {
	matched, err := regexp.MatchString(pattern, v)

	if err != nil {
		return false
	}

	return matched
}

func (p *Passport) isValid() bool {
	if !p.hasRequiredFields() {
		return false
	}

	for k, v := range p.fields {
		if k == "byr" {
			if !p.validNumber(v, 1920, 2002) {
				return false
			}
		}

		if k == "iyr" {
			if !p.validNumber(v, 2010, 2020) {
				return false
			}
		}

		if k == "eyr" {
			if !p.validNumber(v, 2020, 2030) {
				return false
			}
		}

		if k == "hgt" {
			if !p.validHeight(v) {
				return false
			}
		}

		if k == "hcl" {
			if !p.validString(v, `#[0-9a-f]{6}`) {
				return false
			}
		}

		if k == "ecl" {
			if !p.validString(v, `amb|blu|brn|gry|grn|hzl|oth`) {
				return false
			}
		}

		if k == "pid" {
			if !p.validString(v, `[0-9]{9}`) {
				return false
			}
		}
	}
	return true
}

func getFields(passportDetails string) []string {
	fields := strings.Split(passportDetails, " ")
	return fields
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
