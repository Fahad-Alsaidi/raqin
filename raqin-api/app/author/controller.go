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
