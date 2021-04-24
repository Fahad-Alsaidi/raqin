package book

import "io"

type NewBookRequest struct {
	Name     string        `form:"name"`
	Notes    string        `form:"notes"`
	Category []string      `form:"category"`
	Authors  []string      `form:"authors"`
	File     io.ReadCloser `form:"file"`
}

type NewBookResponse struct {
	Name     string   `json:"name"`
	Notes    string   `json:"notes"`
	Category []string `json:"categories"`
	Authors  []string `json:"authors"`
}
