package services

import (
	"errors"
	"fmt"
	"gomen/app/models"
	"gomen/app/requests"
	"gomen/app/responses"
	"gomen/config"
	"gomen/helpers"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetAll(params helpers.PaginationParams) ([]models.Product, responses.Pagination, error) {
	db := config.GetDB()

	var product []models.Product
	var total int64

	db.Model(&models.Product{}).Count(&total)
	db.Scopes(helpers.Paginate(params)).Find(&product)

	pagination := responses.Pagination{
		CurrentPage: params.Page,
		PerPage:     params.PerPage,
		Total:       total,
		TotalPages:  helpers.CalculateTotalPages(total, params.PerPage),
	}

	return product, pagination, nil
}

func (s *ProductService) GetByID(id uint) (*models.Product, error) {
	db := config.GetDB()

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		helpers.Error(err, "Failed to get product by ID").
			Uint("product_id", id).
			Msg("Database query failed")
		return nil, errors.New("product not found")
	}

	return &product, nil
}

func (s *ProductService) Create(req *requests.CreateProductRequest) (*models.Product, error) {
	db := config.GetDB()

	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := db.Create(&product).Error; err != nil {
		helpers.Error(err, "Failed to create product").
			Str("product_name", req.Name).
			Float64("price", req.Price).
			Msg("Database insert failed")
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	helpers.Info("Product created successfully").
		Uint("product_id", product.ID).
		Str("product_name", product.Name).
		Msg("New product added")

	return &product, nil
}

func (s *ProductService) Update(id uint, req *requests.UpdateProductRequest) (*models.Product, error) {
	db := config.GetDB()

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		helpers.Error(err, "Product not found for update").
			Uint("product_id", id).
			Msg("Database query failed")
		return nil, errors.New("product not found")
	}

	// Update fields
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock

	if err := db.Save(&product).Error; err != nil {
		helpers.Error(err, "Failed to update product").
			Uint("product_id", id).
			Str("product_name", req.Name).
			Msg("Database update failed")
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	helpers.Info("Product updated successfully").
		Uint("product_id", product.ID).
		Str("product_name", product.Name).
		Msg("Product information modified")

	return &product, nil
}

func (s *ProductService) Delete(id uint) error {
	db := config.GetDB()

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		helpers.Error(err, "Product not found for deletion").
			Uint("product_id", id).
			Msg("Database query failed")
		return errors.New("product not found")
	}

	productName := product.Name

	if err := db.Delete(&product).Error; err != nil {
		helpers.Error(err, "Failed to delete product").
			Uint("product_id", id).
			Str("product_name", productName).
			Msg("Database delete failed")
		return fmt.Errorf("failed to delete product: %w", err)
	}

	helpers.Info("Product deleted successfully").
		Uint("product_id", id).
		Str("product_name", productName).
		Msg("Product removed from database")

	return nil
}
