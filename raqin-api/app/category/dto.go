package category

import "time"

type NewCategoryRequest struct {
	Name string `json:"name" validate:"required,alpha"`
}

type DeleteCategoryRequest struct {
	ID int `query:"id" validate:"required,number"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id" validate:"required,number"`
	Name string `json:"name" validate:"required,alpha"`
}

type CategoryResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCategoryByIDRequest struct {
	ID int `query:"id" validate:"required,number"`
}
