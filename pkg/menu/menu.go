package menu

import "github.com/MrShanks/pizza-delivery/pkg/restaurant"

func Margherita() restaurant.Pizza {
	return restaurant.Pizza{
		Name: "Margherita",
		Cost: 15,
		Ings: []string{"Mozarella", "Tomato Sauce"},
	}
}

func Capricciosa() restaurant.Pizza {
	return restaurant.Pizza{
		Name: "Capricciosa",
		Cost: 21,
		Ings: []string{"Mozarella", "Tomato Sauce", "Mushrooms", "Prosciutto"},
	}
}

func Diavola() restaurant.Pizza {
	return restaurant.Pizza{
		Name: "Margherita",
		Cost: 21,
		Ings: []string{"Mozarella", "Tomato Sauce", "Chily", "Pepperoni"},
	}
}

func Calabra() restaurant.Pizza {
	return restaurant.Pizza{
		Name: "Margherita",
		Cost: 21,
		Ings: []string{"Mozarella", "Tomato Sauce", "Chily", "Pepperoni", "Nduia", "Provola", "Garlic"},
	}
}
