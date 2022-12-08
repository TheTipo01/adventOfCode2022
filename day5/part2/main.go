package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ss := readFile()

	for _, m := range ss.moves {
		// take ss.moves.quantity from ss.moves.from
		crate := ss.stacks[m.from-1][len(ss.stacks[m.from-1])-m.quantity:]
		ss.stacks[m.from-1] = ss.stacks[m.from-1][:len(ss.stacks[m.from-1])-m.quantity]

		ss.stacks[m.to-1] = append(ss.stacks[m.to-1], crate...)
	}

	for _, s := range ss.stacks {
		fmt.Print(s[len(s)-1])
	}
}

func readFile() SupplyStacks {
	var ss SupplyStacks

	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)

	lines := strings.Split(string(bytes), "\n")

	var emptyline int
	// Search for the empty line delimiting the crates and the moves
	for i, line := range lines {
		if line == "" {
			emptyline = i
			break
		}
	}

	// Parse the moves
	for _, l := range lines[emptyline+1:] {
		var m Move
		// move 7 from 3 to 9
		_, _ = fmt.Sscanf(l, "move %d from %d to %d", &m.quantity, &m.from, &m.to)
		ss.moves = append(ss.moves, m)
	}

	// Parse the number of crates
	splitted := strings.Split(strings.TrimSpace(lines[emptyline-1]), "   ")
	crates := len(splitted)

	// Parse the stacks
	ss.stacks = make([][]string, crates)
	for i := emptyline - 2; i >= 0; i-- {
		padding := 1
		for j := 0; j <= crates - 1; j++ {
			if len(lines[i]) > padding {
				crate := strings.TrimSpace(string(lines[i][padding]))
				if crate != "" {
					ss.stacks[j] = append(ss.stacks[j], crate)
				}

				padding += 4
			}
		}
	}

	return ss
}