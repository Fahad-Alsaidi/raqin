package book

import "io"

type NewBookRequest struct {
	Name     string        `form:"name"`
	Notes    string        `form:"notes"`
	Category []string      `form:"category"`
	Authors  []string      `form:"authors"`
	File     io.ReadCloser `form:"file"`
}
