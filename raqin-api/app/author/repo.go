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

	err := author.Insert(context.Background(), br.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return author, nil

}
