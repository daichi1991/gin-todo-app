package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required, min=2, max=100"`
	Description string `json:"description"`
}
