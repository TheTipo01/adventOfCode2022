package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var score uint
	lines := readFile()

	for _, line := range lines {
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		for _, char := range firstHalf {
			if searchForCharacter(string(char), secondHalf) {
				score += computeScore(char)
				break
			}
		}
	}

	fmt.Println(score)
}

func computeScore(char int32) uint {
	if char >= 97 && char <= 122 {
		return uint(char - 96)
	}

	return uint(char - 38)
}

func readFile() []string {
	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)

	return strings.Split(string(bytes), "\n")
}

func searchForCharacter(char, line string) bool {
	for _, c := range line {
		if string(c) == char {
			return true
		}
	}
	return false
}
