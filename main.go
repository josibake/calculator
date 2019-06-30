package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], "")
	tokens := CmdLineInputParsing(input)
	rpn := ShuntingYardAlgorithm(tokens)
	result := ComputeResult(rpn)
	fmt.Println(result)
}
