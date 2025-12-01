package controllers

import (
	"gulmen/app/requests"
	"gulmen/app/responses"
	"gulmen/app/services"
	"gulmen/helpers"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body requests.RegisterRequest true "Register Request"
// @Success 201 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var req requests.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	user, token, err := ctrl.authService.Register(&req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Created(c, "User registered successfully", gin.H{
		"user":  user,
		"token": token,
	})
}

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body requests.LoginRequest true "Login Request"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var req requests.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	user, token, err := ctrl.authService.Login(&req)
	if err != nil {
		responses.Unauthorized(c, err.Error())
		return
	}

	responses.Success(c, "Login successful", gin.H{
		"user":  user,
		"token": token,
	})
}

// GetProfile godoc
// @Summary Get user profile
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Router /auth/profile [get]
func (ctrl *AuthController) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	user, err := ctrl.authService.GetProfile(userID)
	if err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.Success(c, "Profile retrieved successfully", user)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body requests.UpdateProfileRequest true "Update Profile Request"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /auth/profile [put]
func (ctrl *AuthController) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req requests.UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	user, err := ctrl.authService.UpdateProfile(userID, &req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Success(c, "Profile updated successfully", user)
}

// ChangePassword godoc
// @Summary Change user password
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body requests.ChangePasswordRequest true "Change Password Request"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /auth/change-password [post]
func (ctrl *AuthController) ChangePassword(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req requests.ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	if err := ctrl.authService.ChangePassword(userID, &req); err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Success(c, "Password changed successfully", nil)
}

// RefreshToken godoc
// @Summary Refresh JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Router /auth/refresh [post]
func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	userID := c.GetUint("user_id")
	email := c.GetString("email")

	token, err := helpers.GenerateJWT(userID, email)
	if err != nil {
		responses.InternalServerError(c, "Failed to generate token")
		return
	}

	responses.Success(c, "Token refreshed successfully", gin.H{
		"token": token,
	})
}
