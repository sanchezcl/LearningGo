package main

import (
	"fmt"
	"log"
)

func fullName(name string, lastName string) (string, error) {
	return name + " " + lastName, nil
}
func main() {
	name := "jhon"
	lastName := "Doe"
	fullName, err := fullName(name, lastName)
	if err != nil {
		log.Fatal("ERROR")
	}
	fmt.Printf("Fullname: %v\n", fullName)

	func(message string) {
		fmt.Printf("This is a message from an anonymus function: %s \n", message)
	}("hello!")
}
