package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	Name       string    `json:"name"`
	BirthDate  time.Time `json:"birthDate"`
	Address    string    `json:"address"`
	Job        string    `json:"job"`
	JoinDate   time.Time `json:"joinDate"`
	Attendance []Attendance
}

type CreateEmployeeInput struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Job       string `json:"job"`
	JoinDate  string `json:"joinDate"`
}

type UpdateEmployeeInput struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Job       string `json:"job"`
	JoinDate  string `json:"joinDate"`
}
