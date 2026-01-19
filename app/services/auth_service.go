package services

import (
	"errors"
	"fmt"
	"gomen/app/models"
	"gomen/app/requests"
	"gomen/config"
	"gomen/helpers"
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
		helpers.Warn("Registration attempted with existing email").
			Str("email", req.Email).
			Msg("Email already registered")
		return nil, "", errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		helpers.Error(err, "Failed to hash password during registration").
			Str("email", req.Email).
			Msg("Password hashing error")
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
		helpers.Error(err, "Failed to create user during registration").
			Str("email", req.Email).
			Str("name", req.Name).
			Msg("Database insert failed")
		return nil, "", fmt.Errorf("failed to create user: %w", err)
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(user.ID, user.Email)
	if err != nil {
		helpers.Error(err, "Failed to generate JWT token").
			Uint("user_id", user.ID).
			Str("email", user.Email).
			Msg("Token generation error")
		return nil, "", errors.New("failed to generate token")
	}

	helpers.Info("User registered successfully").
		Uint("user_id", user.ID).
		Str("email", user.Email).
		Msg("New user registration completed")

	return &user, token, nil
}

func (s *AuthService) Login(req *requests.LoginRequest) (*models.User, string, error) {
	db := config.GetDB()

	var user models.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		helpers.Warn("Login attempt with non-existent email").
			Str("email", req.Email).
			Msg("Invalid credentials")
		return nil, "", errors.New("invalid credentials")
	}

	if !user.IsActive {
		helpers.Warn("Login attempt on inactive account").
			Uint("user_id", user.ID).
			Str("email", user.Email).
			Msg("Account is not active")
		return nil, "", errors.New("account is not active")
	}

	if !helpers.CheckPassword(req.Password, user.Password) {
		helpers.Warn("Login attempt with incorrect password").
			Uint("user_id", user.ID).
			Str("email", user.Email).
			Msg("Invalid credentials")
		return nil, "", errors.New("invalid credentials")
	}

	token, err := helpers.GenerateJWT(user.ID, user.Email)
	if err != nil {
		helpers.Error(err, "Failed to generate JWT token during login").
			Uint("user_id", user.ID).
			Str("email", user.Email).
			Msg("Token generation error")
		return nil, "", errors.New("failed to generate token")
	}

	helpers.Info("User logged in successfully").
		Uint("user_id", user.ID).
		Str("email", user.Email).
		Msg("Login successful")

	return &user, token, nil
}

func (s *AuthService) GetProfile(userID uint) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		helpers.Error(err, "Failed to get user profile").
			Uint("user_id", userID).
			Msg("Database query failed")
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s *AuthService) UpdateProfile(userID uint, req *requests.UpdateProfileRequest) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		helpers.Error(err, "User not found for profile update").
			Uint("user_id", userID).
			Msg("Database query failed")
		return nil, errors.New("user not found")
	}

	user.Name = req.Name

	if err := db.Save(&user).Error; err != nil {
		helpers.Error(err, "Failed to update user profile").
			Uint("user_id", userID).
			Str("new_name", req.Name).
			Msg("Database update failed")
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	helpers.Info("User profile updated successfully").
		Uint("user_id", user.ID).
		Str("email", user.Email).
		Msg("Profile information modified")

	return &user, nil
}

func (s *AuthService) ChangePassword(userID uint, req *requests.ChangePasswordRequest) error {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		helpers.Error(err, "User not found for password change").
			Uint("user_id", userID).
			Msg("Database query failed")
		return errors.New("user not found")
	}

	if !helpers.CheckPassword(req.CurrentPassword, user.Password) {
		helpers.Warn("Password change attempt with incorrect current password").
			Uint("user_id", userID).
			Str("email", user.Email).
			Msg("Current password is incorrect")
		return errors.New("current password is incorrect")
	}

	hashedPassword, err := helpers.HashPassword(req.NewPassword)
	if err != nil {
		helpers.Error(err, "Failed to hash new password").
			Uint("user_id", userID).
			Msg("Password hashing error")
		return errors.New("failed to hash password")
	}

	user.Password = hashedPassword

	if err := db.Save(&user).Error; err != nil {
		helpers.Error(err, "Failed to update password").
			Uint("user_id", userID).
			Str("email", user.Email).
			Msg("Database update failed")
		return fmt.Errorf("failed to update password: %w", err)
	}

	helpers.Info("User password changed successfully").
		Uint("user_id", user.ID).
		Str("email", user.Email).
		Msg("Password updated")

	return nil
}
