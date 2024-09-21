package routes

import (
	"github.com/cepot-blip/catalog-app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    productRoutes := r.Group("/products")
    {
        productRoutes.GET("/", controllers.GetAllProducts)
        productRoutes.GET("/:id", controllers.GetProductByID)
        productRoutes.POST("/", controllers.CreateProduct)
        productRoutes.PUT("/:id", controllers.UpdateProduct)
        productRoutes.DELETE("/:id", controllers.DeleteProduct)
    }

    return r
}
