package dto

type CreateStatusInput struct {
	Name string `json:"name" binding:"required"`
}
