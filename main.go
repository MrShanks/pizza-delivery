package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Pizza-Delivery</h1><div>The fastest Pizzeria on the web 3.</div>")
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/menu.html")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/menu", menuHandler)
	// Serve static files (CSS, JavaScript, images, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
