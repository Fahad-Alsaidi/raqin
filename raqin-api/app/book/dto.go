package book

import "io"

type NewBookRequest struct {
	Name     string        `form:"name"`
	Notes    string        `form:"notes"`
	Category []int         `form:"category"`
	Authors  []int         `form:"authors"`
	File     io.ReadCloser `form:"file"`
}

type NewBookResponse struct {
	Name     string   `json:"name"`
	Notes    string   `json:"notes"`
	Category []string `json:"categories"`
	Authors  []string `json:"authors"`
}
