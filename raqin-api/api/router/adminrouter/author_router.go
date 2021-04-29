package adminrouter

import (
	"raqin-api/app/author"
	"raqin-api/storage"

	"github.com/labstack/echo/v4"
)

func authorRouter(gRoute *echo.Group) {
	db := storage.DBInstance()
	authorCtrl := author.NewAuthorController(author.NewBookService(author.NewAuthorRepo(db)))

	gRoute.POST("/new", authorCtrl.NewAuthor)
	gRoute.DELETE("/delete", authorCtrl.DeleteAuthor)
	gRoute.PATCH("/update", authorCtrl.UpdateAuthor)
	gRoute.GET("/all", authorCtrl.AllAuthors)
	gRoute.GET("", authorCtrl.AuthorByID)
}
