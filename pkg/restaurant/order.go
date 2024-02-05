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
