package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var max int

	// Find the biggest value
	for _, value := range readFile() {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}

func readFile() map[int]int {
	var i int
	a := make(map[int]int)

	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)

	for _, l := range strings.Split(string(bytes), "\n") {
		if "" == strings.TrimSpace(l) {
			i++
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			n = 0
		}

		a[i] += n
	}

	return a
}
