package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"

	database "ecommerce.com/m/database"
	"ecommerce.com/m/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func GetmockProduct() models.Product {
	// Parsing datetime of type string  to time.time
	layout := "2006-01-02T15:04:05.000000"
	createdat, err := time.Parse(layout, "2022-12-01T12:23:47.419334")
	checkErr(err)
	updatedat, err := time.Parse(layout, "2022-12-01T12:23:47.419334")
	checkErr(err)
	// Creating the mock Product
	prod1 := models.Product{ID: 1, CreatedAt: createdat, UpdatedAt: updatedat, Product_Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Category: "Laptops", Quantity: 100, Price: 88000, Image: "jhbguhfbvidshgfbjhds"}

	return prod1
}

func GetmockRating(id string) models.Rating {
	// Parsing datetime of type string  to time.time
	layout := "2006-01-02T15:04:05.000000"
	createdat, err := time.Parse(layout, "2022-12-01T12:23:47.419334")
	checkErr(err)
	updatedat, err := time.Parse(layout, "2022-12-01T12:23:47.419334")
	checkErr(err)
	Productid, err := strconv.Atoi(id)
	checkErr(err)
	rating := models.Rating{ID: 3, CreatedAt: createdat, UpdatedAt: updatedat, ProductID: uint(Productid), Name: "Surya", Review: "Super fast and Display quality is amazing", Rating: 5}

	return rating
}
func TestGetProducts(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := database.NewMockDataBase(controller)
	s := Server{
		Db: mockDb,
	}

	// Parsing datetime of type string  to time.time
	layout := "2006-01-02T15:04:05.000000"
	createdat, err := time.Parse(layout, "2022-12-01T12:23:47.419334")
	s.CheckErr(err)
	updatedat, err := time.Parse(layout, "2022-12-01T12:23:47.419334")
	s.CheckErr(err)

	// Creating the mock Product
	prod1 := models.Product{ID: 1, CreatedAt: createdat, UpdatedAt: updatedat, Product_Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Category: "Laptops", Quantity: 100, Price: 88000, Image: "jhbguhfbvidshgfbjhds"}
	mockProducts := []models.Product{prod1}

	mockDb.EXPECT().GetAllProducts().Return(mockProducts, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.GetProducts)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	decoder := json.NewDecoder(rr.Body)
	var got JsonResponse
	err = decoder.Decode(&got)
	s.CheckErr(err)
	var expected = JsonResponse{Type: "success", Data: mockProducts}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}

}

// Test for Adding Product
func TestAddProduct(t *testing.T) {
	product := models.Product{
		Product_Name: "Samsung A20",
		Description:  "Super Amoled Display",
		Category:     "Mobile Phones",
		Quantity:     100,
		Price:        11500,
		Image:        "jhbguhfbdfgbdshgfbjhds",
		Variant: []models.Variant{
			{ProductID: 2, Color: "blue", Image: "hjbdjfdgdsfdfhbv"},
			{ProductID: 2, Color: "black", Image: "hdfgerfgedifjb"},
		},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(product)
	req, _ := http.NewRequest("POST", "/api/products/create", &buf) //BTW check for error
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := database.NewMockDataBase(controller)
	s := Server{
		Db: mockDb,
	}
	s.CheckErr(err)
	mockDb.EXPECT().CreateProduct(product).Return(nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.AddProduct)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.

	expected := JsonResponse{Type: "success", Message: "The product has been inserted successfully!"}
	decoder := json.NewDecoder(rr.Body)
	var got JsonResponse
	err = decoder.Decode(&got)
	s.CheckErr(err)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v ",
			got, expected)
	}

}

func TestGetProduct(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/products/1/reviews/create", nil)

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDataBase(controller)
	s := Server{
		Db: mockDb,
	}

	// Creating the mock Product
	prod1 := GetmockProduct()

	params := mux.Vars(req)
	mockDb.EXPECT().GetParticularProduct(params["id"]).Return(prod1, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.GetProduct)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	decoder := json.NewDecoder(rr.Body)
	var got models.Product
	err = decoder.Decode(&got)
	s.CheckErr(err)
	var expected = prod1
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
}

func TestAddReview(t *testing.T) {
	product := GetmockRating("2")
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(product)
	req, _ := http.NewRequest("POST", "api/products/2/reviews/create", &buf) //BTW check for error
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := database.NewMockDataBase(controller)
	s := Server{
		Db: mockDb,
	}
	s.CheckErr(err)

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	checkErr(err)
	mockDb.EXPECT().CreateReview(uint(id), product).Return(nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.AddReview)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.

	expected := JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}
	decoder := json.NewDecoder(rr.Body)
	var got JsonResponse
	err = decoder.Decode(&got)
	s.CheckErr(err)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v ",
			got, expected)
	}
}
