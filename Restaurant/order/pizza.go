package order

import "fmt"

var OrderCount int

type Pizza struct {
	Cost        float64
	Ingredients []string
}

type Order struct {
	Pizzas  []Pizza
	OrderID int
}

func NewPizza() Pizza {
	return Pizza{Cost: 5.0}
}

func NewOrder() Order {
	OrderCount++
	return Order{OrderID: OrderCount}
}

func (p *Pizza) AddTopping(ingredient string) {
	p.Ingredients = append(p.Ingredients, ingredient)
	p.Cost += 2.5
}

func (p *Pizza) ListIngredients() {
	for _, ingredient := range p.Ingredients {
		fmt.Println(ingredient)
	}
	fmt.Println(p.Cost)
}
