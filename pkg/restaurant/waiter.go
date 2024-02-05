package restaurant

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func AddPizza(ord *Order, pizza Pizza) {
	ord.Pizzas = append(ord.Pizzas, pizza)
}

func SendOrder(ord Order) {
	ordJSON, err := json.Marshal(ord)
	if err != nil {
		log.Fatal("Error marshalling JSON: ", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:3010/kitchen", bytes.NewBuffer(ordJSON))
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}

	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)
}
