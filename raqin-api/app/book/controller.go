package book

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookController struct {
	bookservice BookService
}

func NewBookController(bookservice BookService) *bookController {
	return &bookController{bookservice: bookservice}
}

func (bCtrl *bookController) NewBook(c echo.Context) error {

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

	resp, err := bCtrl.bookservice.NewBook(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)

}

func (bCtrl *bookController) UpdateBook(c echo.Context) error {

	book := UpdateBookRequest{}
	if err := c.Bind(&book); err != nil {
		panic(err)
	}

	resp, err := bCtrl.bookservice.UpdateBook(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (bCtrl *bookController) BookByID(c echo.Context) error {

	book := BookIDRequest{}
	if err := c.Bind(&book); err != nil {
		panic(err)
	}

	resp, err := bCtrl.bookservice.BookByID(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (bCtrl *bookController) ExtractBook(c echo.Context) error {

	book := BookIDRequest{}
	if err := c.Bind(&book); err != nil {
		panic(err)
	}

	resp, err := bCtrl.bookservice.BookToPages(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
