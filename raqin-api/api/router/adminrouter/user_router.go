package adminrouter

import (
	"raqin-api/app/user"
	"raqin-api/storage"

	"github.com/labstack/echo/v4"
)

func userRouter(gRoute *echo.Group) {
	db := storage.DBInstance()
	userRepo := user.NewUserRepo(db)
	userCtrl := user.NewUserController(user.NewUserService(userRepo))

	gRoute.POST("/new", userCtrl.SignUp)
	gRoute.POST("/login", userCtrl.SignIn)
	gRoute.DELETE("/delete", userCtrl.DeleteUser)
	gRoute.PATCH("/update", userCtrl.UpdateUser)
	gRoute.GET("/all", userCtrl.AllUsers)
	gRoute.GET("/promote", userCtrl.PromoteUser)
	gRoute.GET("/demote", userCtrl.DemoteUser)
	gRoute.PATCH("/chpwd", userCtrl.ChangePassword)
}
