package main

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func wrapFunc(f func()) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	f()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

func main() {
	fmt.Printf("Type: %T\n", myFunc)
	wrapFunc(myFunc)
}
