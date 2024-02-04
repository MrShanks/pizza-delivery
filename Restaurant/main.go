package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MrShanks/pizza-delivery/Restaurant/order"
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

	ord := order.NewOrder()

	ord.AddPizza(order.Margherita())
	ord.AddPizza(order.Capricciosa())

	ordJSON, err := json.MarshalIndent(ord, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(ordJSON))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
