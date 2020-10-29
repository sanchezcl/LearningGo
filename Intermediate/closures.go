package main

import (
	"fmt"
)

func getLimit()  func() int {
	limit := 10
	return func() int {
		limit -= 1
		return limit
	}
}

func main() {
	limit := getLimit()
	fmt.Printf("limit: %v\n", limit())
	fmt.Printf("limit: %v\n", limit())
	fmt.Printf("limit: %v\n", limit())

	limit2 := getLimit()
	fmt.Printf("limit2: %v\n", limit2())
}
