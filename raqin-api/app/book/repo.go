package book

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type BookRepo interface {
	NewBook(in *repo.Book) (*repo.Book, error)
}

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepo {
	return &bookRepo{db}
}

func (br *bookRepo) NewBook(book *repo.Book) (*repo.Book, error) {

	err := book.Insert(context.Background(), br.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return book, nil

}
