package main

import "testing"

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
