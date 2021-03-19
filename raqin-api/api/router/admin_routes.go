package router

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// InitAdminRouter will register all admin routes
// هنا نعرف وجهات المستخدم المتحكم كالمثال الموجود
func InitAdminRouter(clientRouter *echo.Group) {

	// profile route وجهة الصفحة الشخصية للمتحكم
	clientRouter.GET("/profile", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World from Admin!")
	})

	clientRouter.POST("/upload", func(c echo.Context) error {
		// Read form fields
		name := c.FormValue("name")
		author := c.FormValue("author")

		//-----------
		// Read file
		//-----------

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, author))
	})
}
