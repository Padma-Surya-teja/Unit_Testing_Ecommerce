package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"ecommerce.com/m/database"
	"ecommerce.com/m/handlers"
	"ecommerce.com/m/models"
	"github.com/gorilla/mux"
)

func main() {
	//connect tyo database
	db_user := flag.String("user", "postgres", "database user")
	db_password := flag.String("password", "root", "database password")
	db_name := flag.String("name", "Ecommerce_microservice", "database name")
	db := models.SetupDb(db_user, db_password, db_name)

	var s handlers.Server = handlers.Server{Db: database.DBClient{Db: db}}
	//Init the mux router
	router := mux.NewRouter()

	//Get all Products
	router.HandleFunc("/api/products", s.GetProducts).Methods("GET")

	// //Add a Product
	router.HandleFunc("/api/products/create", s.AddProduct).Methods("POST")

	//Get a particular Product
	router.HandleFunc("/api/products/{id}", s.GetProduct).Methods("GET")

	// //Add review to a particular product
	router.HandleFunc("/api/products/{id}/reviews/create", s.AddReview).Methods("POST")

	// //Get a particular product review
	// router.HandleFunc("/api/products/{id}/reviews", database.GetProductReviews).Methods("GET")

	// //Update details of a review
	// router.HandleFunc("/api/products/{id}/reviews/{rid}", database.UpdateReview).Methods("PATCH")

	// //Delete a particular review
	// router.HandleFunc("/api/products/{id}/reviews/{rid}", database.DeleteReview).Methods("DELETE")

	fmt.Println("Server at 8080")

	//Serving
	log.Fatal(http.ListenAndServe(":8080", router))
}
