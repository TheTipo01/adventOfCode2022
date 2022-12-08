package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var counter uint

	pairs := readFile()
	for _, p := range pairs {
		if checkLine(p) {
			counter++
		}
	}

	fmt.Println(counter)
}

// Checks whatever the pair of elfs have overlapping assigned sections
func checkLine(p Pair) bool {
	return overlap(p.elf1, p.elf2) || overlap(p.elf2, p.elf1)
}

// Checks if the the first elf overlaps the second, even partially
func overlap(elf1, elf2 Section) bool {
	return elf1.start <= elf2.start && elf1.end >= elf2.start || elf1.start >= elf2.start && elf1.start <= elf2.end
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