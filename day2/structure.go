package main

const (
	rock = iota
	paper
	scissor
)

const (
	lose = iota
	draw
	win
)

type Round struct {
	mine  uint
	enemy uint
}
