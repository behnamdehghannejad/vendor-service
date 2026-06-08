package dto

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
	Path     string `json:"path"`
}

type RequestUpdateCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
