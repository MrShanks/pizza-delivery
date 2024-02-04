package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/MrShanks/pizza-delivery/Restaurant/order"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	pizzaMadeCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "Restaurant_pizza_made_total",
		Help: "The total number of pizzas that the kitchen has made thus far",
	})
	activeChefGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "Restaurant_active_chefs",
		Help: "Number of actively working chefs in the kitchen",
	})
	orderQueueGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "Restaurant_order_queue_size",
		Help: "Number of orders in the queue",
	})
	pizzaTimeHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "Restaurant_pizza_making_duration_seconds",
		Help:    "The amount of time needed to make a pizza",
		Buckets: []float64{1, 2, 4, 8, 16},
	})
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
	orderQueueGauge.Inc()
}

func StartKitchen(orderQueue chan order.Order) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	semaphore := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for ord := range orderQueue {
		wg.Add(1)
		go func(ord order.Order) {
			defer wg.Done()
			semaphore <- struct{}{}
			orderQueueGauge.Dec()
			activeChefGauge.Inc()
			fmt.Printf("Preparing order number: #%d\n", ord.OrderID)
			for _, pizza := range ord.Pizzas {
				start := time.Now()
				for _, ingredient := range pizza.Ingredients {
					time.Sleep(time.Duration(r.Intn(3000)) * time.Millisecond)
					fmt.Printf("#%d: Adding: %q\n", ord.OrderID, ingredient)
				}
				time.Sleep(time.Duration(r.Intn(2000)) * time.Millisecond)
				fmt.Printf("#%d: Pizza is ready\n", ord.OrderID)
				pizzaTimeHistogram.Observe(time.Since(start).Seconds())
				fmt.Println(time.Since(start).Seconds())
				pizzaMadeCounter.Inc()

			}
			<-semaphore
			activeChefGauge.Dec()
		}(ord)
	}
	wg.Wait()
}

func main() {

	http.HandleFunc("/kitchen", handler)
	http.Handle("/metrics", promhttp.Handler())
	go StartKitchen(orderQueue)
	http.ListenAndServe(":3010", nil)
}
