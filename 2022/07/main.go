package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type Directory struct {
	name        string
	directories []Directory
}

func getInput() []string {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(contents), "")
}

func main() {
	d := Directory{}
	log.Print(d)
}
