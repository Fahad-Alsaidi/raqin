package book

import (
	"io"
	"raqin-api/app/author"
	"raqin-api/app/category"
	"raqin-api/app/user"
)

type UpdateBookRequest struct {
	ID    int    `json:"id" validate:"required,number"`
	Name  string `json:"name" validate:"required,ascii"`
	Notes string `json:"notes" validate:"required,ascii"`
}

type NewBookRequest struct {
	Name       string        `form:"name" validate:"required,ascii"`
	Notes      string        `form:"notes" validate:"required,ascii"`
	Category   []int         `form:"category" validate:"required,dive,number"`
	Authors    []int         `form:"authors" validate:"required,dive,number"`
	Initiators []int         `form:"initiators" validate:"required,dive,number"`
	File       io.ReadCloser `form:"file" validate:"required"`
}
type BookResponse struct {
	ID       int                         `json:"id"`
	Name     string                      `json:"name"`
	Notes    string                      `json:"notes"`
	Category []category.CategoryResponse `json:"categories"`
	Authors  []author.AuthorResponse     `json:"authors"`
	Users    []user.UserResponse         `json:"initiators"`
}

type BookIDRequest struct {
	ID int `query:"id" validate:"required,number"`
}

type AddBookRel struct {
	ID     int `query:"id" validate:"required,number"`
	BookID int `query:"book_id" validate:"required,number"`
}

type RemoveBookRel struct {
	ID     int `query:"id" validate:"required,number"`
	BookID int `query:"book_id" validate:"required,number"`
}
