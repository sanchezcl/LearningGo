package main

import "fmt"

func increment(a int) {
	a += 2
	fmt.Printf("incrementa 'a': %v\n", a)
}

func incrementWithPointers(p *int) {
	*p += 2 //p es el puntero de la variable con '*' conseguimos asignar valores nuevos desde el puntero
	fmt.Printf("incrementa 'a': %v\n", *p) //para imprimir el valor de la variable con el puntero se agrega '*'
}

func main() {
	a := 1
	fmt.Printf("'a' antes de la funcion 'increment': %v\n", a)
	increment(a)
	fmt.Printf("'a' despues de la funcion 'increment': %v\n", a)
	fmt.Println("Se pierde el cambio por el scope de la funcion")
	fmt.Println("************************************************")
	fmt.Printf("'a' antes de la funcion 'increment': %v\n", a)
	fmt.Printf("puntero de 'a': %v\n", &a) //&a obtiene el puntero de la variable a
	incrementWithPointers(&a)
	fmt.Printf("'a' despues de la funcion 'increment': %v\n", a)
}
