package services

import (
	"errors"
	"gomen/app/models"
	"gomen/app/requests"
	"gomen/app/responses"
	"gomen/config"
	"gomen/helpers"
)

type AboutService struct{}

func NewAboutService() *AboutService {
	return &AboutService{}
}

func (s *AboutService) GetAll(params helpers.PaginationParams) ([]models.About, responses.Pagination, error) {
	db := config.GetDB()

	var about []models.About
	var total int64

	db.Model(&models.About{}).Count(&total)
	db.Scopes(helpers.Paginate(params)).Find(&about)

	pagination := responses.Pagination{
		CurrentPage: params.Page,
		PerPage:     params.PerPage,
		Total:       total,
		TotalPages:  helpers.CalculateTotalPages(total, params.PerPage),
	}

	return about, pagination, nil
}

func (s *AboutService) GetByID(id uint) (*models.About, error) {
	db := config.GetDB()

	var about models.About
	if err := db.First(&about, id).Error; err != nil {
		return nil, errors.New("about not found")
	}

	return &about, nil
}

func (s *AboutService) Create(req *requests.CreateAboutRequest) (*models.About, error) {
	db := config.GetDB()

	about := models.About{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
	}

	if err := db.Create(&about).Error; err != nil {
		return nil, errors.New("failed to create about")
	}

	return &about, nil
}

func (s *AboutService) Update(id uint, req *requests.UpdateAboutRequest) (*models.About, error) {
	db := config.GetDB()

	var about models.About
	if err := db.First(&about, id).Error; err != nil {
		return nil, errors.New("about not found")
	}

	about.Title = req.Title
	about.Description = req.Description
	about.Content = req.Content

	if err := db.Save(&about).Error; err != nil {
		return nil, errors.New("failed to update about")
	}

	return &about, nil
}

func (s *AboutService) Delete(id uint) error {
	db := config.GetDB()

	var about models.About
	if err := db.First(&about, id).Error; err != nil {
		return errors.New("about not found")
	}

	if err := db.Delete(&about).Error; err != nil {
		return errors.New("failed to delete about")
	}

	return nil
}
