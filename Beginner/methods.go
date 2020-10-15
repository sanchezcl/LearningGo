package main

import "fmt"

type Employee struct {
	Name     string
	LastName string
	Wage     float32
}

func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) UpdateLastName(newLastName string) {
	e.LastName = newLastName
}

func (e *Employee) GetFullName() string {
	return e.Name + " " + e.LastName
}

func (e *Employee) incrementWageByPercentage(percentage float32) {
	e.Wage = e.Wage * (1 + (percentage / 100))
}

func main() {
	employee := Employee{
		Name:     "Jhon",
		LastName: "Doe",
		Wage:     115000,
	}
	fmt.Printf("Current fullname: %v \n", employee.GetFullName())
	fmt.Printf("Current Wage: %.2f \n", employee.Wage)
	employee.UpdateName("Jane")
	employee.UpdateLastName("D'oe")
	employee.incrementWageByPercentage(40)
	fmt.Println("****** Changes in employee ******")
	fmt.Printf("Current fullname: %v \n", employee.GetFullName())
	fmt.Printf("Current Wage: %.2f \n", employee.Wage)
}
