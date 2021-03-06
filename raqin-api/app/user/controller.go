package user

import (
	"fmt"
	"net/http"
	"raqin-api/api/middleware"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userService UserService
}

func NewUserController(userService UserService) *userController {
	return &userController{userService: userService}
}

func (uCtrl *userController) SignUp(c echo.Context) error {

	user := UserSignUp{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	_, err := uCtrl.userService.SignUp(user)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	token, err := middleware.GenerateJWTToken(name, user.Email, user.Gender, "VOLUNTEER")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}

func (uCtrl *userController) SignIn(c echo.Context) error {

	user := UserSignIn{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp, err := uCtrl.userService.SignIn(user)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("%s %s", resp.FirstName, resp.LastName)
	token, err := middleware.GenerateJWTToken(name, resp.Email, resp.Gender, resp.Role)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}

func (uCtrl *userController) DeleteUser(c echo.Context) error {

	user := UserIDRequest{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err := uCtrl.userService.DeleteUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (uCtrl *userController) UpdateUser(c echo.Context) error {

	user := UpdateUserRequest{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err := uCtrl.userService.UpdateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (uCtrl *userController) AllUsers(c echo.Context) error {

	users, err := uCtrl.userService.AllUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (uCtrl *userController) ChangePassword(c echo.Context) error {

	user := ChangePasswordRequest{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err := uCtrl.userService.ChangePassword(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (uCtrl *userController) PromoteUser(c echo.Context) error {

	user := UserIDRequest{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err := uCtrl.userService.PromoteUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (uCtrl *userController) DemoteUser(c echo.Context) error {

	user := UserIDRequest{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err := uCtrl.userService.DemoteUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
