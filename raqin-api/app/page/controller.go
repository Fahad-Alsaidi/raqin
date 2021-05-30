package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type pageController struct {
	pageservice PageService
}

func NewPageController(pageservice PageService) *pageController {
	return &pageController{pageservice: pageservice}
}

func (pCtrl *pageController) NewPageRevision(c echo.Context) error {

	pgRev := NewPageRevision{}
	if err := c.Bind(&pgRev); err != nil {
		return err
	}

	err := pCtrl.pageservice.NewRevision(pgRev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (pCtrl *pageController) UpdatePageRevision(c echo.Context) error {

	pgRev := UpdatePageRevision{}
	if err := c.Bind(&pgRev); err != nil {
		return err
	}

	err := pCtrl.pageservice.UpdateRevision(pgRev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, err)
}

func (pCtrl *pageController) DeletePageRevision(c echo.Context) error {

	pgRev := ByID{}
	if err := c.Bind(&pgRev); err != nil {
		return err
	}

	err := pCtrl.pageservice.DeleteRevision(pgRev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (pCtrl *pageController) RevisionsByPageID(c echo.Context) error {

	pgRev := ByID{}
	if err := c.Bind(&pgRev); err != nil {
		panic(err)
	}

	resp, err := pCtrl.pageservice.RevisionsByPageID(pgRev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (pCtrl *pageController) ApproveRevision(c echo.Context) error {

	pgRev := ByID{}
	if err := c.Bind(&pgRev); err != nil {
		panic(err)
	}

	resp, err := pCtrl.pageservice.ApproveRevision(pgRev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (pCtrl *pageController) NewReaction(c echo.Context) error {

	revReac := NewReaction{}
	if err := c.Bind(&revReac); err != nil {
		panic(err)
	}

	err := pCtrl.pageservice.NewReaction(revReac)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (pCtrl *pageController) UpdateReaction(c echo.Context) error {

	revReac := UpdateReaction{}
	if err := c.Bind(&revReac); err != nil {
		panic(err)
	}

	err := pCtrl.pageservice.UpdateReaction(revReac)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (pCtrl *pageController) NewComment(c echo.Context) error {

	comment := NewComment{}
	if err := c.Bind(&comment); err != nil {
		return err
	}

	err := pCtrl.pageservice.NewComment(comment)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (pCtrl *pageController) UpdateComment(c echo.Context) error {

	comment := UpdateComment{}
	if err := c.Bind(&comment); err != nil {
		return err
	}

	err := pCtrl.pageservice.UpdateComment(comment)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (pCtrl *pageController) DeleteComment(c echo.Context) error {

	comment := ByID{}
	if err := c.Bind(&comment); err != nil {
		return err
	}

	err := pCtrl.pageservice.DeleteComment(comment)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (pCtrl *pageController) CommentsByRevisionID(c echo.Context) error {

	comment := ByID{}
	if err := c.Bind(&comment); err != nil {
		panic(err)
	}

	resp, err := pCtrl.pageservice.CommentsByRevisionID(comment)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
