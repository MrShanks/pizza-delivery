package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrShanks/pizza-delivery/Restaurant/order"
	"github.com/MrShanks/pizza-delivery/Restaurant/waiter"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Pizza-Delivery</h1><div>The fastest Pizzeria on the web 3.</div>")
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/menu.html")
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Order Placed")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/menu", menuHandler)
	http.HandleFunc("/order", orderHandler)

	// Serve static files (CSS, JavaScript, images, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	order1 := order.NewOrder()

	waiter.AddPizza(&order1, order.Margherita())
	waiter.AddPizza(&order1, order.Capricciosa())
	waiter.AddPizza(&order1, order.Diavola())
	waiter.SendOrder(order1)

	order2 := order.NewOrder()
	waiter.AddPizza(&order2, order.Margherita())
	waiter.AddPizza(&order2, order.Capricciosa())
	waiter.AddPizza(&order2, order.Diavola())
	waiter.SendOrder(order2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
