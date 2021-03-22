package adminrouter

import (
	"github.com/labstack/echo/v4"
)

// InitAdminRouter will register all admin routes
// هنا نعرف وجهات المستخدم المتحكم كالمثال الموجود
func Init(gRouter *echo.Group) {

	// all book routes will be inside bookRouter function
	bookRouter(gRouter.Group("/book"))
}
