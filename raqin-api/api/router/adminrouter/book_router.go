package adminrouter

import (
	"raqin-api/app/book"

	"github.com/labstack/echo/v4"
)

func bookRouter(gRoute *echo.Group) {
	// init NewBookController with NewBookService
	bookCtrl := book.NewBookController(book.NewBookService())

	gRoute.POST("/upload", bookCtrl.UploadNewBook)
}
