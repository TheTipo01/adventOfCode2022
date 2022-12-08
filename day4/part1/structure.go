package main

type Section struct {
	start int
	end   int
}

type Pair struct {
	elf1 Section
	elf2 Section
}