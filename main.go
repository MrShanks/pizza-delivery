package main

import (
	"github.com/MrShanks/pizza-delivery/pkg/restaurant"
)

func main() {

	order1 := restaurant.NewOrder()

	restaurant.AddPizza(&order1, restaurant.Margherita())
	restaurant.AddPizza(&order1, restaurant.Capricciosa())
	restaurant.AddPizza(&order1, restaurant.Capricciosa())
	restaurant.AddPizza(&order1, restaurant.Capricciosa())
	restaurant.AddPizza(&order1, restaurant.Diavola())

	restaurant.SendOrder(order1)

}
