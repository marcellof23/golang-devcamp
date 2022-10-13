package repository

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/marcellof23/ocbc-practice-day4/core/entity"
	repository_intf "github.com/marcellof23/ocbc-practice-day4/core/repository"
)

type repository struct {
}

func New() repository_intf.EmployeeRepository {
	return &repository{}
}

func (r *repository) FindAll(c *gin.Context) ([]entity.Employee, error) {
	employees := []entity.Employee{}

	db := c.MustGet("db").(*gorm.DB)
	err := db.Find(&employees).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordEmployeeNotFound
		}
		return nil, err
	}

	return employees, nil
}

func (r *repository) FindSingle(c *gin.Context) (entity.Employee, error) {
	employee := entity.Employee{}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		return entity.Employee{}, repository_intf.ErrRecordEmployeeNotFound
	}

	return employee, nil
}

func (r *repository) Create(c *gin.Context) error {
	var input entity.CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return errors.New("failed to create employee")
	}
	date := "2006-01-02"
	joinDate, _ := time.Parse(date, input.JoinDate)
	birthDate, _ := time.Parse(date, input.BirthDate)

	// Create Employee
	Employee := entity.Employee{Name: input.Name, BirthDate: birthDate, Address: input.Address, Job: input.Job, JoinDate: joinDate}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&Employee).Error; err != nil {
		return errors.New("failed to create employee")
	}

	return nil
}

func (r *repository) Delete(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)
	var employee entity.Employee
	if err := db.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found!"})
		return errors.New("failed to delete employee")
	}

	if err := db.Delete(&employee).Error; err != nil {
		return errors.New("failed to delete employee")
	}

	return nil
}
