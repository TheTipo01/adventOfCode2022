package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var counter uint

	for _, p := range readFile() {
		if checkLine(p) {
			counter++
		}
	}

	fmt.Println(counter)
}

// Checks whatever the pair of elfs have overlapping assigned sections
func checkLine(p Pair) bool {
	if overlap(p) {
		return true
	} else {
		p.elf2, p.elf1 = p.elf1, p.elf2
		return overlap(p)
	}
}

// Checks if the the first elf overlaps the second
func overlap(p Pair) bool {
	return p.elf2.start >= p.elf1.start && p.elf2.end <= p.elf1.end
}

func readFile() []Pair {
	var sections []Pair
	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)

	for _, l := range strings.Split(string(bytes), "\n") {
		var p Pair
		fmt.Sscanf(l, "%d-%d,%d-%d", &p.elf1.start, &p.elf1.end, &p.elf2.start, &p.elf2.end)

		sections = append(sections, p)
	}

	return sections
}