package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce.com/m/models"
	"github.com/gorilla/mux"
)

// Adding Review to a Particular Product
func (s Server) AddReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	s.CheckErr(err)
	decoder := json.NewDecoder(r.Body)
	var t models.Rating
	err = decoder.Decode(&t)

	s.CheckErr(err)

	s.PrintMessage("Inserting Review into Db")

	s.Db.CreateReview(uint(id), t)

	response := JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}
	s.PrintMessage("Inserted Review into Db")

	json.NewEncoder(w).Encode(response)
}

// func (s DataBase) GetProductReviews(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	id := params["id"]

// 	s.PrintMessage("Getting Product Reviews...")

// 	var product models.Product
// 	s.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

// 	json.NewEncoder(w).Encode(product.Rating)
// }

// func (s DataBase) UpdateReview(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	id, err := strconv.Atoi(params["id"])
// 	s.CheckErr(err)

// 	rid, err := strconv.Atoi(params["rid"])
// 	s.CheckErr(err)

// 	decoder := json.NewDecoder(r.Body)
// 	var t models.Rating
// 	err = decoder.Decode(&t)

// 	s.CheckErr(err)

// 	var response = JsonResponse{}

// 	if t.ID == 0 {
// 		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
// 	} else {

// 		s.PrintMessage("Inserting Review into Db")

// 		var rating models.Rating

// 		s.Db.Model(&models.Rating{}).Where("product_id=? and id=?", id, rid).Find(&rating)

// 		s.Db.First(&rating)
// 		rating.Rating = t.Rating
// 		rating.Review = t.Review
// 		s.Db.Save(&rating)
// 		// Db.Save(&t);

// 		response = JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}

// 	}

// 	json.NewEncoder(w).Encode(response)
// }

// func (s DataBase) DeleteReview(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	pid, err := strconv.Atoi(params["id"])
// 	s.CheckErr(err)
// 	rid, err := strconv.Atoi(params["rid"])

// 	s.CheckErr(err)
// 	var response = JsonResponse{}
// 	if pid == 0 || rid == 0 {
// 		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
// 	} else {

// 		s.PrintMessage("Deleting a Review in Db")

// 		s.Db.Model(models.Rating{}).Where("product_id=? and id=?", pid, rid).Delete(&models.Rating{})

// 		response = JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}

// 	}

// 	json.NewEncoder(w).Encode(response)

// }
