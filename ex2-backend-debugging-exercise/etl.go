package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Product struct {
	ID    int    `json:"product_id"`
	Name  string `json:"product_name"`
	Price int    `json:"price"`
}

type Order struct {
	ID         int       `json:"order_id"`
	CustomerId int       `json:"customer"`
	Products   []Product `json:"products"`
}

type Customer struct {
	Id        int    `json:"customer_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func fetchAndUnmarshal[T any](url string) ([]T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []T
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func main() {
	orderUrl := os.Getenv("ORDER_SERVICE_URL")
	if orderUrl == "" {
		orderUrl = "http://localhost:3000/orders"
	}

	customerUrl := os.Getenv("CUSTOMER_SERVICE_URL")
	if customerUrl == "" {
		customerUrl = "http://localhost:3001/customers"
	}

	orders, err := fetchAndUnmarshal[Order](orderUrl)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	customers, err := fetchAndUnmarshal[Customer](customerUrl)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	for _, order := range orders {
		fmt.Printf("\nOrder ID: %d\n", order.ID)
		var customer Customer
		for _, c := range customers {
			if c.Id == order.CustomerId {
				customer = c
			}
		}
		fmt.Printf("Customer: { %s, %s, %s }\n", customer.FirstName, customer.LastName, customer.Email)
		fmt.Printf("Products:\n")
		for _, product := range order.Products {
			fmt.Printf("\t{ %s, %d }\n", product.Name, product.Price)
		}
	}
}
