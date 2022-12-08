package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	trees := readFile()

	var max int

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			current := check(trees, i, j)
			if current > max {
				max = current
			}
		}
	}

	fmt.Println(max)
}

func readFile() (numbers [][]int) {
	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)
	lines := strings.Split(string(bytes), "\n")

	// Why did I waste so much time on optimizing the numbers of memory allocations?
	numbers = make([][]int, len(lines))
	for i, l := range lines {
		numbers[i] = make([]int, len(l))
		for j, c := range l {
			numbers[i][j] = int(c) - 48
		}
	}

	return numbers
}
