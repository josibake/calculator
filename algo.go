package main

import (
	"math"
	"strconv"
	"strings"
)

var ops = map[string]struct {
	prec   int
	rAssoc bool
}{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

func IsParentheses(tok string) bool {
	switch tok {
	case "(",
		")":
		return true
	}
	return false
}

func CmdLineInputParsing(input string, ops map[string]struct {
	prec   int
	rAssoc bool
}) []string {
	var output []string
	input = strings.Replace(input, " ", "", -1)
	i := 0
	for j, token := range input {
		token := string(token)
		if _, exists := ops[token]; exists || IsParentheses(token) {
			if j == i {
				output = append(output, token)
			} else {
				output = append(output, input[i:j], token)
			}
			i = j + 1
		} else if _, exists := ops[string(token)]; !exists && j+1 == len(input) {
			output = append(output, input[i:])
		} else {
			continue
		}
	}
	return output
}

func ShuntingYardAlgorithm(input []string, ops map[string]struct {
	prec   int
	rAssoc bool
}) []string {
	var stack []string
	var rpn []string
	for _, tok := range input {
		if _, isOp := ops[tok]; isOp {
			for {
				if len(stack) == 0 {
					stack = append(stack, tok)
					break
				} else {
					prevOp := stack[len(stack)-1]
					if (ops[tok].prec < ops[prevOp].prec || (ops[tok].prec == ops[prevOp].prec && !ops[prevOp].rAssoc)) && prevOp != "(" {
						rpn = append(rpn, prevOp)
						stack = stack[:len(stack)-1]
					} else {
						stack = append(stack, tok)
						break
					}
				}
			}
		} else if tok == "(" {
			stack = append(stack, tok)
		} else if tok == ")" {
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
			rpn = append(rpn, tok)
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

func ComputeResult(rpn []string, ops map[string]struct {
	prec   int
	rAssoc bool
}) float64 {
	var result []float64
	for _, item := range rpn {
		if _, isOp := ops[item]; isOp {
			// pop y
			y := result[len(result)-1]
			result = result[:len(result)-1]
			// pop x
			x := result[len(result)-1]
			result = result[:len(result)-1]
			switch item {
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
			f, _ := strconv.ParseFloat(item, 64)
			result = append(result, f)
		}
	}
	return result[0]
}
