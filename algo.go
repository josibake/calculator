package main

import (
	"math"
	"strconv"
	"strings"
)

var operators = map[string]struct {
	prec   int
	rAssoc bool
}{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

func isParentheses(token string) bool {
	switch token {
	case "(",
		")":
		return true
	}
	return false
}

func CmdLineInputParsing(input string) []string {
	var output []string
	input = strings.Replace(input, " ", "", -1)
	i := 0
	for j, token := range input {
		token := string(token)
		if _, exists := operators[token]; exists || isParentheses(token) {
			if j == i {
				output = append(output, token)
			} else {
				output = append(output, input[i:j], token)
			}
			i = j + 1
		} else if _, exists := operators[token]; !exists && j+1 == len(input) {
			output = append(output, input[i:])
		} else {
			continue
		}
	}
	return output
}

func ShuntingYardAlgorithm(input []string) []string {
	var stack []string
	var rpn []string
	for _, token := range input {
		if _, exists := operators[token]; exists {
			for {
				if len(stack) == 0 {
					stack = append(stack, token)
					break
				} else {
					prevOp := stack[len(stack)-1]
					if (operators[token].prec < operators[prevOp].prec || (operators[token].prec == operators[prevOp].prec && !operators[prevOp].rAssoc)) && prevOp != "(" {
						rpn = append(rpn, prevOp)
						stack = stack[:len(stack)-1]
					} else {
						stack = append(stack, token)
						break
					}
				}
			}
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for {
				prevOp := stack[len(stack)-1]
				if prevOp != "(" {
					rpn = append(rpn, prevOp)
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
					break
				}
			}
		} else {
			rpn = append(rpn, token)
		}
	}
	// drain the stack
	for len(stack) > 0 {
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rpn = append(rpn, op)
	}
	return rpn
}

func ComputeResult(rpn []string) float64 {
	var result []float64
	for _, token := range rpn {
		if _, exists := operators[token]; exists {
			// pop y
			y := result[len(result)-1]
			result = result[:len(result)-1]
			// pop x
			x := result[len(result)-1]
			result = result[:len(result)-1]
			switch token {
			case "+":
				x += y
				result = append(result, x)
			case "*":
				x *= y
				result = append(result, x)
			case "-":
				x -= y
				result = append(result, x)
			case "/":
				x = x / y
				result = append(result, x)
			case "^":
				x = math.Pow(x, y)
				result = append(result, x)
			}
		} else {
			f, _ := strconv.ParseFloat(token, 64)
			result = append(result, f)
		}
	}
	return result[0]
}
