package main

import (
	"github.com/MrShanks/pizza-delivery/pkg/menu"
	"github.com/MrShanks/pizza-delivery/pkg/restaurant"
)

func main() {

	order1 := restaurant.NewOrder()

	restaurant.AddPizza(&order1, menu.Margherita())
	restaurant.AddPizza(&order1, menu.Capricciosa())
	restaurant.AddPizza(&order1, menu.Capricciosa())
	restaurant.AddPizza(&order1, menu.Capricciosa())
	restaurant.AddPizza(&order1, menu.Diavola())
	restaurant.AddPizza(&order1, menu.Diavola())

	restaurant.SendOrder(order1)

	order2 := restaurant.NewOrder()

	restaurant.AddPizza(&order2, menu.Margherita())
	restaurant.AddPizza(&order2, menu.Capricciosa())
	restaurant.AddPizza(&order2, menu.Calabra())

	restaurant.SendOrder(order2)

	order3 := restaurant.NewOrder()

	restaurant.AddPizza(&order3, menu.Margherita())
	restaurant.AddPizza(&order3, menu.Capricciosa())
	restaurant.AddPizza(&order3, menu.Calabra())
	restaurant.AddPizza(&order3, menu.Calabra())
	restaurant.AddPizza(&order3, menu.Calabra())
	restaurant.AddPizza(&order3, menu.Calabra())
	restaurant.AddPizza(&order3, menu.Calabra())
	restaurant.AddPizza(&order3, menu.Calabra())

	restaurant.SendOrder(order3)
}
