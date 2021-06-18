package page

import (
	"context"
	"raqin-api/storage/repo"
	"raqin-api/utils/irror"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	errCanNotInsertComment = irror.New("can not insert comment")
	errCanNotDeleteComment = irror.New("can not delete comment")
	errCanNotUpdateComment = irror.New("can not update comment")
	errCanNotGetComments   = irror.New("can not get comments")
)

func (pr *pageRepo) NewComment(comment *repo.PageRevisionComment) error {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return errCanNotInsertComment.Wrap(err)
	}

	err = comment.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return errCanNotInsertComment.Wrap(err)
	}

	tx.Commit()
	return nil
}

func (pr *pageRepo) UpdateComment(comment *repo.PageRevisionComment) (int64, error) {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotUpdateComment.Wrap(err)
	}

	n, err := comment.Update(ctx, tx, boil.Whitelist("comment", "updated_at"))
	if err != nil {
		tx.Rollback()
		return n, errCanNotUpdateComment.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) DeleteComment(comment *repo.PageRevisionComment) (int64, error) {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotDeleteComment.Wrap(err)
	}

	n, err := comment.Update(ctx, tx, boil.Whitelist("deleted_at"))
	if err != nil {
		tx.Rollback()
		return n, errCanNotDeleteComment.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) CommentsByRevisionID(id int) (*repo.PageRevisionCommentSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetComments.Wrap(err)
	}

	comments, err := repo.PageRevisionComments(
		qm.Where("page_revision_id = ?", id),
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Load(repo.PageRevisionCommentRels.PageRevision),
		qm.Load(repo.PageRevisionCommentRels.Commenter)).All(ctx, tx)
	if err != nil {
		return nil, errCanNotGetComments.Wrap(err)
	}

	return &comments, nil
}
