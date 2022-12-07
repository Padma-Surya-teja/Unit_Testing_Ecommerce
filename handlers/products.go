package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "ecommerce.com/m/database"
	"ecommerce.com/m/models"
)

type JsonResponse struct {
	Type    string           `json:"type"`
	Data    []models.Product `json:"data"`
	Message string           `json:"message"`
}

type Server struct {
	db database.DataBase
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

	products, _ := s.db.GetAllProducts()
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

	var response = JsonResponse{}

	s.PrintMessage("Inserting Product into Db")

	s.db.CreateProduct(t)
	response = JsonResponse{Type: "success", Message: "The product has been inserted successfully!"}

	json.NewEncoder(w).Encode(response)
}

// func (s Server) GetProduct(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	id := params["id"]

// 	s.PrintMessage("Getting Products...")

// 	var product []models.Product
// 	s.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

// 	json.NewEncoder(w).Encode(product)
// }
