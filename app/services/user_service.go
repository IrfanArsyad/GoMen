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
		helpers.Error(err, "Failed to get user by ID").
			Uint("user_id", id).
			Msg("Database query failed")
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s *UserService) Create(req *requests.CreateUserRequest) (*models.User, error) {
	db := config.GetDB()

	// Check if email already exists
	var existingUser models.User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		helpers.Warn("Attempted to create user with existing email").
			Str("email", req.Email).
			Msg("Email already registered")
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		helpers.Error(err, "Failed to hash password").
			Str("email", req.Email).
			Msg("Password hashing error")
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: true,
	}

	if err := db.Create(&user).Error; err != nil {
		helpers.Error(err, "Failed to create user").
			Str("email", req.Email).
			Str("name", req.Name).
			Msg("Database insert failed")
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	helpers.Info("User created successfully").
		Uint("user_id", user.ID).
		Str("email", user.Email).
		Msg("New user registered")

	return &user, nil
}

func (s *UserService) Update(id uint, req *requests.UpdateUserRequest) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		helpers.Error(err, "User not found for update").
			Uint("user_id", id).
			Msg("Database query failed")
		return nil, errors.New("user not found")
	}

	// Check if email is being changed and already exists
	if req.Email != user.Email {
		var existingUser models.User
		if err := db.Where("email = ? AND id != ?", req.Email, id).First(&existingUser).Error; err == nil {
			helpers.Warn("Attempted to update user with existing email").
				Uint("user_id", id).
				Str("new_email", req.Email).
				Msg("Email already registered")
			return nil, errors.New("email already registered")
		}
	}

	user.Name = req.Name
	user.Email = req.Email

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := db.Save(&user).Error; err != nil {
		helpers.Error(err, "Failed to update user").
			Uint("user_id", id).
			Str("email", req.Email).
			Msg("Database update failed")
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	helpers.Info("User updated successfully").
		Uint("user_id", user.ID).
		Str("email", user.Email).
		Msg("User information modified")

	return &user, nil
}

func (s *UserService) Delete(id uint) error {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		helpers.Error(err, "User not found for deletion").
			Uint("user_id", id).
			Msg("Database query failed")
		return errors.New("user not found")
	}

	userEmail := user.Email

	if err := db.Delete(&user).Error; err != nil {
		helpers.Error(err, "Failed to delete user").
			Uint("user_id", id).
			Str("email", userEmail).
			Msg("Database delete failed")
		return fmt.Errorf("failed to delete user: %w", err)
	}

	helpers.Info("User deleted successfully").
		Uint("user_id", id).
		Str("email", userEmail).
		Msg("User removed from database")

	return nil
}
