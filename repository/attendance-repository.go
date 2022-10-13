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

type repositoryAttendance struct {
}

func NewAttendance() repository_intf.AttendanceRepository {
	return &repositoryAttendance{}
}

func (r *repositoryAttendance) FindAllAttendance(c *gin.Context) ([]entity.Attendance, error) {
	attendances := []entity.Attendance{}

	db := c.MustGet("db").(*gorm.DB)
	err := db.Find(&attendances).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordAttendanceNotFound
		}
		return nil, err
	}

	return attendances, nil
}

func (r *repositoryAttendance) CreateAttendance(c *gin.Context) error {
	var input entity.CreateAttendanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return errors.New("failed to create attendance")
	}
	dates := "2006-01-02"
	date, _ := time.Parse(dates, input.Date)

	// Create Employee
	Attendance := entity.Attendance{EmployeeID: input.EmployeeID, Date: date}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&Attendance).Error; err != nil {
		return errors.New("failed to create attendance")
	}

	return nil
}
