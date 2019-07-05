package calculator

import (
	"math"
	"reflect"
	"testing"
)

func TestCmdLineInputParsing(t *testing.T) {
	tables := []struct {
		i string
		o []string
	}{
		{"sin(0)", []string{"sin", "(", "0", ")"}},
		{"1 + 3*2", []string{"1", "+", "3", "*", "2"}},
		{"1 +  3 *2", []string{"1", "+", "3", "*", "2"}},
		{"1+3*2", []string{"1", "+", "3", "*", "2"}},
		{"11+3*2", []string{"11", "+", "3", "*", "2"}},
		{"1+32*2", []string{"1", "+", "32", "*", "2"}},
		{"1+3*203", []string{"1", "+", "3", "*", "203"}},
		{"1.1+3*2", []string{"1.1", "+", "3", "*", "2"}},
		{"1.1+(3*2)", []string{"1.1", "+", "(", "3", "*", "2", ")"}},
		{"(1+1)*5+3", []string{"(", "1", "+", "1", ")", "*", "5", "+", "3"}},
		{"3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3", []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "^", "2", "^", "3"}},
	}

	for _, table := range tables {
		total := CmdLineInputParsing(table.i)
		if reflect.DeepEqual(total, table.o) == false {
			t.Errorf("Parsing of %s was incorrect, got: %s, want: %s.", table.i, total, table.o)
		}
	}
}

func TestShuntingYardAlgorithm(t *testing.T) {
	tables := []struct {
		i []string
		o []string
	}{
		{[]string{"sin", "(", "0", ")"}, []string{"0", "sin"}},
		{[]string{"1", "+", "3", "+", "2"}, []string{"1", "3", "+", "2", "+"}},
		{[]string{"1", "+", "3", "*", "2"}, []string{"1", "3", "2", "*", "+"}},
		{[]string{"11", "+", "3", "*", "2"}, []string{"11", "3", "2", "*", "+"}},
		{[]string{"1", "+", "32", "*", "2"}, []string{"1", "32", "2", "*", "+"}},
		{[]string{"1", "+", "3", "*", "203"}, []string{"1", "3", "203", "*", "+"}},
		{[]string{"1.1", "+", "3", "*", "2"}, []string{"1.1", "3", "2", "*", "+"}},
		{[]string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "^", "2", "^", "3"}, []string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "^", "^", "/", "+"}},
	}

	for _, table := range tables {
		rpn := ShuntingYardAlgorithm(table.i)
		if reflect.DeepEqual(rpn, table.o) == false {
			t.Errorf("Parsing of %s was incorrect, got: %s, want: %s.", table.i, rpn, table.o)
		}
	}
}

func TestComputeResult(t *testing.T) {
	tables := []struct {
		i []string
		o float64
	}{
		{[]string{"0", "sin"}, 0},
		{[]string{"1", "2", "-"}, -1},
		{[]string{"1", "3", "+", "2", "+"}, 6},
		{[]string{"1", "3", "2", "*", "+"}, 7},
		{[]string{"11", "3", "2", "*", "+"}, 17},
		{[]string{"1", "32", "2", "*", "+"}, 65},
		{[]string{"1", "3", "203", "*", "+"}, 610},
		{[]string{"1.1", "3", "2", "*", "+"}, 7.1},
		{[]string{"2", "3", "^"}, 8},
		{[]string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "^", "^", "/", "+"}, 3.0001220703125},
	}

	for _, table := range tables {
		result := ComputeResult(table.i)
		if result != table.o {
			t.Errorf("Parsing of %s was incorrect, got: %f, want: %f.", table.i, result, table.o)
		}
	}
}

func TestCalculate(t *testing.T) {
	tables := []struct {
		i string
		o float64
	}{
		{"sin(1)", math.Sin(1)},
		{"sin(pi)", math.Sin(math.Pi)},
		{"cos(0)", math.Cos(0)},
		{"1+sin(1)", 1 + math.Sin(1)},
		{"1+sin(1+1)", 1 + math.Sin(1+1)},
		{"1+sin(sin(1+1))", 1 + math.Sin(math.Sin(1+1))},
		{"1+2*sin(1)", 1 + 2*math.Sin(1)},
		{"1 + 3*2", 7},
		{"1 - 2", -1},
		{"1+32*2", 65},
		{"1+3*203", 610},
		{"1.1+(3*2)", 7.1},
		{"2^3", 8},
		{"3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3", 3.0001220703125},
	}

	for _, table := range tables {
		result := Calculate(table.i)
		if result != table.o {
			t.Errorf("Parsing of %s was incorrect, got: %f, want: %f.", table.i, result, table.o)
		}
	}
}
