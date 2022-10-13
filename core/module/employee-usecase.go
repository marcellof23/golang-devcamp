package module

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/marcellof23/ocbc-practice-day4/core/entity"
	"github.com/marcellof23/ocbc-practice-day4/core/repository"
)

type EmployeeUsecase interface {
	GetEmployees(c *gin.Context) ([]entity.Employee, error)
	GetEmployee(c *gin.Context) (entity.Employee, error)
	CreateEmployee(c *gin.Context) error
	DeleteEmployee(c *gin.Context) error
}

type employeeUsecase struct {
	employeeRepo repository.EmployeeRepository
}

// NewEmployeeUsecase use for initiate new employee usecase
func NewEmployeeUsecase(repo repository.EmployeeRepository) EmployeeUsecase {
	return &employeeUsecase{repo}
}

var ErrEmployeeNotFound = errors.New("employee not found")

func (em *employeeUsecase) GetEmployees(c *gin.Context) ([]entity.Employee, error) {
	data, err := em.employeeRepo.FindAll(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordEmployeeNotFound) {
			return nil, fmt.Errorf("%w.", ErrEmployeeNotFound)
		}
		return nil, fmt.Errorf("%w: %v", ErrEmployeeNotFound, err)
	}
	return data, nil
}

func (em *employeeUsecase) GetEmployee(c *gin.Context) (entity.Employee, error) {
	data, err := em.employeeRepo.FindSingle(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordEmployeeNotFound) {
			return entity.Employee{}, fmt.Errorf("%w.", ErrEmployeeNotFound)
		}
		return entity.Employee{}, fmt.Errorf("%w: %v", ErrEmployeeNotFound, err)
	}
	return data, nil
}

func (em *employeeUsecase) CreateEmployee(c *gin.Context) error {
	err := em.employeeRepo.Create(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordEmployeeNotFound) {
			return fmt.Errorf("%w.", ErrEmployeeNotFound)
		}
		return fmt.Errorf("%w: %v", ErrEmployeeNotFound, err)
	}
	return nil
}

func (em *employeeUsecase) DeleteEmployee(c *gin.Context) error {
	err := em.employeeRepo.Delete(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordEmployeeNotFound) {
			return fmt.Errorf("%w.", ErrEmployeeNotFound)
		}
		return fmt.Errorf("%w: %v", ErrEmployeeNotFound, err)
	}
	return nil
}
