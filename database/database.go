package database

import (
	"ecommerce.com/m/models"
	"github.com/jinzhu/gorm"
)

type DataBase interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(models.Product) error
}

type DBClient struct {
	Db *gorm.DB
}

func (db DBClient) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	db.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Debug().Find(&products)
	return products, nil
}

func (db DBClient) CreateProduct(t models.Product) error {
	db.Db.Create(&t)
	return nil
}
