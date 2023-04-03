package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func main() {

	// HandleRequest()
	Products = []Product{
		Product{ID: 1, Name: "Chair", Quantity: 10, Price: 100.00},
		Product{ID: 2, Name: "Desk", Quantity: 200, Price: 220.0},
	}

	HandleRequest()
}

func HandleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", ReturnAllProducts)
	myRouter.HandleFunc("/homepage", homepage)
	myRouter.HandleFunc("/products/{id}", getProduct)
	http.ListenAndServe("localhost:10000", myRouter)

}

func getProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	for _, product := range Products{
		if product.ID == key{
			json.NewEncoder(w).Encode(product)
		}
	}
}

func ReturnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hits: returnall")
	json.NewEncoder(w).Encode(Products)



}

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hits: homepage")
	fmt.Fprintf(w, "Welcome to Homepage !!")
}
