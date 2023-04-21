package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       int
	Name     string
	Quantity int
	Pirce    float64
}

var Products = []Product{}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage")
	log.Println("Homepage endpoint hit")
}

func handleAllRequests() {
	http.HandleFunc("/home", homepage)
	http.HandleFunc("/products", returAllProducts)
	http.ListenAndServe("localhost:10000", nil)
}

func returAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("All products endpoint hit")
	json.NewEncoder(w).Encode(Products)
}

func main() {
	
	Products = []Product{
		{Id: 1, Name: "Pen", Quantity: 10, Pirce: 5.99},
		{Id: 2, Name: "Pencil", Quantity: 5, Pirce: 2.99},
	}

	handleAllRequests()
}
