package page

import (
	"context"
	"raqin-api/storage/repo"
	"raqin-api/utils/irror"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	errCanNotInsert          = irror.New("can not insert revision")
	errCanNotDelete          = irror.New("can not delete revision")
	errCanNotUpdate          = irror.New("can not update revision")
	errCanNotGetRevision     = irror.New("can not get revision")
	errCanNotGetRevisions    = irror.New("can not get revisions")
	errCanNotApproveRevision = irror.New("can not approve revision")
)

func (pr *pageRepo) NewRevision(pageRev *repo.PageRevision) (*repo.PageRevision, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotInsert.Wrap(err)
	}

	err = pageRev.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, errCanNotInsert.Wrap(err)
	}

	tx.Commit()
	return pageRev, nil
}

func (pr *pageRepo) UpdateRevision(pageRev *repo.PageRevision) (int64, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotUpdate.Wrap(err)
	}

	n, err := pageRev.Update(ctx, tx, boil.Whitelist("page_text"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotUpdate.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) DeleteRevision(pageRev *repo.PageRevision) (int64, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotDelete.Wrap(err)
	}

	n, err := pageRev.Update(ctx, tx, boil.Whitelist("deleted_at"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotDelete.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) RevisionsByReviewerID(id int) (*repo.PageRevisionSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetRevisions.Wrap(err)
	}

	revs, err := repo.PageRevisions(
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Where("reviewer_id = ?", id),
		qm.Load(repo.PageRevisionRels.Reviewer),
		qm.Load(repo.PageRevisionRels.Page)).All(ctx, tx)
	if err != nil {
		return nil, errCanNotGetRevisions.Wrap(err)
	}

	return &revs, nil
}

func (pr *pageRepo) RevisionsByPageID(id int) (*repo.PageRevisionSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetRevisions.Wrap(err)
	}

	revs, err := repo.PageRevisions(
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Where("page_id = ?", id),
		qm.Load(repo.PageRevisionRels.Reviewer),
		qm.Load(repo.PageRevisionRels.Page)).All(ctx, tx)
	if err != nil {
		return nil, errCanNotGetRevisions.Wrap(err)
	}

	return &revs, nil
}

func (pr *pageRepo) ApproveRevision(id int, stage string) (*repo.Page, error) {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotApproveRevision.Wrap(err)
	}

	rev, err := repo.FindPageRevision(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotApproveRevision.Wrap(err)
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
		return nil, errCanNotApproveRevision.Wrap(err)
	}

	err = pg.Reload(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotApproveRevision.Wrap(err)
	}

	tx.Commit()
	return pg, nil
}

// the follwing functions may not be used
func (pr *pageRepo) RevisionByID(id int) (*repo.PageRevision, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetRevision.Wrap(err)
	}

	pageRev, err := repo.PageRevisions(
		qm.Where("id = ?", id),
		qm.Load(repo.PageRevisionRels.Reviewer),
		qm.Load(repo.PageRevisionRels.Page)).One(ctx, tx)
	if err != nil {
		return nil, errCanNotGetRevision.Wrap(err)
	}

	return pageRev, nil
}

func (pr *pageRepo) AllRevisions() (*repo.PageRevisionSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetRevisions.Wrap(err)
	}

	revs, err := repo.PageRevisions(
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Load(repo.PageRevisionRels.Reviewer),
		qm.Load(repo.PageRevisionRels.Page)).All(ctx, tx)
	if err != nil {
		return nil, errCanNotGetRevisions.Wrap(err)
	}

	return &revs, nil
}
