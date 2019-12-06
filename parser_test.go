package main

import "testing"

func TestParse(t *testing.T) {
	parser := NewParser()
	lv, err := parser.Parse("levels/simple/map1")
	if err != nil {
		t.Error(err)
	}
	t.Log(lv)

	printer := NewPrinter()
	t.Log(printer.Print(lv))
}
