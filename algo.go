package main

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
	rpn := []string{"1", "2", "+"}
	return rpn
}

func ComputeResult(rpn []string, ops map[string]struct {
	prec   int
	rAssoc bool
}) float64 {
	return 42.0
}
