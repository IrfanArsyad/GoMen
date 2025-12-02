package controllers

import (
	"gomen/app/requests"
	"gomen/app/responses"
	"gomen/app/services"
	"gomen/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AboutController struct {
	aboutService *services.AboutService
}

func NewAboutController() *AboutController {
	return &AboutController{
		aboutService: services.NewAboutService(),
	}
}

// Index godoc
// @Summary Get all Abouts
// @Tags Abouts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Success 200 {object} responses.PaginatedResponse
// @Failure 401 {object} responses.Response
// @Router /abouts [get]
func (ctrl *AboutController) Index(c *gin.Context) {
	params := helpers.GetPaginationParams(c)

	about, pagination, err := ctrl.aboutService.GetAll(params)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Paginated(c, "Abouts retrieved successfully", about, pagination)
}

// Show godoc
// @Summary Get About by ID
// @Tags Abouts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "About ID"
// @Success 200 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /abouts/{id} [get]
func (ctrl *AboutController) Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid about ID", nil)
		return
	}

	about, err := ctrl.aboutService.GetByID(uint(id))
	if err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.Success(c, "About retrieved successfully", about)
}

// Store godoc
// @Summary Create a new About
// @Tags Abouts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body requests.CreateAboutRequest true "Create About Request"
// @Success 201 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /abouts [post]
func (ctrl *AboutController) Store(c *gin.Context) {
	var req requests.CreateAboutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	about, err := ctrl.aboutService.Create(&req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Created(c, "About created successfully", about)
}

// Update godoc
// @Summary Update a About
// @Tags Abouts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "About ID"
// @Param request body requests.UpdateAboutRequest true "Update About Request"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /abouts/{id} [put]
func (ctrl *AboutController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid about ID", nil)
		return
	}

	var req requests.UpdateAboutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	about, err := ctrl.aboutService.Update(uint(id), &req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Success(c, "About updated successfully", about)
}

// Delete godoc
// @Summary Delete a About
// @Tags Abouts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "About ID"
// @Success 204 {object} nil
// @Failure 404 {object} responses.Response
// @Router /abouts/{id} [delete]
func (ctrl *AboutController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid about ID", nil)
		return
	}

	if err := ctrl.aboutService.Delete(uint(id)); err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.NoContent(c)
}
