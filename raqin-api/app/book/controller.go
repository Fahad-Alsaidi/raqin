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

func (bCtrl *bookController) DeleteBook(c echo.Context) error {

	book := BookIDRequest{}
	if err := c.Bind(&book); err != nil {
		return err
	}

	err := bCtrl.bookservice.DeleteBook(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (bCtrl *bookController) AllBooks(c echo.Context) error {

	books, err := bCtrl.bookservice.AllBooks()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, books)
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

func (bCtrl *bookController) AddBookAuthor(c echo.Context) error {

	bkAuthor := AddBookRel{}
	if err := c.Bind(&bkAuthor); err != nil {
		panic(err)
	}

	err := bCtrl.bookservice.AddBookAuthor(bkAuthor)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (bCtrl *bookController) RemoveBookAuthor(c echo.Context) error {

	bkAuthor := RemoveBookRel{}
	if err := c.Bind(&bkAuthor); err != nil {
		panic(err)
	}

	err := bCtrl.bookservice.RemoveBookAuthor(bkAuthor)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (bCtrl *bookController) AddBookCategory(c echo.Context) error {

	bkCategory := AddBookRel{}
	if err := c.Bind(&bkCategory); err != nil {
		panic(err)
	}

	err := bCtrl.bookservice.AddBookCategory(bkCategory)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (bCtrl *bookController) RemoveBookCategory(c echo.Context) error {

	bkCategory := RemoveBookRel{}
	if err := c.Bind(&bkCategory); err != nil {
		panic(err)
	}

	err := bCtrl.bookservice.RemoveBookCategory(bkCategory)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (bCtrl *bookController) AddBookInitiator(c echo.Context) error {

	bkInitiater := AddBookRel{}
	if err := c.Bind(&bkInitiater); err != nil {
		panic(err)
	}

	err := bCtrl.bookservice.AddBookInitiator(bkInitiater)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (bCtrl *bookController) RemoveBookInitiator(c echo.Context) error {

	bkInitiater := RemoveBookRel{}
	if err := c.Bind(&bkInitiater); err != nil {
		panic(err)
	}

	err := bCtrl.bookservice.RemoveBookInitiator(bkInitiater)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
