package database

import (
	"ecommerce.com/m/models"
	"github.com/jinzhu/gorm"
)

type DataBase interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(models.Product) error
	GetParticularProduct(string) (models.Product, error)
	CreateReview(uint, models.Rating) error
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

func (db DBClient) GetParticularProduct(id string) (models.Product, error) {
	var product models.Product
	db.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)
	return product, nil
}

func (db DBClient) CreateReview(id uint, t models.Rating) error {
	t.ProductID = id
	var product models.Product
	db.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

	product.Rating = append(product.Rating, t)
	db.Db.Save(&product)

	return nil
}
