package controllers

import (
	"net/http"
	"strconv"

	"github.com/cepot-blip/catalog-app/models"
	"github.com/cepot-blip/catalog-app/services"
	"github.com/cepot-blip/catalog-app/utils"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
    products, err := services.GetAllProducts()
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get products")
        return
    }
    c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    product, err := services.GetProductByID(uint(id))
    if err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Product not found")
        return
    }
    c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
        return
    }

    if err := services.CreateProduct(&product); err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create product")
        return
    }

    c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
    var product models.Product
    id, _ := strconv.Atoi(c.Param("id"))

    if err := c.ShouldBindJSON(&product); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
        return
    }

    product.ID = uint(id)
    if err := services.UpdateProduct(&product); err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update product")
        return
    }

    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    if err := services.DeleteProduct(uint(id)); err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete product")
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
