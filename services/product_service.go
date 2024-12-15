package services

import (
	"errors"
	"product-management/database"
	"product-management/models"
)

func CreateProductService(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}

func GetProductService(id string) (*models.Product, error) {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil
}
package services

import (
	"errors"
	"product-management/database"
	"product-management/models"
)

func CreateProductService(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}

func GetProductService(id string) (*models.Product, error) {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil
}
package services

import (
	"errors"
	"product-management/database"
	"product-management/models"
)

func CreateProductService(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}

func GetProductService(id string) (*models.Product, error) {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil
}
