package category

import (
	"net/http"
	"raqin-api/utils/irror"

	"github.com/labstack/echo/v4"
)

type categoryController struct {
	categoryService CategoryService
}

func NewCategoryController(categoryService CategoryService) *categoryController {
	return &categoryController{categoryService: categoryService}
}

func (aCtrl *categoryController) NewCategory(c echo.Context) error {

	category := NewCategoryRequest{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, irror.New("can not bind data").Wrap(err))
	}

	resp, err := aCtrl.categoryService.NewCategory(category)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (aCtrl *categoryController) DeleteCategory(c echo.Context) error {

	category := DeleteCategoryRequest{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, irror.New("can not bind data").Wrap(err))
	}

	err := aCtrl.categoryService.DeleteCategory(category)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (aCtrl *categoryController) UpdateCategory(c echo.Context) error {

	category := UpdateCategoryRequest{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, irror.New("can not bind data").Wrap(err))
	}

	err := aCtrl.categoryService.UpdateCategory(category)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (aCtrl *categoryController) AllCategories(c echo.Context) error {

	categories, err := aCtrl.categoryService.AllCategories()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, categories)
}
