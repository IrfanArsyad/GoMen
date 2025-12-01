package helpers

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginationParams struct {
	Page    int
	PerPage int
	Offset  int
}

func GetPaginationParams(c *gin.Context) PaginationParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100
	}

	offset := (page - 1) * perPage

	return PaginationParams{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
	}
}

func Paginate(params PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(params.Offset).Limit(params.PerPage)
	}
}

func CalculateTotalPages(total int64, perPage int) int {
	return int(math.Ceil(float64(total) / float64(perPage)))
}
