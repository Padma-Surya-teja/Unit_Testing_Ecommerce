package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupDb(db_user *string, db_password *string, db_name *string) *gorm.DB {
	database_link := "user=" + *db_user + " password=" + *db_password + " dbname=" + *db_name + " sslmode=disable"
	db, err := gorm.Open("postgres", database_link)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Rating{})
	db.AutoMigrate(&Variant{})

	return db
}

type Product struct {
	ID           int `gorm:"auto_increment"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Product_Name string `sql:"unique_index:id_product"`
	Description  string
	Category     string
	Quantity     int
	Price        int
	Image        string `gorm:"type:varchar(255);"`
	Variant      []Variant
	Rating       []Rating
}

type Rating struct {
	ID        int `gorm:"auto_increment"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductID uint   `sql:"unique_index:id_name"`
	Name      string `sql:"unique_index:id_name"`
	Review    string
	Rating    int `gorm:"check:rating>1&rating<5"`
}

type Variant struct {
	ID        int `gorm:"auto_increment"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductID uint   `json:"productid" sql:"unique_index:product_and_color"`
	Color     string `json:"color" sql:"unique_index:product_and_color"`
	Image     string `json:"image"`
}
