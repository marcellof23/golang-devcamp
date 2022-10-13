package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/marcellof23/ocbc-practice-day4/handler"
)

func SetupRoutes(db *gorm.DB, employeeHdl handler.EmployeeHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/employees", employeeHdl.GetAll)
	r.GET("/employees/:id", employeeHdl.GetSingle)
	r.POST("/employees", employeeHdl.Create)
	r.DELETE("/employees/:id", employeeHdl.Delete)
	return r
}
