package main

type SupplyStacks struct {
	stacks [][]string
	moves []Move
}

type Move struct {
	quantity int
	from int
	to int
}