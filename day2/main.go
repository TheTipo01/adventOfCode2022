package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	rounds := readFile()

	// Part 1
	var score uint

	for _, r := range rounds {
		score += computeScore(r)
	}

	fmt.Printf("Part 1: %d\n", score)

	// Part 2
	score = 0
	for _, r := range rounds {
		switch r.mine {
		case lose:
			r.mine = whatLoses(r.enemy)
		case draw:
			r.mine = r.enemy
		case win:
			r.mine = whatWins(r.enemy)
		}

		score += computeScore(r)
	}

	fmt.Printf("Part 2: %d\n", score)
}

func wonRound(r Round) bool {
	return r.mine == whatWins(r.enemy)
}

func readFile() []Round {
	var a []Round
	f, _ := os.Open("input.txt")
	bytes, _ := io.ReadAll(f)

	for _, l := range strings.Split(string(bytes), "\n") {
		splitted := strings.Split(l, " ")
		if len(splitted) == 2 {
			a = append(a, computeRound(splitted))
		}
	}

	return a
}

func computeRound(split []string) (r Round) {
	switch split[0] {
	case "A":
		r.enemy = rock
	case "B":
		r.enemy = paper
	case "C":
		r.enemy = scissor
	default:
		panic("what")
	}

	switch split[1] {
	case "X":
		r.mine = rock
	case "Y":
		r.mine = paper
	case "Z":
		r.mine = scissor
	default:
		panic("what")
	}

	return r
}

func computeScore(r Round) (score uint) {
	score += r.mine + 1

	if wonRound(r) {
		score += 6
	} else {
		if r.mine == r.enemy {
			score += 3
		}
	}

	return score
}
