package book

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// book struct that will implement IBook interface
type Controller struct {
	bookservice IBook
}

// NewBookController return bookController struct with IBook service
// this method is the entry to this controller
func NewBookController(bookservice IBook) Controller {
	return Controller{bookservice: bookservice}
}

// NewBook returns the new created book
func (b Controller) UploadNewBook(c echo.Context) error {

	book := NewBookRequest{}
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	src, err := file.Open()
	if err != nil {
		panic(err)
	}

	book.File = src

	if err := c.Bind(&book); err != nil {
		panic(err)
	}

	b.bookservice.UploadNewBook(book)

	return c.String(http.StatusOK, "book has been uploaded")

}
