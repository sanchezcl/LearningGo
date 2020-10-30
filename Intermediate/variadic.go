package main

import (
	"fmt"
	"strings"
)

func aVaraidcFunction(args ...string)  {
	var result string
	for _, k := range args {
		result = fmt.Sprintf("%v %v", result, k)
	}
	fmt.Println(strings.Trim(result, " "))
}

func main() {
	aVaraidcFunction("This", "is", "a", "variadic", "func")
}
