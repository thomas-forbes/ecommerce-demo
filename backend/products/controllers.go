package products

import (
	"backend/database"
)

func (pm *productManager) GetProducts() (*[]database.Product, error) {
	var products []database.Product
	pm.DB.Find(&products)
	return &products, nil
}

func (pm *productManager) CreateProduct(product *database.Product) (*database.Product, error) {
	pm.DB.Create(product)
	return product, nil

}
