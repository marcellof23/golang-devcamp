package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Attendance struct {
	gorm.Model
	EmployeeID uint
	Date       time.Time `json:"date"`
}

type CreateAttendanceInput struct {
	EmployeeID uint
	Date       string `json:"date"`
}
