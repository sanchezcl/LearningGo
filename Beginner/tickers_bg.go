package main

import (
	"fmt"
	"time"
)

func tickerBg() {
	ticker := time.NewTicker(1 * time.Second)

	for _ = range ticker.C {
		fmt.Println("tock")
	}
}

func main() {
	fmt.Println("Go Tickers Tutorial")

	go tickerBg()

	fmt.Println("The rest of my application can continue")

	// here we use an empty select{} in order to keep our main function alive indefinitely as it would
	// complete before our backgroundTask has a chance to execute if we didn't.
	select{}
}
