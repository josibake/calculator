package main

import (
	"fmt"
	"github.com/josibake/calculator"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], "")
	fmt.Printf("%v: %v\n", input, calculator.Calculate(input))
}
