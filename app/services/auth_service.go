package services

import (
	"errors"
	"gulmen/app/models"
	"gulmen/app/requests"
	"gulmen/config"
	"gulmen/helpers"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(req *requests.RegisterRequest) (*models.User, string, error) {
	db := config.GetDB()

	// Check if email already exists
	var existingUser models.User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, "", errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, "", errors.New("failed to hash password")
	}

	// Create user
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: true,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, "", errors.New("failed to create user")
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return &user, token, nil
}

func (s *AuthService) Login(req *requests.LoginRequest) (*models.User, string, error) {
	db := config.GetDB()

	var user models.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, "", errors.New("account is not active")
	}

	if !helpers.CheckPassword(req.Password, user.Password) {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := helpers.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return &user, token, nil
}

func (s *AuthService) GetProfile(userID uint) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s *AuthService) UpdateProfile(userID uint, req *requests.UpdateProfileRequest) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	user.Name = req.Name

	if err := db.Save(&user).Error; err != nil {
		return nil, errors.New("failed to update profile")
	}

	return &user, nil
}

func (s *AuthService) ChangePassword(userID uint, req *requests.ChangePasswordRequest) error {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	if !helpers.CheckPassword(req.CurrentPassword, user.Password) {
		return errors.New("current password is incorrect")
	}

	hashedPassword, err := helpers.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = hashedPassword

	if err := db.Save(&user).Error; err != nil {
		return errors.New("failed to update password")
	}

	return nil
}
