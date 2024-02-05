package order

var OrderCount int

type Pizza struct {
	Name        string
	Cost        float64
	Ingredients []string
	Ing         Ing
}

type Order struct {
	OrderID int
	Pizzas  []Pizza
}

type Ing struct {
	Mozzarella  bool
	TomatoSauce bool
	Chily       bool
	Mushroom    bool
	Prosciutto  bool
}

func Add(options ...func(*Ing)) *Ing {
	ing := new(Ing)
	for _, opt := range options {
		opt(ing)
	}
	return ing
}

func Mozzarella() func(*Ing) {
	return func(i *Ing) {
		i.Mozzarella = true
	}
}

func TomatoSauce() func(*Ing) {
	return func(i *Ing) {
		i.TomatoSauce = true
	}
}

func Chily() func(*Ing) {
	return func(i *Ing) {
		i.Chily = true
	}
}

func Mushroom() func(*Ing) {
	return func(i *Ing) {
		i.Mushroom = true
	}
}

func Prosciutto() func(*Ing) {
	return func(i *Ing) {
		i.Prosciutto = true
	}
}

func NewOrder() Order {
	OrderCount++
	return Order{OrderID: OrderCount}
}

func Margherita() Pizza {
	return Pizza{
		Name: "Margherita",
		Ing: *Add(
			Mozzarella(),
			TomatoSauce())}
}

func Capricciosa() Pizza {
	return Pizza{
		Name: "Capricciosa",
		Ing: *Add(
			Mozzarella(),
			TomatoSauce(),
			Mushroom(),
			Prosciutto())}
}

func Diavola() Pizza {
	return Pizza{
		Name: "Diavola",
		Ing: *Add(
			Mozzarella(),
			TomatoSauce(),
			Chily())}
}
