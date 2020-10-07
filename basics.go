package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	aString := "this is a string"

	fmt.Println(aString)
	var anotherString string

	anotherString = "another string"
	fmt.Println(anotherString)
	fmt.Printf("Printed with format: %v \n", aString)
	fmt.Printf("Printed with format and concatenated: %v | %v \n", aString, anotherString)

	var aInt int = 5
	var aFloat float32 = 5.3333

	fmt.Printf("This is a %T: %v and this is a %T: %v \n", aInt, aInt, aFloat, aFloat)
}
