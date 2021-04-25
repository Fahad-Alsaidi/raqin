package author

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AuthorRepo interface {
	NewAuthor(in *repo.Author) (*repo.Author, error)
}

type authorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) *authorRepo {
	return &authorRepo{db}
}

func (br *authorRepo) NewAuthor(author *repo.Author) (*repo.Author, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = author.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return author, nil

}
