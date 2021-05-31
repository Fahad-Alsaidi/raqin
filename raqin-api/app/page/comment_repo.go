package page

import (
	"context"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (pr *pageRepo) NewComment(comment *repo.PageRevisionComment) error {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = comment.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (pr *pageRepo) UpdateComment(comment *repo.PageRevisionComment) (int64, error) {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	n, err := comment.Update(ctx, tx, boil.Whitelist("comment", "updated_at"))
	if err != nil {
		tx.Rollback()
		return n, err
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) DeleteComment(comment *repo.PageRevisionComment) (int64, error) {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	n, err := comment.Update(ctx, tx, boil.Whitelist("deleted_at"))
	if err != nil {
		tx.Rollback()
		return n, err
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) CommentsByRevisionID(id int) (*repo.PageRevisionCommentSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	comments, err := repo.PageRevisionComments(
		qm.Where("page_revision_id = ?", id),
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Load(repo.PageRevisionCommentRels.PageRevision),
		qm.Load(repo.PageRevisionCommentRels.Commenter)).All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &comments, nil
}
