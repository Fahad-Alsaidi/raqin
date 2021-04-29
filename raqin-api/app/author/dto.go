package author

import "time"

type NewAuthorRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type DeleteAuthorRequest struct {
	ID int `query:"id"`
}

type UpdateAuthorRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AuthorResponse struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAuthorByIDRequest struct {
	ID int `query:"id"`
}
