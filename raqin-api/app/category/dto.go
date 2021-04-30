package category

import "time"

type NewCategoryRequest struct {
	Name string `json:"name"`
}

type DeleteCategoryRequest struct {
	ID int `query:"id"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCategoryByIDRequest struct {
	ID int `query:"id"`
}
