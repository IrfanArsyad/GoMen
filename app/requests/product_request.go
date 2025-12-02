package requests

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description" validate:"max=1000"`
	Price       float64 `json:"price" validate:"required,gte=0"`
	Stock       int     `json:"stock" validate:"gte=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description" validate:"max=1000"`
	Price       float64 `json:"price" validate:"required,gte=0"`
	Stock       int     `json:"stock" validate:"gte=0"`
}
