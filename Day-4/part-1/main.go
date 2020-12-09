package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkValues(values map[string]string) int {
	required := []string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}

	counter := 0
	for _, idx := range required {
		if values[idx] != " " {
			counter = counter + 1
		}
	}

	if counter == 7 {
		return 1
	}

	return 0
}

func main() {

	data, err := ioutil.ReadFile("input.txt")
	check(err)

	counter := 0
	parsedData := strings.Split(string(data), "\n\n")
	for _, value := range parsedData {
		splitData := strings.Split(value, " ")

		regex, _ := regexp.Compile("[a-z]*:#?[a-z0-9]*")
		matchedData := regex.FindAllString(strings.Join(splitData, " "), 8)

		values := map[string]string{
			"byr": " ",
			"cid": " ",
			"ecl": " ",
			"eyr": " ",
			"hcl": " ",
			"hgt": " ",
			"iyr": " ",
			"pid": " ",
		}

		for _, idx := range matchedData {
			splitMatched := strings.Split(idx, ":")
			values[splitMatched[0]] = splitMatched[1]
		}
		counter = counter + checkValues(values)
	}

	fmt.Println(counter)
}
