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

	// divide all the lines in group of 3
	for i := 0; i < len(lines); i += 3 {
		for _, char := range lines[i] {
			if searchForCharacter(char, lines[i+1], lines[i+2]) {
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

func searchForCharacter(char int32, lines ...string) bool {
	for _, l := range lines {
		if !strings.ContainsRune(l, char) {
			return false
		}
	}

	return true
}
