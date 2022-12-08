package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input := readFile()

	fmt.Printf("Part 1 found at %d\n", startOfMessage(input, 4))
	fmt.Printf("Part 2 found at %d\n", startOfMessage(input, 14))
}

func startOfMessage(input string, size int) int {
	for i := range input {
		if marker(input[i : i+size]) {
			return i + size
		}
	}

	return -1
}

func marker(s string) bool {
	mappa := make(map[rune]bool)

	for _, c := range s {
		if _, ok := mappa[c]; ok {
			return false
		} else {
			mappa[c] = true
		}
	}

	return true
}

func readFile() string {
	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)

	return string(bytes)
}
