package controllers

import (
	"gomen/app/requests"
	"gomen/app/responses"
	"gomen/app/services"
	"gomen/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		productService: services.NewProductService(),
	}
}

// Index godoc
// @Summary Get all Products
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Success 200 {object} responses.PaginatedResponse
// @Failure 401 {object} responses.Response
// @Router /products [get]
func (ctrl *ProductController) Index(c *gin.Context) {
	params := helpers.GetPaginationParams(c)

	product, pagination, err := ctrl.productService.GetAll(params)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Paginated(c, "Products retrieved successfully", product, pagination)
}

// Show godoc
// @Summary Get Product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /products/{id} [get]
func (ctrl *ProductController) Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid product ID", nil)
		return
	}

	product, err := ctrl.productService.GetByID(uint(id))
	if err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.Success(c, "Product retrieved successfully", product)
}

// Store godoc
// @Summary Create a new Product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body requests.CreateProductRequest true "Create Product Request"
// @Success 201 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /products [post]
func (ctrl *ProductController) Store(c *gin.Context) {
	var req requests.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	product, err := ctrl.productService.Create(&req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Created(c, "Product created successfully", product)
}

// Update godoc
// @Summary Update a Product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param request body requests.UpdateProductRequest true "Update Product Request"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Router /products/{id} [put]
func (ctrl *ProductController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid product ID", nil)
		return
	}

	var req requests.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body", nil)
		return
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		responses.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	product, err := ctrl.productService.Update(uint(id), &req)
	if err != nil {
		responses.BadRequest(c, err.Error(), nil)
		return
	}

	responses.Success(c, "Product updated successfully", product)
}

// Delete godoc
// @Summary Delete a Product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 204 {object} nil
// @Failure 404 {object} responses.Response
// @Router /products/{id} [delete]
func (ctrl *ProductController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.BadRequest(c, "Invalid product ID", nil)
		return
	}

	if err := ctrl.productService.Delete(uint(id)); err != nil {
		responses.NotFound(c, err.Error())
		return
	}

	responses.NoContent(c)
}
