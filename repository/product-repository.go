package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/marcellof23/devcamp-day4/core/entity"
	repository_intf "github.com/marcellof23/devcamp-day4/core/repository"
)

type repository struct {
}

func New() repository_intf.ProductRepository {
	return &repository{}
}

func (r *repository) FindAll(c *gin.Context) ([]entity.Product, error) {
	var products []entity.Product

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	err := db.Model(&entity.Product{}).Preload("Variants").Find(&products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordProductNotFound
		}
		return nil, err
	}

	return products, nil
}

func (r *repository) FindSingle(c *gin.Context) (entity.Product, error) {
	product := entity.Product{}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.Product{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		return entity.Product{}, repository_intf.ErrRecordProductNotFound
	}

	return product, nil
}

func (r *repository) Create(c *gin.Context) error {
	var input entity.ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return errors.New("failed to create product")
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	var arrayVariants []entity.Variant
	for _, v := range input.Variants {
		arrayVariants = append(arrayVariants, entity.Variant{
			VariantName: v.VariantName,
			Price:       v.Price,
			Stock:       v.Stock,
		})
	}

	// Create Product
	Product := entity.Product{
		ProductName: input.ProductName,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Discount:    input.Discount,
		Variants:    arrayVariants,
	}

	if err := db.Create(&Product).Error; err != nil {
		return errors.New("failed to create product")
	}

	return nil
}

func (r *repository) Update(c *gin.Context) error {
	var input entity.ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return errors.New("failed to update product")
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	var arrayVariants []entity.Variant
	for _, v := range input.Variants {
		arrayVariants = append(arrayVariants, entity.Variant{
			VariantName: v.VariantName,
			Price:       v.Price,
			Stock:       v.Stock,
		})
	}

	// Create Product
	Product := entity.Product{
		ProductName: input.ProductName,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Discount:    input.Discount,
		Variants:    arrayVariants,
	}

	if err := db.Where("id = ?", c.Param("id")).Save(&Product).Error; err != nil {
		return errors.New("failed to update product")
	}

	return nil
}

func (r *repository) Delete(c *gin.Context) error {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	var product entity.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		return errors.New("failed to delete product")
	}

	if err := db.Delete(&product).Error; err != nil {
		return errors.New("failed to delete product")
	}

	return nil
}
