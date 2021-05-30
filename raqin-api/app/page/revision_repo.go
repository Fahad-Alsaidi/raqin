package page

import (
	"context"
	"raqin-api/storage/repo"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (pr *pageRepo) NewRevision(pageRev *repo.PageRevision) (*repo.PageRevision, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = pageRev.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return pageRev, nil
}

func (pr *pageRepo) UpdateRevision(pageRev *repo.PageRevision) (int64, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	n, err := pageRev.Update(ctx, tx, boil.Whitelist("page_text"))
	if err != nil || n == 0 {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) DeleteRevision(pageRev *repo.PageRevision) (int64, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	n, err := pageRev.Update(ctx, tx, boil.Whitelist("deleted_at"))
	if err != nil || n == 0 {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) RevisionsByReviewerID(id int) (*repo.PageRevisionSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	revs, err := repo.PageRevisions(
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Where("reviewer_id = ?", id),
		qm.Load(repo.PageRevisionRels.Reviewer),
		qm.Load(repo.PageRevisionRels.Page)).All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &revs, nil
}

func (pr *pageRepo) RevisionsByPageID(id int) (*repo.PageRevisionSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	revs, err := repo.PageRevisions(
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Where("page_id = ?", id),
		qm.Load(repo.PageRevisionRels.Reviewer),
		qm.Load(repo.PageRevisionRels.Page)).All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &revs, nil
}

func (pr *pageRepo) ApproveRevision(id int, stage string) (*repo.Page, error) {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	rev, err := repo.FindPageRevision(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	pg := &repo.Page{
		ID:               rev.PageID,
		ApprovedRevision: null.IntFrom(id),
		UpdatedAt:        time.Now(),
		Stage:            stage,
	}

	_, err = pg.Update(ctx, tx, boil.Whitelist("approved_revision", "stage", "updated_at"))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = pg.Reload(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return pg, nil
}

// the follwing functions may not be used
func (pr *pageRepo) RevisionByID(id int) (*repo.PageRevision, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	pageRev, err := repo.PageRevisions(
		qm.Where("id = ?", id),
		qm.Load(repo.PageRevisionRels.Reviewer),
		qm.Load(repo.PageRevisionRels.Page)).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return pageRev, nil
}
func (pr *pageRepo) AllRevisions() (*repo.PageRevisionSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	revs, err := repo.PageRevisions(
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Load(repo.PageRevisionRels.Reviewer),
		qm.Load(repo.PageRevisionRels.Page)).All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &revs, nil
}
