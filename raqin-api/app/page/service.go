package page

import (
	"fmt"
	"raqin-api/storage/repo"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
)

type PageService interface {
	NewRevision(NewPageRevision) error
	UpdateRevision(UpdatePageRevision) error
	DeleteRevision(ByID) error
	RevisionsByPageID(ByID) (*[]PageRevisionResponse, error)

	ApproveRevision(ByID) (*PageResponse, error)
	NewReaction(NewReaction) error
	UpdateReaction(UpdateReaction) error

	// PAGE_REVISION_COMMENTS
	NewComment(NewComment) error
	UpdateComment(UpdateComment) error
	DeleteComment(ByID) error
	CommentsByRevisionID(ByID) (*[]CommentResponse, error)
}

type pageService struct {
	pageRepo PageRepo
}

func NewPageService(pageRepo PageRepo) *pageService {
	return &pageService{pageRepo}
}

func (bgSrvc *pageService) NewRevision(pageRev NewPageRevision) error {

	pg := &repo.PageRevision{
		PageText:   null.StringFrom(pageRev.Text),
		PageID:     pageRev.PageID,
		ReviewerID: 1,
	}

	_, err := bgSrvc.pageRepo.NewRevision(pg)
	if err != nil {
		return err
	}

	return nil
}

func (bgSrvc *pageService) UpdateRevision(pageRev UpdatePageRevision) error {

	pg := &repo.PageRevision{
		PageText: null.StringFrom(pageRev.Text),
		ID:       pageRev.ID,
	}

	page, err := bgSrvc.pageRepo.PageByRevisionID(pageRev.ID)
	if err != nil {
		return err
	}

	if page.Stage == "INIT" || page.Stage == "NONE" {
		n, err := bgSrvc.pageRepo.UpdateRevision(pg)
		if err != nil {
			return err
		}
		if n == 0 {
			errors.New("nothing changed")
		}
	}

	// return errors.New("cannot change an approved revision")
	return nil
}

func (bgSrvc *pageService) DeleteRevision(pageRev ByID) error {

	pg := &repo.PageRevision{
		ID:        pageRev.ID,
		DeletedAt: time.Now(),
	}

	page, err := bgSrvc.pageRepo.PageByRevisionID(pageRev.ID)
	if err != nil {
		return err
	}

	if page.Stage == "INIT" || page.Stage == "NONE" {
		n, err := bgSrvc.pageRepo.DeleteRevision(pg)
		if err != nil {
			return err
		}
		if n == 0 {
			errors.New("nothing deleted")
		}
	}

	// return errors.New("cannot delete an approved revision")
	return nil
}

func (bgSrvc *pageService) RevisionsByPageID(pageRev ByID) (*[]PageRevisionResponse, error) {
	revs, err := bgSrvc.pageRepo.RevisionsByPageID(pageRev.ID)
	if err != nil {
		return nil, err
	}

	revResponse := []PageRevisionResponse{}
	for _, revision := range *revs {
		rev := PageRevisionResponse{
			ID:          revision.ID,
			Text:        revision.PageText.String,
			LastUpdated: revision.UpdatedAt,
			PageNum:     revision.R.Page.Number,
			// BookName:    revision.R.Page.R.Book.Name,
			Reviewer: Reviewer{
				Name:   revision.R.Reviewer.FirstName + " " + revision.R.Reviewer.LastName,
				Gender: revision.R.Reviewer.Gender.String,
			},
		}
		revResponse = append(revResponse, rev)
	}

	return &revResponse, nil
}

func (bgSrvc *pageService) ApproveRevision(pageRev ByID) (*PageResponse, error) {
	b, err := bgSrvc.pageRepo.ApproveRevision(pageRev.ID, "REV")
	if err != nil {
		return nil, err
	}

	return &PageResponse{
		Stage:   b.Stage,
		ID:      b.ID,
		PageNum: b.Number,
		BookID:  b.BookID,
	}, nil
}

func (bgSrvc *pageService) NewReaction(reaction NewReaction) error {

	re := &repo.PageRevisionReaction{
		Reaction:       reaction.Reaction,
		PageRevisionID: reaction.RevisionID,
		ReactorID:      1,
	}

	err := bgSrvc.pageRepo.NewReaction(re)
	if err != nil {
		return err
	}

	page, err := bgSrvc.pageRepo.PageByID(reaction.PageID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if page.Stage == "NONE" || page.Stage == "INIT" {
		rev, err := bgSrvc.pageRepo.MostApprovedRevisionByPageID(reaction.PageID)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		_, err = bgSrvc.pageRepo.ApproveRevision(rev.ID, "INIT")
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return nil
}

func (bgSrvc *pageService) UpdateReaction(reaction UpdateReaction) error {

	re := &repo.PageRevisionReaction{
		ID:       reaction.ID,
		Reaction: reaction.Reaction,
	}

	n, err := bgSrvc.pageRepo.UpdateReaction(re)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("nothing changed")
	}

	return nil
}

func (bgSrvc *pageService) NewComment(in NewComment) error {

	comment := &repo.PageRevisionComment{
		PageRevisionID: in.RevisionID,
		CommenterID:    in.CommenterID,
		Comment:        in.Comment,
	}

	err := bgSrvc.pageRepo.NewComment(comment)
	if err != nil {
		return err
	}

	return nil
}

func (bgSrvc *pageService) UpdateComment(in UpdateComment) error {

	comment := &repo.PageRevisionComment{
		ID:      in.ID,
		Comment: in.Comment,
	}

	n, err := bgSrvc.pageRepo.UpdateComment(comment)
	if err != nil {
		return err
	}
	if n == 0 {
		fmt.Println(n, "none")
		return errors.New("nothing changed")
	}

	return nil
}

func (bgSrvc *pageService) DeleteComment(in ByID) error {

	comment := &repo.PageRevisionComment{
		ID:        in.ID,
		DeletedAt: time.Now(),
	}

	n, err := bgSrvc.pageRepo.DeleteComment(comment)
	if err != nil {
		return err
	}
	if n == 0 {
		fmt.Println(n, "none")
		return errors.New("nothing changed")
	}

	return nil
}

func (bgSrvc *pageService) CommentsByRevisionID(in ByID) (*[]CommentResponse, error) {

	comments, err := bgSrvc.pageRepo.CommentsByRevisionID(in.ID)
	if err != nil {
		return nil, err
	}

	commentResponse := []CommentResponse{}
	for _, comment := range *comments {
		co := CommentResponse{
			ID:      comment.ID,
			Comment: comment.Comment,
		}
		commentResponse = append(commentResponse, co)
	}

	return &commentResponse, nil
}
