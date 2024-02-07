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

func NewOrder() Order {
	OrderCount++
	return Order{OrderID: OrderCount}
}
