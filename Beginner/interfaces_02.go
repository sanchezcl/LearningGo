package main

import "fmt"

type Car interface {
	TurnOn()
	TurnOff()
}

type OffRoad struct {
	Name   string
}

func (o OffRoad) TurnOn() {
	fmt.Printf("%v is turned on \n", o.Name)
}

func (o OffRoad) TurnOff() {
	fmt.Printf("%v is turned off \n", o.Name)
}

type SportCar struct {
	Name   string
	Status bool
}

func (s SportCar) TurnOn() {
	fmt.Printf("%v is turned on \n", s.Name)
}

func (s SportCar) TurnOff() {
	fmt.Printf("%v is turned off \n", s.Name)
}

func main() {
	var offRoad OffRoad
	offRoad.Name = "Land Cruiser"

	var sport SportCar
	sport.Name = "Mustang GT"

	offRoad.TurnOn()
	sport.TurnOn()


	var cars []Car
	cars = append(cars, offRoad)
	cars = append(cars, sport)

	fmt.Println(cars)
}
