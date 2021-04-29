package adminrouter

import (
	"raqin-api/app/book"
	"raqin-api/storage"

	"github.com/labstack/echo/v4"
)

func bookRouter(gRoute *echo.Group) {
	db := storage.DBInstance()
	bookCtrl := book.NewBookController(book.NewBookService(book.NewBookRepo(db)))

	gRoute.POST("/upload", bookCtrl.NewBook)
}
