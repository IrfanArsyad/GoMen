package services

import (
	"errors"
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
		return nil, errors.New("product not found")
	}

	return &product, nil
}

func (s *ProductService) Create(req *requests.CreateProductRequest) (*models.Product, error) {
	db := config.GetDB()

	product := models.Product{
		// Map request fields to model
		// Name: req.Name,
	}

	if err := db.Create(&product).Error; err != nil {
		return nil, errors.New("failed to create product")
	}

	return &product, nil
}

func (s *ProductService) Update(id uint, req *requests.UpdateProductRequest) (*models.Product, error) {
	db := config.GetDB()

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, errors.New("product not found")
	}

	// Update fields
	// product.Name = req.Name

	if err := db.Save(&product).Error; err != nil {
		return nil, errors.New("failed to update product")
	}

	return &product, nil
}

func (s *ProductService) Delete(id uint) error {
	db := config.GetDB()

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return errors.New("product not found")
	}

	if err := db.Delete(&product).Error; err != nil {
		return errors.New("failed to delete product")
	}

	return nil
}
