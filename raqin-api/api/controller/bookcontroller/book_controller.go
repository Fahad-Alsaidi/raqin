package bookController

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"raqin-api/core/model"
	"raqin-api/core/service/bookservice"

	"github.com/labstack/echo/v4"
)

// book struct that will implement IBook interface
type bookController struct {
	bookservice bookservice.IBook
}

// NewBookController return bookController struct with IBook service
// this method is the entry to this controller
func NewBookController(bookservice bookservice.IBook) bookController {
	return bookController{bookservice: bookservice}
}

// NewBook returns the new created book
func (b bookController) NewBook(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	author := c.FormValue("author")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, author))
}

// Books returns all books
func (b bookController) Books(c echo.Context) error {

	return nil
}

// BookByID returns a book by its id
func (b bookController) BookByID(c echo.Context) error {
	book := model.Book{}
	if err := c.Bind(book); err != nil {
		return err
	}
	return nil
}
