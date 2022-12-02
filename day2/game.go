package main

func whatWins(c uint) uint {
	switch c {
	case rock:
		return paper
	case paper:
		return scissor
	case scissor:
		return rock
	}

	panic("invalid value")
}

func whatLoses(c uint) uint {
	switch c {
	case rock:
		return scissor
	case paper:
		return rock
	case scissor:
		return paper
	}

	panic("invalid value")
}
