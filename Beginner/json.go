package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Sales int `json:"book_sales"`
	Age int `json:"age"`
	Developer bool `json:"is_developer"`
}

type Book struct {
	Title string `json:"title"`
	Author Author
}

func main() {
	fmt.Println("JSON Marshal: ")
	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	byteArray, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	jsonString := string(byteArray)
	fmt.Println(jsonString)

	fmt.Println("JSON unMarshal: ")
	var book2 Book
	err = json.Unmarshal([]byte(jsonString), &book2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Type: %T\n", book2)
	fmt.Printf("Content: %+v\n", book2)

	fmt.Println("JSON unstructured: ")
	str := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var reading map[string]interface{}
	err = json.Unmarshal([]byte(str), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
