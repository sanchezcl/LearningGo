package main

import (
	"fmt"
)

type pizza interface {
	getPrice() int
}

type basicPizza struct {
}

func (p *basicPizza) getPrice() int {
	return 10
}

//**** veggeMania ****
type veggeMania struct {
	pizza pizza
}

func (p *veggeMania) getPrice() int {
	return p.pizza.getPrice() + 5
}

//**** end veggeMania ****

//**** tomatoTopping ****
type tomatoTopping struct {
	pizza pizza
}

func (p *tomatoTopping) getPrice() int {
	return p.pizza.getPrice() + 7
}

//**** end tomatoTopping ****

//**** cheeseTopping ****
type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

//**** end cheeseTopping ****

func main() {
	fmt.Println("### Pizza Decorator ###")
	pizza := &basicPizza{}

	pizzaVeggie := &veggeMania{
		pizza: pizza,
	}
	//Add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizzaVeggie,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())

}
