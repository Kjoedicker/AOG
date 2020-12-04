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
		// fmt.Print(string(b))
		x, y, focus, data := matchRegex(string(b))
		// fmt.Printf("%v\n%v\n%v\n%v\n", x, y, focus, data)

		counter = counter + followsRules(x, y, focus, data)

		if err != nil {
			break
		}
	}

	fmt.Println(counter)
}

// TODO:
func main() {
	fileReader()
	//Isolated instance
	// x, y, focus, data := matchRegex("19-20b: bbbbbbbbbbbbbbbbbhffdasfdasfdsdasfdgfbbb")
	// fmt.Printf("%v\n%v\n%v\n%v\n", x, y, focus, data)
	// if followsRules(x, y, focus, data) == 1 {
	// 	fmt.Println("true")
	// }

	// a := 1
	// fmt.Println(os.Args)
	// for a+2 <= (len(os.Args) + 1) {
	// 	fmt.Println(os.Args[a : a+3])

	// 	// x, y, focus, data := matchRegex(strings.Join(os.Args[a:a+3], ","))
	// 	// breaksRules(x, y, focus, data)

	// 	a = a + 3
	// 	// fmt.Println(a)

	// }

}
