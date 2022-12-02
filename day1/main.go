package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	a := readFile()
	sort.Ints(a)

	fmt.Printf("Part 1: %d\n", a[len(a)-1])
	fmt.Printf("Part 2: %d\n", a[len(a)-1]+a[len(a)-2]+a[len(a)-3])
}

func readFile() []int {
	var cont int
	var a []int

	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)

	for _, l := range strings.Split(string(bytes), "\n") {
		if "" == strings.TrimSpace(l) {
			a = append(a, cont)
			cont = 0
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			n = 0
		}

		cont += n
	}

	return a
}
