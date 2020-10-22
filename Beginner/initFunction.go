package main

import "fmt"

var name string

func init()  {
	fmt.Println("This run first! asing 'Eliot' to name")
	name = "Elliot"
}

func main() {
	fmt.Println("This run second")
	fmt.Printf("Name: %s\n", name)
}
