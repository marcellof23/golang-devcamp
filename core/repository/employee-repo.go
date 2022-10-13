package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/marcellof23/ocbc-practice-day4/core/entity"
)

var ErrRecordEmployeeNotFound = errors.New("record not found")

type EmployeeRepository interface {
	FindAll(c *gin.Context) ([]entity.Employee, error)
	FindSingle(c *gin.Context) (entity.Employee, error)
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
}
