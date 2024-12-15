package controllers

import (
	"net/http"
	"product-management/models"
	"product-management/services"
	"product-management/utils"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := services.CreateProductService(&product); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.RespondJSON(c, http.StatusCreated, product)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := services.GetProductService(id)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Product not found")
		return
	}

	utils.RespondJSON(c, http.StatusOK, product)
}
