package page

import (
	"context"
	"raqin-api/storage/repo"
	"raqin-api/utils/irror"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	errCanNotInsertReaction = irror.New("can not insert reaction")
	errCanNotUpdateReaction = irror.New("can not update reaction")
	errCanNotGetReactions   = irror.New("can not get reactions")

	errCanNotGetMostApprovedRevision = irror.New("can not get most approved revision")
)

func (pr *pageRepo) NewReaction(pReact *repo.PageRevisionReaction) error {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return errCanNotInsertReaction.Wrap(err)
	}

	err = pReact.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return errCanNotInsertReaction.Wrap(err)
	}

	tx.Commit()
	return nil
}

func (pr *pageRepo) UpdateReaction(pReact *repo.PageRevisionReaction) (int64, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotUpdateReaction.Wrap(err)
	}

	n, err := pReact.Update(ctx, tx, boil.Whitelist("reaction", "updated_at"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotUpdateReaction.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (pr *pageRepo) ReactionsByRevisionID(id int) (*repo.PageRevisionReactionSlice, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetReactions.Wrap(err)
	}

	reactions, err := repo.PageRevisionReactions(
		qm.Where("page_revision_id = ?", id),
		qm.Load(repo.PageRevisionReactionRels.PageRevision),
		qm.Load(repo.PageRevisionReactionRels.Reactor)).All(ctx, tx)
	if err != nil {
		return nil, errCanNotGetReactions.Wrap(err)
	}

	return &reactions, nil
}

func (pr *pageRepo) MostApprovedRevisionByPageID(id int) (*repo.PageRevision, error) {

	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetMostApprovedRevision.Wrap(err)
	}

	revs, err := repo.PageRevisions(
		qm.Where("page_id = ?", id)).All(ctx, tx)
	if err != nil {
		return nil, errCanNotGetMostApprovedRevision.Wrap(err)
	}

	var approved int = 0
	var mostCount int64 = 0
	if len(revs) > 1 {
		for i, v := range revs {
			n, err := repo.PageRevisionReactions(
				qm.Where("page_revision_id = ?", v.ID),
				qm.Where("reaction = ?", repo.PageRevisionReactionReactionAPPROVE)).
				Count(ctx, tx)
			if err != nil {
				return nil, errCanNotGetMostApprovedRevision.Wrap(err)
			}

			if n > mostCount {
				approved = i
			}
			mostCount = n
		}
	}

	return revs[approved], nil
}
