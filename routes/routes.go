package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/marcellof23/devcamp-day4/handler"
)

func SetupRoutes(db *gorm.DB, productHdl handler.ProductHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/products", productHdl.GetAll)
	r.GET("/products/:id", productHdl.GetSingle)
	r.PUT("/products/:id", productHdl.Update)
	r.POST("/products", productHdl.Create)
	r.DELETE("/products/:id", productHdl.Delete)
	return r
}
