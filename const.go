package main

// consts
const (
	FLOOR = 0
	GOAL  = 1
	MAN   = 2
	AVOID = 4
	BAG   = 8
	WALL  = 16
	EXTRA = 32
	BLANK = 64

	MONG = MAN | GOAL
	MONA = MAN | AVOID
	BONG = BAG | GOAL
)
