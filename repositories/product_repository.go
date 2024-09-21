package repositories

import (
	"github.com/cepot-blip/catalog-app/config"
	"github.com/cepot-blip/catalog-app/models"
)

func GetAllProducts() ([]models.Product, error) {
    var products []models.Product
    err := config.DB.Find(&products).Error
    return products, err
}

func GetProductByID(id uint) (models.Product, error) {
    var product models.Product
    err := config.DB.First(&product, id).Error
    return product, err
}

func CreateProduct(product *models.Product) error {
    return config.DB.Create(product).Error
}

func UpdateProduct(product *models.Product) error {
    return config.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
    return config.DB.Delete(&models.Product{}, id).Error
}
