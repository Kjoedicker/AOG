package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLine() []string {
	route := []string{}

	path := "input.txt"
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		route = append(route, scanner.Text())
	}

	return route
}

func main() {
	route := readLine()

	multi := 1
	param := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	for _, p := range param {

		counter := 0
		for x, y := 0, 0; x < len(route); x, y = x+p[0], y+p[1] {
			if y >= 31 {
				z := y - 31
				y = z
			}

			if string(route[x][y]) == "#" {
				counter = counter + 1
			}
		}
		multi = multi * counter
	}

	fmt.Println(multi)
}
