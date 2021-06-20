package adminrouter

import (
	"raqin-api/app/page"
	"raqin-api/storage"

	"github.com/labstack/echo/v4"
)

func pageRouter(gRoute *echo.Group) {
	db := storage.DBInstance()
	pageRepo := page.NewPageRepo(db)
	pageSrvc := page.NewPageService(pageRepo)
	pageCtrl := page.NewPageController(pageSrvc)

	// page revision related
	gRoute.GET("/book-pages", pageCtrl.PagesByBookID)

	// revision related routes
	revision := gRoute.Group("/revision")
	revision.POST("", pageCtrl.NewPageRevision)
	revision.DELETE("", pageCtrl.DeletePageRevision)
	revision.PATCH("", pageCtrl.UpdatePageRevision)
	revision.GET("/by-page", pageCtrl.RevisionsByPageID)
	// here is the manual approval for raqin users
	revision.GET("/approve", pageCtrl.ApproveRevision)

	// reaction related routes
	reaction := revision.Group("/reaction")
	// auto approval for volunteers
	reaction.POST("", pageCtrl.NewReaction)
	reaction.PATCH("", pageCtrl.UpdateReaction)

	// comment related routes
	comment := revision.Group("/comment")
	comment.POST("", pageCtrl.NewComment)
	comment.DELETE("", pageCtrl.DeleteComment)
	comment.PATCH("", pageCtrl.UpdateComment)
	comment.GET("", pageCtrl.CommentsByRevisionID)
}
