package requests

type CreateAboutRequest struct {
	Title       string `json:"title" validate:"required,min=2,max=255"`
	Description string `json:"description" validate:"max=1000"`
	Content     string `json:"content"`
}

type UpdateAboutRequest struct {
	Title       string `json:"title" validate:"required,min=2,max=255"`
	Description string `json:"description" validate:"max=1000"`
	Content     string `json:"content"`
}
