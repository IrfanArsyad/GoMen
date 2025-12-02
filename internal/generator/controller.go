package generator

import (
	"fmt"
	"path/filepath"
)

// MakeController generates a new controller file
func MakeController(name string) {
	pascalName := toPascalCase(name)
	snakeName := toSnakeCase(pascalName)
	camelName := toCamelCase(pascalName)
	pluralName := toPlural(pascalName)
	pluralSnake := toSnakeCase(pluralName)

	content := fmt.Sprintf(`package controllers

import (
	"gomen/app/requests"
	"gomen/app/responses"
	"gomen/app/services"
	"gomen/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type %sController struct {
	%sService *services.%sService
}

func New%sController() *%sController {
	return &%sController{
		%sService: services.New%sService(),
	}
}

// Index godoc
// @Summary Get all %s
// @Tags %s
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Success 200 {object} responses.PaginatedResponse
// @Failure 401 {object} responses.Response
// @Router /%s [get]
func (ctrl *%sController) Index(c *gin.Context) {
	params := helpers.GetPaginationParams(c)

	%s, pagination, err := ctrl.%sService.GetAll(params)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Paginated(c, "%s retrieved successfully", %s, pagination)
}

// Show godoc
// @Summary Get %s by ID
// @Tags %s
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "%s ID"
// @Success 200 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /%s/{id} [get]
func (ctrl *%sController) Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid %s ID", nil)
		return
	}

	%s, err := ctrl.%sService.GetByID(uint(id))
	if err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.Success(c, "%s retrieved successfully", %s)
}

// Store godoc
// @Summary Create a new %s
// @Tags %s
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body requests.Create%sRequest true "Create %s Request"
// @Success 201 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /%s [post]
func (ctrl *%sController) Store(c *gin.Context) {
	var req requests.Create%sRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	%s, err := ctrl.%sService.Create(&req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Created(c, "%s created successfully", %s)
}

// Update godoc
// @Summary Update a %s
// @Tags %s
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "%s ID"
// @Param request body requests.Update%sRequest true "Update %s Request"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /%s/{id} [put]
func (ctrl *%sController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid %s ID", nil)
		return
	}

	var req requests.Update%sRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	%s, err := ctrl.%sService.Update(uint(id), &req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Success(c, "%s updated successfully", %s)
}

// Delete godoc
// @Summary Delete a %s
// @Tags %s
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "%s ID"
// @Success 204 {object} nil
// @Failure 404 {object} responses.Response
// @Router /%s/{id} [delete]
func (ctrl *%sController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid %s ID", nil)
		return
	}

	if err := ctrl.%sService.Delete(uint(id)); err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.NoContent(c)
}
`,
		// Struct and constructor
		pascalName, camelName, pascalName,
		pascalName, pascalName, pascalName, camelName, pascalName,
		// Index
		pluralName, pluralName, pluralSnake,
		pascalName,
		camelName, camelName,
		pluralName, camelName,
		// Show
		pascalName, pluralName, pascalName, pluralSnake,
		pascalName,
		snakeName,
		camelName, camelName,
		pascalName, camelName,
		// Store
		pascalName, pluralName, pascalName, pascalName, pluralSnake,
		pascalName, pascalName,
		camelName, camelName,
		pascalName, camelName,
		// Update
		pascalName, pluralName, pascalName, pascalName, pascalName, pluralSnake,
		pascalName,
		snakeName,
		pascalName,
		camelName, camelName,
		pascalName, camelName,
		// Delete
		pascalName, pluralName, pascalName, pluralSnake,
		pascalName,
		snakeName,
		camelName,
	)

	filePath := filepath.Join(getProjectRoot(), "app", "controllers", snakeName+"_controller.go")

	if err := writeFile(filePath, content); err != nil {
		printError(err)
		return
	}

	printSuccess("Controller", filePath)
}
