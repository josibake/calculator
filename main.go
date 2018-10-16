package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], "")
	tokens := CmdLineInputParsing(input, ops)
	rpn := ShuntingYardAlgorithm(tokens, ops)
	result := ComputeResult(rpn, ops)
	fmt.Println(result)
}
