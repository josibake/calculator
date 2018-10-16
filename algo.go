package main

import "strconv"

var ops = map[string]struct {
	prec   int
	rAssoc bool
}{
	"(": {5, false},
	")": {5, true},
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

func CmdLineInputParsing(input string, ops map[string]struct {
	prec   int
	rAssoc bool
}) []string {
	var output []string
	i := 0
	for j, token := range input {
		if _, exists := ops[string(token)]; exists {
			number := input[i:j]
			if j == i {
				output = append(output, string(token))
			} else {
				output = append(output, number, string(token))
			}
			i = j + 1
		} else if _, exists := ops[string(token)]; !exists && j+1 == len(input) {
			number := input[i:]
			output = append(output, number)
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
			if len(stack) > 0 {
				prevOp := stack[len(stack)-1]
				if ops[tok].prec > ops[prevOp].prec {
					stack = append(stack)
				} else {
					rpn = append(rpn, tok)
				}
			}
			stack = append(stack, tok)
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
			a := result[len(result)-1]
			result = result[:len(result)-1]
			b := result[len(result)-1]
			result = result[:len(result)-1]
			if item == "+" {
				a += b
				result = append(result, a)
			} else {
				a *= b
				result = append(result, a)
			}
		} else {
			f, _ := strconv.ParseFloat(item, 64)
			result = append(result, f)
		}
	}
	return result[0]
}
