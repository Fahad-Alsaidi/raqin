package adminrouter

import (
	"raqin-api/app/book"
	"raqin-api/storage"

	"github.com/labstack/echo/v4"
)

func bookRouter(gRoute *echo.Group) {
	db := storage.DBInstance()
	bookRepo := book.NewBookRepo(db)
	bookCtrl := book.NewBookController(book.NewBookService(bookRepo))

	gRoute.POST("/new", bookCtrl.NewBook)
	gRoute.DELETE("/delete", bookCtrl.DeleteBook)
	gRoute.POST("/update", bookCtrl.UpdateBook)
	gRoute.GET("/all", bookCtrl.AllBooks)
	gRoute.GET("/extract", bookCtrl.ExtractBook)
	gRoute.GET("/add/author", bookCtrl.AddBookAuthor)
	gRoute.GET("/add/category", bookCtrl.AddBookCategory)
	gRoute.GET("/add/initiater", bookCtrl.AddBookInitiator)
	gRoute.GET("/remove/author", bookCtrl.RemoveBookAuthor)
	gRoute.GET("/remove/category", bookCtrl.RemoveBookCategory)
	gRoute.GET("/remove/initiater", bookCtrl.RemoveBookInitiator)
}
