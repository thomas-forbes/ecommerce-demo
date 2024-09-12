package products

import (
	"backend/database"

	"gorm.io/gorm"
)

type productManager struct {
	DB *gorm.DB
}

func NewProductManager(db *gorm.DB) ProductProvider {
	return &productManager{DB: db}
}

type ProductProvider interface {
	GetProducts() (*[]database.Product, error)
	CreateProduct(product *database.Product) (*database.Product, error)
}
