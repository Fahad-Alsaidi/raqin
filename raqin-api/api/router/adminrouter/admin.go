package adminrouter

import (
	"github.com/labstack/echo/v4"
)

// Init will register all admin routes
func Init(gRouter *echo.Group) {

	// all book routes will be inside bookRouter
	bookRouter(gRouter.Group("/book"))

	// all author routes will be inside authorRouter
	authorRouter(gRouter.Group("/author"))

	// all category routes will be inside categoryRouter
	categoryRouter(gRouter.Group("/category"))
}
