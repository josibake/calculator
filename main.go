package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Addition(x, y float64) float64 {
	return x + y
}

func Multiply(x, y float64) float64 {
	return x * y
}

func Subtract(x, y float64) float64 {
	return x - y
}

func Divide(x, y float64) float64 {
	return x / y
}

func Parse(input, sep string) (float64, float64) {
	ab := strings.Split(input, sep)
	a, _ := strconv.ParseFloat(ab[0], 64)
	b, _ := strconv.ParseFloat(ab[1], 64)
	return a, b
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func InputParsing(input string, operators map[string]int) []string {
	var output []string
	startFrom := 0
	for i, token := range input {
		if operators[string(token)] != 0 {
			number := input[startFrom:i]
			if i == startFrom {
				output = append(output, string(token))
			} else {
				output = append(output, number, string(token))
			}
			startFrom = i + 1
		} else if operators[string(token)] == 0 && i+1 == len(input) {
			number := input[startFrom:]
			output = append(output, number)
		} else {
			continue
		}
	}
	return output
}

func main() {
	// collapse inputs into a single string
	precedence := map[string]int{"*": 2, "/": 2, "+": 3, "-": 3, "(": 1, ")": 1}
	input := InputParsing(strings.Join(os.Args[1:], ""), precedence)
	fmt.Println(input)
}
