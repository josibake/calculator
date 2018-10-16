package main

import (
	"reflect"
	"testing"
)

func TestAddition(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		n float64
	}{
		{0, 0, 0},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Addition(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%f+%f) was incorrect, got: %f, want: %f.", table.x, table.y, total, table.n)
		}
	}
}

func TestInputParsing(t *testing.T) {
	tables := []struct {
		i string
		o []string
	}{
		{"1+3*2", []string{"1", "+", "3", "*", "2"}},
		{"11+3*2", []string{"11", "+", "3", "*", "2"}},
		{"1+32*2", []string{"1", "+", "32", "*", "2"}},
		{"1+3*203", []string{"1", "+", "3", "*", "203"}},
		{"1.1+3*2", []string{"1.1", "+", "3", "*", "2"}},
		{"1.1+(3*2)", []string{"1.1", "+", "(", "3", "*", "2", ")"}},
		{"(1+1)*5+3", []string{"(", "1", "+", "1", ")", "*", "5", "+", "3"}},
	}

	for _, table := range tables {
		precedence := map[string]int{"*": 2, "/": 2, "+": 3, "-": 3, "(": 1, ")": 1}
		total := InputParsing(table.i, precedence)
		if reflect.DeepEqual(total, table.o) == false {
			t.Errorf("Parsing of %s was incorrect, got: %s, want: %s.", table.i, total, table.o)
		}
	}
}
