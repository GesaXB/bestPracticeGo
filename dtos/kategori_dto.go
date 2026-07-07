package dtos

type KategoriCreateRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}

type KategoriResponse struct {
	Name string `json:"name"`
}
