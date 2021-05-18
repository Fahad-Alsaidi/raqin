package book

import (
	"io"
	"raqin-api/app/author"
	"raqin-api/app/category"
	"raqin-api/app/user"
)

type UpdateBookRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Notes string `json:"notes"`
}

type NewBookRequest struct {
	Name       string        `form:"name"`
	Notes      string        `form:"notes"`
	Category   []int         `form:"category"`
	Authors    []int         `form:"authors"`
	Initiators []int         `form:"initiators"`
	File       io.ReadCloser `form:"file"`
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
	ID int `query:"id"`
}

type AddBookRel struct {
	ID     int `query:"id"`
	BookID int `query:"book_id"`
}

type RemoveBookRel struct {
	ID int `query:"id"`
}
