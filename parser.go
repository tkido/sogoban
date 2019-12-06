package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Parser is level parser
type Parser struct {
	Map map[rune]int
}

// NewParser returns new Parser
func NewParser() *Parser {
	m := map[rune]int{
		' ': FLOOR,
		'.': GOAL,
		'@': MAN,
		'+': MONG,
		'$': BAG,
		'*': BONG,
		'#': WALL,
	}
	return &Parser{
		Map: m,
	}
}

// Parse returns Level
func (p *Parser) Parse(url string) (lv *Level, err error) {
	lines, err := p.Readlines(url)
	if err != nil {
		return lv, err
	}
	log.Println(lines)
	width := 0
	for _, line := range lines {
		w := len(line)
		if width < w {
			width = w
		}
	}
	height := len(lines)

	board := make([][]int, 0)
	for _, line := range lines {
		row := make([]int, 0)
		for _, r := range line {
			i, ok := p.Map[r]
			if !ok {
				return nil, fmt.Errorf("unknown character")
			}
			row = append(row, i)
		}
		for len(row) < width {
			row = append(row, BLANK)
		}
		board = append(board, row)
	}

	lv = &Level{
		ID:     url,
		Board:  board,
		Width:  width,
		Height: height,
	}
	return lv, nil
}

// Readlines read text file and return lines
func (p *Parser) Readlines(path string) (lines []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err = s.Err(); err != nil {
		return lines, err
	}
	return lines, err
}

// Level is sokoban level
type Level struct {
	ID string
	// Name        string
	// Description string
	Board  [][]int
	Width  int
	Height int
}
