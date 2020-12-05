package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

//based on x and y iter through value and find if it matches the parameters
func followsRules(x int, y int, focus string, value string) int {
	param := regexp.MustCompile(focus)
	matches := len(param.FindAllStringIndex(value, -1))

	if matches >= x && matches <= y {
		return 1
	}

	return 0
}

//convert a slice of strings into a slice of ints
func stoiSlice(input []string) ([]int, bool) {
	new := make([]int, len(input))

	for idx, value := range input {
		var err error
		new[idx], err = strconv.Atoi(value)
		if err != nil {
			return []int(nil), false
		}
	}

	return new, true
}

func matchRegex(input string) (int, int, string, string) {

	regex, _ := regexp.Compile("([1-9]?[0-9])")
	fmt.Println(regex.FindAllString(input, 2))
	constraint, ok := stoiSlice(regex.FindAllString(input, 2))
	if !ok {
		return 0, 0, "", ""
	}

	regex, _ = regexp.Compile("[a-zA-z]")
	focus := regex.FindString(input)

	regex, _ = regexp.Compile("[a-zA-z]{2,}")
	data := regex.FindString(input)

	return constraint[0], constraint[1], focus, data
}

func fileReader() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	counter := 0
	for {
		b, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		x, y, focus, data := matchRegex(string(b))

		counter = counter + followsRules(x, y, focus, data)

		if err != nil {
			break
		}
	}

	fmt.Println(counter)
}

func main() {
	fileReader()
}
