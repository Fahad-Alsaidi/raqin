package page

import (
	"time"
)

type AddPageText struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type PageResponse struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Stage   string `json:"stage"`
	PageNum int    `json:"page_number"`
	BookID  int    `json:"book_id"`
}

type NewPageRevision struct {
	PageID int    `json:"id"`
	Text   string `json:"text"`
}

type UpdatePageRevision struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type ByID struct {
	ID int `query:"id"`
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
	RevisionID int    `json:"revision_id"`
	Reaction   string `json:"reaction"`
	PageID     int    `json:"page_id"`
}

type UpdateReaction struct {
	ID       int    `json:"id"`
	Reaction string `json:"reaction"`
}

type NewComment struct {
	RevisionID  int    `json:"revision_id"`
	CommenterID int    `json:"commentor_id"`
	Comment     string `json:"comment"`
}

type UpdateComment struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
}

type CommentResponse struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
}
