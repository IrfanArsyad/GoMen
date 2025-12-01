package models

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"size:255;not null"`
	Email    string `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Password string `json:"-" gorm:"size:255;not null"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}

func (User) TableName() string {
	return "users"
}
