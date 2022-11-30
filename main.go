package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Car struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	Model *Model `json:"model"`
}
type Model struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var cars []Car

// getALL CARS function
func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)

}

// Delete CARS function
func deleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode(cars)
}

// Get car function
func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreateCar function (postMethod)
func createCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = strconv.Itoa(rand.Intn(1000000000))
	cars = append(cars, car)
	json.NewEncoder(w).Encode(car)
}

// updateCar Function //Created by ad1tya
func updateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			var car Car
			_ = json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars, car)
			json.NewEncoder(w).Encode(car)
			return

		}
	}
}
func main() {
	r := mux.NewRouter()

	cars = append(cars, Car{ID: "1", Isbn: "446688", Title: "koenigsegg", Model: &Model{Firstname: "Agera", Lastname: "rs"}})
	cars = append(cars, Car{ID: "2", Isbn: "446622", Title: "Lamborghini", Model: &Model{Firstname: "Aventador", Lastname: "svj"}})
	r.HandleFunc("/cars", getCars).Methods("GET")
	r.HandleFunc("/cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/cars", createCar).Methods("POST")
	r.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")
	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
