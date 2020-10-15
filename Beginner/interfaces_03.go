package main

import "fmt"

type Worker interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Slave struct {
	Name string
}

func (s Slave) Age() int {
	panic("implement me")
}

func (s Slave) Random() (string, error) {
	panic("implement me")
}

func (s Slave) Language() string {
	return s.Name + " programs in Go"
}

func main() {
	// This will throw an error
	var programmers []Worker
	elliot := Slave{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
	fmt.Println(programmers)
}
