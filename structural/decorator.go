package main

import "fmt"

type pizza interface {
	getPrice() int
}

type basePizza struct{}

func (p *basePizza) getPrice() int {
	return 10
}

type mushroomTopping struct {
	pizza pizza
}

func (m *mushroomTopping) getPrice() int {
	return m.pizza.getPrice() + 2
}

type cornTopping struct {
	pizza pizza
}

func (c *cornTopping) getPrice() int {
	return c.pizza.getPrice() + 5
}

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 10
}

func main() {
	pizza := &basePizza{}

	pizzaWitchCheese := &cheeseTopping{pizza}
	pizzaWitchCheeseAndCorn := &cornTopping{pizzaWitchCheese}
	pizzaWithMushroomAndCorn := &mushroomTopping{&cornTopping{pizza}}
	pizzaWitchCheeseMushroomAndCorn := &mushroomTopping{pizzaWitchCheeseAndCorn}

	fmt.Printf("Price of base pizza is $%d\n", pizza.getPrice())
	fmt.Printf("Price of base pizza with cheese topping is $%d\n", pizzaWitchCheese.getPrice())
	fmt.Printf("Price of base pizza with corn and cheese topping is $%d\n", pizzaWitchCheeseAndCorn.getPrice())
	fmt.Printf("Price of base pizza with corn and mushrooms topping is $%d\n", pizzaWithMushroomAndCorn.getPrice())
	fmt.Printf("Price of base pizza with corn, mushrooms and cheese topping is $%d\n", pizzaWitchCheeseMushroomAndCorn.getPrice())
}
