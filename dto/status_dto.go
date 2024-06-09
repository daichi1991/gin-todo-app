package dto

type CreateStatusInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateStatusInput struct {
	Name string `json:"name" binding:"required"`
}
