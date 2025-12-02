package models

type Product struct {
	BaseModel
	Name        string  `json:"name" gorm:"size:255;not null"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"not null;default:0"`
	Stock       int     `json:"stock" gorm:"not null;default:0"`
}

func (Product) TableName() string {
	return "products"
}
