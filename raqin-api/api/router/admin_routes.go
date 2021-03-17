package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// InitAdminRouter will register all admin routes
// هنا نعرف وجهات المستخدم المتحكم كالمثال الموجود
func InitAdminRouter(clientRouter *echo.Group) {

	// profile route وجهة الصفحة الشخصية للمتحكم
	clientRouter.GET("/profile", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World from Admin!")
	})
}
