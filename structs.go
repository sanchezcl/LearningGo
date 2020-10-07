package main

import "fmt"

type Person struct {
	name     string
	lastName string
	age      int
}

type Team struct {
	name    string
	players [2]Person
}

func main() {
	var JDPerson Person
	JDPerson.name = "Jhon"
	JDPerson.lastName = "Doe"
	JDPerson.age = 33
	fmt.Printf("Person struct print: %v\n", JDPerson)
	fmt.Printf("Name: %v, Last name: %v, Age: %v \n", JDPerson.name, JDPerson.lastName, JDPerson.age)

	JnDPerson := Person{
		name:     "Jane",
		lastName: "Doe",
		age:      33,
	}
	team := [...]Person{JDPerson, JnDPerson}
	fmt.Println(team)

	fmt.Println("******Struct in a Struct*****")
	var doeTeam Team
	fmt.Printf("Empty struct: ")
	fmt.Println(doeTeam)

	doeTeam = Team{
		name:    "unkowns of the world",
		players: team,
	}
	fmt.Printf("Data from struct: ")
	fmt.Printf("Team name: %v \n", doeTeam.name)
	fmt.Printf("this is the team: %v \n", doeTeam.players)
	fmt.Println("*****************************")
}
