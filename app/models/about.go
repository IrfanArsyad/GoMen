package models

type About struct {
	BaseModel
	Title       string `json:"title" gorm:"size:255;not null"`
	Description string `json:"description" gorm:"type:text"`
	Content     string `json:"content" gorm:"type:text"`
}

func (About) TableName() string {
	return "abouts"
}
