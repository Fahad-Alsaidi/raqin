package book

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type BookRepo interface {
	NewBook(*repo.Book, NewBookRequest) (*repo.Book, NewBookResponse, error)
}

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepo {
	return &bookRepo{db}
}

func (br *bookRepo) NewBook(book *repo.Book, in NewBookRequest) (*repo.Book, NewBookResponse, error) {
	res := NewBookResponse{}
	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, res, err
	}

	err = book.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, res, err
	}

	// Add Book Authors relationship to db
	for _, author := range in.Authors {
		bookAuthor := &repo.BookAuthor{
			BookID:   book.ID,
			AuthorID: author,
		}
		err := book.AddBookAuthors(ctx, tx, true, bookAuthor)
		if err != nil {
			tx.Rollback()
			return nil, res, err
		}
	}

	// Add Book Categories relationship to db
	for _, category := range in.Category {
		bookCategory := &repo.BookCategory{
			BookID:     book.ID,
			CategoryID: category,
		}
		err := book.AddBookCategories(ctx, tx, true, bookCategory)
		if err != nil {
			tx.Rollback()
			return nil, res, err
		}
	}

	// TODO: Add Book Initiators relationship to db

	tx.Commit()

	return book, res, nil

}
