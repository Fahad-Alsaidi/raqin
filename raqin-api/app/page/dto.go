package page

import (
	"time"
)

type PageResponse struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Stage   string `json:"stage"`
	PageNum int    `json:"page_number"`
	BookID  int    `json:"book_id"`
}

type NewPageRevision struct {
	PageID int    `json:"id" validate:"required,number"`
	Text   string `json:"text" validate:"required,ascii"`
}

type UpdatePageRevision struct {
	ID   int    `json:"id" validate:"required,number"`
	Text string `json:"text" validate:"required,ascii"`
}

type ByID struct {
	ID int `query:"id" validate:"required,number"`
}

type PageRevisionResponse struct {
	ID          int       `json:"id"`
	Text        string    `json:"text"`
	LastUpdated time.Time `json:"last_updated"`
	Reviewer    Reviewer  `json:"reviewer"`
	PageNum     int       `json:"page_number"`
	BookName    string    `json:"book_name"`
}

type Reviewer struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type NewReaction struct {
	RevisionID int    `json:"revision_id" validate:"required,number"`
	Reaction   string `json:"reaction" validate:"required,alpha"`
	PageID     int    `json:"page_id" validate:"required,number"`
}

type UpdateReaction struct {
	ID       int    `json:"id" validate:"required,number"`
	Reaction string `json:"reaction" validate:"required,alpha"`
}

type NewComment struct {
	RevisionID  int    `json:"revision_id" validate:"required,number"`
	CommenterID int    `json:"commentor_id" validate:"required,number"`
	Comment     string `json:"comment" validate:"required,alpha"`
}

type UpdateComment struct {
	ID      int    `json:"id" validate:"required,number"`
	Comment string `json:"comment" validate:"required,alpha"`
}

type CommentResponse struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
}
