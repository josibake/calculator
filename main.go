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

func main() {
	// collapse inputs into a single string
	input := strings.Join(os.Args[1:], "")

	// determine type of operation
	switch {
	case strings.Contains(input, "+"):
		fmt.Println(Addition(Parse(input, "+")))
	case strings.Contains(input, "*"):
		fmt.Println(Multiply(Parse(input, "*")))
	case strings.Contains(input, "-"):
		fmt.Println(Subtract(Parse(input, "-")))
	case strings.Contains(input, "/"):
		fmt.Println(Divide(Parse(input, "/")))
	}
}
