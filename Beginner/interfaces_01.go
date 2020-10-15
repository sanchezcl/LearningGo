package main

import "fmt"

func printAnything(a interface{}) {
	fmt.Printf("Print a %T: %v\n", a, a)
}

func main() {
	printAnything("string")
	printAnything(666)
	printAnything(666.999)
	printAnything([...]int{2, 3, 5, 7, 11, 13})
	printAnything(map[string]int{"TutorialEdge": 2240, "": 6580350, "FunFunFunction": 171220,})
}
