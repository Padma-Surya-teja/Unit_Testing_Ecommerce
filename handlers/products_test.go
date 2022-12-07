package handlers

import (
	"net/http"
	"testing"

	database "ecommerce.com/m/database"
	"github.com/golang/mock/gomock"
)

func TestGetProducts(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := database.NewMockDataBase(controller)
	s := Server{
		db: mockDb,
	}
	mockDb.EXPECT().GetAllProducts().Return(1)

	var w http.ResponseWriter
	var r *http.Request
	s.GetProducts(w, r)
}
