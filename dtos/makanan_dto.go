package dtos

type MakananCreateRequest struct {
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"required"`
}

type MakananResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MakananUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
