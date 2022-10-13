package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marcellof23/devcamp-day4/core/module"
)

type ProductHandler struct {
	productUc module.ProductUsecase
}

func NewProductHandler(productUc module.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUc: productUc,
	}
}

func (hdl *ProductHandler) GetAll(c *gin.Context) {
	Products, err := hdl.productUc.GetProducts(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Products})
}

func (hdl *ProductHandler) GetSingle(c *gin.Context) {
	Product, err := hdl.productUc.GetProduct(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Product})
}

func (hdl *ProductHandler) Create(c *gin.Context) {
	err := hdl.productUc.CreateProduct(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "product succesfully created"})
}

func (hdl *ProductHandler) Update(c *gin.Context) {
	err := hdl.productUc.UpdateProduct(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "product succesfully updated"})
}
func (hdl *ProductHandler) Delete(c *gin.Context) {
	err := hdl.productUc.DeleteProduct(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "product succesfully deleted"})
}
