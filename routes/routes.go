package routes

import (
	"be-go-product-sales/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB, r *gin.Engine) {
	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/product", controller.GetAllProduct)
	r.POST("/product", controller.AddProduct)
	r.PATCH("/product/:id", controller.UpdateProduct)
	r.DELETE("/product/:id", controller.DeleteProduct)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
