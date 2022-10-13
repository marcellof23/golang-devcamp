package module

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/marcellof23/ocbc-practice-day4/core/entity"
	"github.com/marcellof23/ocbc-practice-day4/core/repository"
)

type AttendanceUsecase interface {
	GetAttendances(c *gin.Context) ([]entity.Attendance, error)
	CreateAttendance(c *gin.Context) error
}

type attendanceUsecase struct {
	attendanceRepo repository.AttendanceRepository
}

// NewAttendanceUsecase use for initiate new attendance usecase
func NewAttendanceUsecase(repo repository.AttendanceRepository) AttendanceUsecase {
	return &attendanceUsecase{repo}
}

var ErrAttendanceNotFound = errors.New("attendance not found")

func (em *attendanceUsecase) GetAttendances(c *gin.Context) ([]entity.Attendance, error) {
	data, err := em.attendanceRepo.FindAllAttendance(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordAttendanceNotFound) {
			return nil, fmt.Errorf("%w.", ErrAttendanceNotFound)
		}
		return nil, fmt.Errorf("%w: %v", ErrAttendanceNotFound, err)
	}
	return data, nil
}

func (em *attendanceUsecase) CreateAttendance(c *gin.Context) error {
	err := em.attendanceRepo.CreateAttendance(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordAttendanceNotFound) {
			return fmt.Errorf("%w.", ErrAttendanceNotFound)
		}
		return fmt.Errorf("%w: %v", ErrAttendanceNotFound, err)
	}
	return nil
}
