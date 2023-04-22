package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "gorilla/mux"
)

type Product struct {
	Id       string
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
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/home", homepage)
	myRouter.HandleFunc("/products", returAllProducts)
	myRouter.HandleFunc("/product/{id}", getProduct)
	http.ListenAndServe("localhost:10000", myRouter)
}
func getProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	for _, product := range Products {
		if product.Id == id{
			json.NewEncoder(w).Encode(product)
		}
   }
}

func returAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("All products endpoint hit")
	json.NewEncoder(w).Encode(Products)
}

func main() {
	
	Products = []Product{
		{Id: "1", Name: "Pen", Quantity: 10, Pirce: 5.99},
		{Id: "2", Name: "Pencil", Quantity: 5, Pirce: 2.99},
	}

	handleAllRequests()
}
