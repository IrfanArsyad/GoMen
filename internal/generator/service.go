package generator

import (
	"fmt"
	"path/filepath"
)

// MakeService generates a new service file
func MakeService(name string) {
	pascalName := toPascalCase(name)
	snakeName := toSnakeCase(pascalName)
	camelName := toCamelCase(pascalName)

	content := fmt.Sprintf(`package services

import (
	"errors"
	"gomen/app/models"
	"gomen/app/requests"
	"gomen/app/responses"
	"gomen/config"
	"gomen/helpers"
)

type %sService struct{}

func New%sService() *%sService {
	return &%sService{}
}

func (s *%sService) GetAll(params helpers.PaginationParams) ([]models.%s, responses.Pagination, error) {
	db := config.GetDB()

	var %s []models.%s
	var total int64

	db.Model(&models.%s{}).Count(&total)
	db.Scopes(helpers.Paginate(params)).Find(&%s)

	pagination := responses.Pagination{
		CurrentPage: params.Page,
		PerPage:     params.PerPage,
		Total:       total,
		TotalPages:  helpers.CalculateTotalPages(total, params.PerPage),
	}

	return %s, pagination, nil
}

func (s *%sService) GetByID(id uint) (*models.%s, error) {
	db := config.GetDB()

	var %s models.%s
	if err := db.First(&%s, id).Error; err != nil {
		return nil, errors.New("%s not found")
	}

	return &%s, nil
}

func (s *%sService) Create(req *requests.Create%sRequest) (*models.%s, error) {
	db := config.GetDB()

	%s := models.%s{
		// Map request fields to model
		// Name: req.Name,
	}

	if err := db.Create(&%s).Error; err != nil {
		return nil, errors.New("failed to create %s")
	}

	return &%s, nil
}

func (s *%sService) Update(id uint, req *requests.Update%sRequest) (*models.%s, error) {
	db := config.GetDB()

	var %s models.%s
	if err := db.First(&%s, id).Error; err != nil {
		return nil, errors.New("%s not found")
	}

	// Update fields
	// %s.Name = req.Name

	if err := db.Save(&%s).Error; err != nil {
		return nil, errors.New("failed to update %s")
	}

	return &%s, nil
}

func (s *%sService) Delete(id uint) error {
	db := config.GetDB()

	var %s models.%s
	if err := db.First(&%s, id).Error; err != nil {
		return errors.New("%s not found")
	}

	if err := db.Delete(&%s).Error; err != nil {
		return errors.New("failed to delete %s")
	}

	return nil
}
`,
		// Struct and constructor
		pascalName, pascalName, pascalName, pascalName,
		// GetAll
		pascalName, pascalName,
		camelName, pascalName,
		pascalName, camelName,
		camelName,
		// GetByID
		pascalName, pascalName,
		camelName, pascalName, camelName,
		snakeName,
		camelName,
		// Create
		pascalName, pascalName, pascalName,
		camelName, pascalName,
		camelName, snakeName,
		camelName,
		// Update
		pascalName, pascalName, pascalName,
		camelName, pascalName, camelName, snakeName,
		camelName,
		camelName, snakeName,
		camelName,
		// Delete
		pascalName,
		camelName, pascalName, camelName, snakeName,
		camelName, snakeName,
	)

	filePath := filepath.Join(getProjectRoot(), "app", "services", snakeName+"_service.go")

	if err := writeFile(filePath, content); err != nil {
		printError(err)
		return
	}

	printSuccess("Service", filePath)
	fmt.Printf("  → Don't forget to create the model '%s' and request '%sRequest'\n", pascalName, pascalName)
}
