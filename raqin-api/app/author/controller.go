package author

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type authorController struct {
	authorservice AuthorService
}

func NewAuthorController(authorservice AuthorService) *authorController {
	return &authorController{authorservice: authorservice}
}

func (aCtrl *authorController) NewAuthor(c echo.Context) error {

	author := NewAuthorRequest{}
	if err := c.Bind(&author); err != nil {
		return err
	}

	resp, err := aCtrl.authorservice.NewAuthor(author)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)

}

func (aCtrl *authorController) DeleteAuthor(c echo.Context) error {

	author := DeleteAuthorRequest{}
	if err := c.Bind(&author); err != nil {
		return err
	}

	err := aCtrl.authorservice.DeleteAuthor(author)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)

}

func (aCtrl *authorController) UpdateAuthor(c echo.Context) error {

	author := UpdateAuthorRequest{}
	if err := c.Bind(&author); err != nil {
		return err
	}

	resp, err := aCtrl.authorservice.UpdateAuthor(author)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &resp)

}

func (aCtrl *authorController) AllAuthors(c echo.Context) error {

	authors, err := aCtrl.authorservice.AllAuthors()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, authors)

}

func (aCtrl *authorController) AuthorByID(c echo.Context) error {

	author := GetAuthorByIDRequest{}
	if err := c.Bind(&author); err != nil {
		return err
	}

	authors, err := aCtrl.authorservice.AuthorByID(author)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &authors)

}
