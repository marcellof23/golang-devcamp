package main

import (
	"github.com/marcellof23/devcamp-day4/config"
	"github.com/marcellof23/devcamp-day4/core/entity"
	"github.com/marcellof23/devcamp-day4/core/module"
	"github.com/marcellof23/devcamp-day4/handler"
	productrepository "github.com/marcellof23/devcamp-day4/repository"
	"github.com/marcellof23/devcamp-day4/routes"
)

func main() {

	db := config.Init()
	db.AutoMigrate(&entity.Product{}, &entity.Variant{})

	productRepo := productrepository.New()
	productUc := module.NewProductUsecase(productRepo)
	productHdl := handler.NewProductHandler(productUc)

	r := routes.SetupRoutes(db, *productHdl)
	r.Run(":8080")
}
