package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/MrShanks/pizza-delivery/Restaurant/order"
)

var orderQueue = make(chan order.Order, 10)

func handler(w http.ResponseWriter, r *http.Request) {
	ord := order.NewOrder()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(body, &ord)

	orderQueue <- ord
}

func StartKitchen(orderQueue chan order.Order) {
	semaphore := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for ord := range orderQueue {
		wg.Add(1)
		go func(ord order.Order) {
			defer wg.Done()
			semaphore <- struct{}{}
			fmt.Printf("Preparing order number: #%d\n", ord.OrderID)
			for _, pizza := range ord.Pizzas {
				for _, ingredient := range pizza.Ingredients {
					time.Sleep(time.Second * 2)
					fmt.Printf("#%d: Adding: %q\n", ord.OrderID, ingredient)
				}
				time.Sleep(time.Second * 2)
				fmt.Printf("#%d: Pizza is ready\n", ord.OrderID)

			}
			<-semaphore
		}(ord)
	}
	wg.Wait()
}

func main() {

	http.HandleFunc("/kitchen", handler)
	go StartKitchen(orderQueue)
	http.ListenAndServe(":3000", nil)
}
