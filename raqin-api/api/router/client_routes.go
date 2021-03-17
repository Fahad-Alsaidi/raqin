package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// InitClientRouter will register all client routes
// هنا نعرف وجهات المستخدم العادي كالمثال الموجود
func InitClientRouter(clientRouter *echo.Group) {

	// profile route وجهة الصفحة الشخصية للمستخدم العادي
	clientRouter.POST("/profile", func(c echo.Context) error {
		answer := c.FormValue("answer")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, Did you send '%v'", answer))
	})
}
