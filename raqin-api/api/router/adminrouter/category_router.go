package adminrouter

import (
	"raqin-api/app/category"
	"raqin-api/storage"

	"github.com/labstack/echo/v4"
)

func categoryRouter(gRoute *echo.Group) {
	// init NewBookController with NewBookService
	db := storage.DBInstance()
	categoryRepo := category.NewCategoryRepo(db)
	categoryCtrl := category.NewCategoryController(category.NewCategoryService(categoryRepo))

	gRoute.POST("/new", categoryCtrl.NewCategory)
	gRoute.DELETE("/delete", categoryCtrl.DeleteCategory)
	gRoute.PATCH("/update", categoryCtrl.UpdateCategory)
	gRoute.GET("/all", categoryCtrl.AllCategories)
	gRoute.GET("", categoryCtrl.CategoryByID)
}
