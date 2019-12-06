package main

import (
	"bytes"
)

// Printer is level printer
type Printer struct {
	Map map[int]rune
}

// NewPrinter returns new Printer
func NewPrinter() *Printer {
	m := map[int]rune{
		FLOOR: '　',
		GOAL:  '・',
		MAN:   '足',
		MONG:  '距',
		AVOID: '×',
		BAG:   '田',
		BONG:  '回',
		WALL:  '■',
		EXTRA: '△',
		BLANK: '　',
	}
	return &Printer{
		Map: m,
	}
}

// Print returns Level
func (p *Printer) Print(lv *Level) string {
	buf := bytes.Buffer{}
	board := lv.Board
	for _, row := range board {
		for _, cell := range row {
			r := p.Map[cell]
			buf.WriteRune(r)
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}
