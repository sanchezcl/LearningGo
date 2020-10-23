package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4,5,2,6,7,3,1,0,9}
	fmt.Println("Array of int unsorted")
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println("Array of int sorted")
	fmt.Println(a)

	b := []string{
		"lip",
		"salt",
		"ignore",
		"ultra",
		"switch",
		"hang",
		"tow",
		"bow",
		"well-off",
		"motionless",
		"quicksand",
		"abnormal",
	}
	fmt.Println("Array of string unsorted")
	fmt.Println(b)
	sort.Strings(b)
	fmt.Println("Array of string sorted")
	fmt.Println(b)
}
