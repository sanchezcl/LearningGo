package main

import "fmt"

func main() {
	days_array := [7]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	fmt.Println(days_array[0], days_array[5])

	var number_array [4]int
	fmt.Printf("Array vacio: %v \n", number_array)
	number_array[0] = 50
	number_array[1] = 23
	number_array[2] = 666
	number_array[3] = 69
	fmt.Printf("Array lleno: %v \n", number_array)

	fmt.Printf("Slice from array: %v \n", number_array[1:3])

	mapExample1 := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}
	fmt.Printf("this a map %v \n", mapExample1)
	fmt.Printf("Only a map item %v \n", mapExample1["MKBHD"])

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Printf("This is a map declared with map() function %v\n", m)
	fmt.Printf("This is a leng of a map %v \n", len(m))
}
