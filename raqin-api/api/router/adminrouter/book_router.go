package adminrouter

import (
	bookController "raqin-api/api/controller/bookcontroller"
	"raqin-api/core/service/bookservice"

	"github.com/labstack/echo/v4"
)

func bookRouter(gRoute *echo.Group) {
	// init NewBookController with NewBookService
	bookCtrl := bookController.NewBookController(bookservice.NewBookService())

	gRoute.POST("/upload", bookCtrl.NewBook)
}
