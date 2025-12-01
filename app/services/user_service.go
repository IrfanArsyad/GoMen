package services

import (
	"errors"
	"gulmen/app/models"
	"gulmen/app/requests"
	"gulmen/app/responses"
	"gulmen/config"
	"gulmen/helpers"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAll(params helpers.PaginationParams) ([]models.User, responses.Pagination, error) {
	db := config.GetDB()

	var users []models.User
	var total int64

	db.Model(&models.User{}).Count(&total)
	db.Scopes(helpers.Paginate(params)).Find(&users)

	pagination := responses.Pagination{
		CurrentPage: params.Page,
		PerPage:     params.PerPage,
		Total:       total,
		TotalPages:  helpers.CalculateTotalPages(total, params.PerPage),
	}

	return users, pagination, nil
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s *UserService) Create(req *requests.CreateUserRequest) (*models.User, error) {
	db := config.GetDB()

	// Check if email already exists
	var existingUser models.User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: true,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}

func (s *UserService) Update(id uint, req *requests.UpdateUserRequest) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Check if email is being changed and already exists
	if req.Email != user.Email {
		var existingUser models.User
		if err := db.Where("email = ? AND id != ?", req.Email, id).First(&existingUser).Error; err == nil {
			return nil, errors.New("email already registered")
		}
	}

	user.Name = req.Name
	user.Email = req.Email

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := db.Save(&user).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	return &user, nil
}

func (s *UserService) Delete(id uint) error {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	if err := db.Delete(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
