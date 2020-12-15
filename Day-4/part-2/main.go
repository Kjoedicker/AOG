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

func verifyValues(values map[string]string) bool {
	required := []string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}

	check := map[string]string{
		"byr": "19[2-9][0-9]$|200[0-2]$",
		"iyr": "201[0-9]$|2020$",
		"ecl": "amb$|blu$|brn$|gry$|grn$|hzl$|oth$",
		"eyr": "202[0-9]$|2030$",
		"hgt": "1[5-8][0-9]cm$|19[0-3]cm$|59in$|6[0-9]in$|7[0-6]in$",
		"hcl": "^#[a-f0-9]{6}$",
		"pid": "^[0-9]{9}$",
	}

	for _, value := range required {
		regex, _ := regexp.Compile(check[value])
		matchedRegex := regex.MatchString(values[value])

		if matchedRegex == false {
			return false
		}
	}

	return true
}

func checkFields(values map[string]string) bool {
	required := []string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}

	counter := 0
	for _, idx := range required {
		if values[idx] != " " {
			counter = counter + 1
		}
	}

	if counter == 7 {
		return true
	}

	return false
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

		ok := checkFields(values)
		if ok {

			if verifyValues(values) {
				fmt.Println(values)
				counter = counter + 1
			}
		}

	}
	fmt.Println(counter)
}
