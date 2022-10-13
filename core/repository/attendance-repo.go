package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/marcellof23/ocbc-practice-day4/core/entity"
)

var ErrRecordAttendanceNotFound = errors.New("record not found")

type AttendanceRepository interface {
	FindAllAttendance(c *gin.Context) ([]entity.Attendance, error)
	CreateAttendance(c *gin.Context) error
}
