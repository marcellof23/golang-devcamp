package main

import (
	"github.com/marcellof23/ocbc-practice-day4/config"
	"github.com/marcellof23/ocbc-practice-day4/core/entity"
	"github.com/marcellof23/ocbc-practice-day4/core/module"
	"github.com/marcellof23/ocbc-practice-day4/handler"
	employeerepository "github.com/marcellof23/ocbc-practice-day4/repository"
	"github.com/marcellof23/ocbc-practice-day4/routes"
)

func main() {

	db := config.Init()
	db.Model(&entity.Employee{}).Related(&entity.Attendance{})
	db.AutoMigrate(&entity.Employee{}, entity.Attendance{})
	db.Model(&entity.Attendance{}).AddForeignKey("employee_id", "employees(id)", "RESTRICT", "RESTRICT")

	employeeRepo := employeerepository.New()
	employeeUc := module.NewEmployeeUsecase(employeeRepo)
	employeeHdl := handler.NewEmployeeHandler(employeeUc)

	defer db.Close()
	r := routes.SetupRoutes(db, *employeeHdl)
	r.Run(":8081")
}
