package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/MrShanks/pizza-delivery/pkg/restaurant"
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
		Buckets: []float64{8, 16, 24, 32},
	})
)

var orderQueue = make(chan restaurant.Order, 10)

func preparationHandler(w http.ResponseWriter, r *http.Request) {
	ord := restaurant.NewOrder()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(body, &ord)
	log.Printf("Order #%d added to the queue, %+v your order is going to be ready soon", ord.OrderID, ord.Pizzas)

	orderQueue <- ord
	orderQueueGauge.Inc()
}
func Preparing(ord restaurant.Order, ingredient string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	preparingTime := time.Duration(r.Intn(4000)+1000) * time.Millisecond
	time.Sleep(preparingTime)
	log.Printf("#%d Preparing... Adding %q. Time elapsed: %v", ord.OrderID, ingredient, preparingTime)
}

func Baking(ord restaurant.Order, pizzaName string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bakingTime := time.Duration(r.Intn(5000)+1000) * time.Millisecond
	time.Sleep(bakingTime)
	log.Printf("#%d Baking... Time elapsed: %v", ord.OrderID, bakingTime)
	time.Sleep(100 * time.Millisecond)
	log.Printf("#%d Pizza %s is ready!", ord.OrderID, pizzaName)
}

func preparePizza(ord restaurant.Order) {
	log.Printf("Starting order #%d", ord.OrderID)
	for _, pizza := range ord.Pizzas {
		start := time.Now()
		for _, ingredient := range pizza.Ings {
			Preparing(ord, ingredient)
		}
		Baking(ord, pizza.Name)
		pizzaTimeHistogram.Observe(time.Since(start).Seconds())
		pizzaMadeCounter.Inc()
	}

}

func StartKitchen(orderQueue chan restaurant.Order) {
	semaphore := make(chan struct{}, 5)
	var wg sync.WaitGroup
	for ord := range orderQueue {
		wg.Add(1)
		go func(ord restaurant.Order) {
			defer wg.Done()
			semaphore <- struct{}{}
			orderQueueGauge.Dec()
			activeChefGauge.Inc()
			preparePizza(ord)
			<-semaphore
			activeChefGauge.Dec()
		}(ord)
	}
	wg.Wait()
}

func main() {

	http.HandleFunc("/kitchen", preparationHandler)
	http.Handle("/metrics", promhttp.Handler())
	go StartKitchen(orderQueue)
	http.ListenAndServe(":3010", nil)
}
