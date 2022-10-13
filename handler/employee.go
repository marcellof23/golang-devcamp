package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcellof23/ocbc-practice-day4/core/module"
)

type EmployeeHandler struct {
	employeeUc module.EmployeeUsecase
}

func NewEmployeeHandler(employeeUc module.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{
		employeeUc: employeeUc,
	}
}

func (hdl *EmployeeHandler) GetAll(c *gin.Context) {
	Employees, err := hdl.employeeUc.GetEmployees(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Employees})
}

func (hdl *EmployeeHandler) GetSingle(c *gin.Context) {
	Employee, err := hdl.employeeUc.GetEmployee(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

func (hdl *EmployeeHandler) Create(c *gin.Context) {
	err := hdl.employeeUc.CreateEmployee(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "employee succesfully created"})
}

func (hdl *EmployeeHandler) Delete(c *gin.Context) {
	err := hdl.employeeUc.DeleteEmployee(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "employee succesfully deleted"})
}
