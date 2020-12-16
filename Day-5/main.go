package main

import (
	"fmt"
	"os"
)

func findRow(a string, b string, length int, pis int, pos string, l int, u int) int {

	if pis == length {
		switch string(pos[pis-1]) {
		case a:
			// fmt.Println(l, u)
			return l
		case b:
			// fmt.Println(l, u)
			return u - 1
		}
	}

	value := (u - l) / 2
	switch string(pos[pis]) {
	case a:
		// fmt.Println(a, l, u-value)
		return findRow(a, b, length, pis+1, pos, l, u-value)
	case b:
		// fmt.Println(b, l+value, u)
		return findRow(a, b, length, pis+1, pos, l+value, u)
	}

	return 0
}

func findCol() {

}

func main() {
	//test case
	// value := []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL", "FBFBBFFRLR"}

	max := 0
	for _, v := range os.Args {
		a := findRow("F", "B", len(v)-3, 0, v, 0, 128)
		b := findRow("L", "R", 3, 0, v[7:], 0, 8)
		c := (a * 8) + b

		if c > max {
			max = c
		}

	}
	fmt.Println(max)
}
