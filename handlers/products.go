package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "ecommerce.com/m/database"
	"ecommerce.com/m/models"
	"github.com/gorilla/mux"
)

type JsonResponse struct {
	Type    string           `json:"type"`
	Data    []models.Product `json:"data"`
	Message string           `json:"message"`
}

type Server struct {
	Db database.DataBase
}

type PRODUCT interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
}

// Function for checking errors
func (s Server) CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for handling messages
func (s Server) PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// returns all the products
func (s Server) GetProducts(w http.ResponseWriter, r *http.Request) {

	s.PrintMessage("Getting Products...")

	w.WriteHeader(http.StatusOK)
	products, _ := s.Db.GetAllProducts()
	s.PrintMessage("Received the Products...")
	var response = JsonResponse{Type: "success", Data: products}

	json.NewEncoder(w).Encode(response)
}

// Add products
func (s Server) AddProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.Product
	err := decoder.Decode(&t)

	s.CheckErr(err)

	w.WriteHeader(http.StatusOK)
	var response = JsonResponse{}

	s.PrintMessage("Inserting Product into Db")

	s.Db.CreateProduct(t)
	response = JsonResponse{Type: "success", Message: "The product has been inserted successfully!"}

	json.NewEncoder(w).Encode(response)
}

// Getting a Particular Product by ID
func (s Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	w.WriteHeader(http.StatusOK)
	s.PrintMessage("Getting Products...")

	product, err := s.Db.GetParticularProduct(id)
	s.CheckErr(err)
	json.NewEncoder(w).Encode(product)
}
