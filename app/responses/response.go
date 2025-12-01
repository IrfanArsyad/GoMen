package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// Success response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created response
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// NoContent response
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// Error response
func Error(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// BadRequest response
func BadRequest(c *gin.Context, message string, errors interface{}) {
	Error(c, http.StatusBadRequest, message, errors)
}

// Unauthorized response
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message, nil)
}

// Forbidden response
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message, nil)
}

// NotFound response
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, nil)
}

// UnprocessableEntity response
func UnprocessableEntity(c *gin.Context, message string, errors interface{}) {
	Error(c, http.StatusUnprocessableEntity, message, errors)
}

// InternalServerError response
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message, nil)
}

// Paginated response
func Paginated(c *gin.Context, message string, data interface{}, pagination Pagination) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}
