package controllers

import (
	"gomen/app/requests"
	"gomen/app/responses"
	"gomen/app/services"
	"gomen/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// Index godoc
// @Summary Get all users
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Success 200 {object} responses.PaginatedResponse
// @Failure 401 {object} responses.Response
// @Router /users [get]
func (ctrl *UserController) Index(c *gin.Context) {
	params := helpers.GetPaginationParams(c)

	users, pagination, err := ctrl.userService.GetAll(params)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Paginated(c, "Users retrieved successfully", users, pagination)
}

// Show godoc
// @Summary Get user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /users/{id} [get]
func (ctrl *UserController) Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid user ID", nil)
		return
	}

	user, err := ctrl.userService.GetByID(uint(id))
	if err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.Success(c, "User retrieved successfully", user)
}

// Store godoc
// @Summary Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body requests.CreateUserRequest true "Create User Request"
// @Success 201 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /users [post]
func (ctrl *UserController) Store(c *gin.Context) {
	var req requests.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	user, err := ctrl.userService.Create(&req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Created(c, "User created successfully", user)
}

// Update godoc
// @Summary Update a user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param request body requests.UpdateUserRequest true "Update User Request"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /users/{id} [put]
func (ctrl *UserController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid user ID", nil)
		return
	}

	var req requests.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	user, err := ctrl.userService.Update(uint(id), &req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Success(c, "User updated successfully", user)
}

// Delete godoc
// @Summary Delete a user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 204 {object} nil
// @Failure 404 {object} responses.Response
// @Router /users/{id} [delete]
func (ctrl *UserController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid user ID", nil)
		return
	}

	if err := ctrl.userService.Delete(uint(id)); err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.NoContent(c)
}
