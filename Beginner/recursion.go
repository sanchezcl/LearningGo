package main

import "fmt"

func factorial(n int)int  {
	fmt.Printf("%v ", n)
	if n <= 1 {
		return 1
	}
	return n * factorial(n -1)
}

func fibonacci(n int)int  {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
	x := 10
	fmt.Printf("Factorial of %v is %v \n", x, factorial(x))
	fmt.Printf("el elemento %v de la serie de fibonacci es: %v \n", x, fibonacci(x))
}
