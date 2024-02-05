package restaurant

var OrderCount int

type Pizza struct {
	Name string
	Cost float64
	Ings []string
}

type Order struct {
	OrderID int
	Pizzas  []Pizza
}

func (p *Pizza) Add(ingredient string) {
	p.Ings = append(p.Ings, ingredient)
}

func NewOrder() Order {
	OrderCount++
	return Order{OrderID: OrderCount}
}

func Margherita() Pizza {
	return Pizza{
		Name: "Margherita",
		Cost: 15,
		Ings: []string{"Mozarella", "Tomato Sauce"},
	}
}

func Capricciosa() Pizza {
	return Pizza{
		Name: "Capricciosa",
		Cost: 21,
		Ings: []string{"Mozarella", "Tomato Sauce", "Mushrooms", "Prosciutto"},
	}
}

func Diavola() Pizza {
	return Pizza{
		Name: "Margherita",
		Cost: 21,
		Ings: []string{"Mozarella", "Tomato Sauce", "Chily", "Pepperoni"},
	}
}
