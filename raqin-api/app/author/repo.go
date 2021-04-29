package author

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AuthorRepo interface {
	NewAuthor(in *repo.Author) (*repo.Author, error)
	DeleteAuthor(author *repo.Author) (int64, error)
	UpdateAuthor(author *repo.Author) (*repo.Author, error)
	AllAuthors() (repo.AuthorSlice, error)
	AuthorByID(id int) (*repo.Author, error)
}

type authorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) *authorRepo {
	return &authorRepo{db}
}

func (aurepo *authorRepo) NewAuthor(author *repo.Author) (*repo.Author, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
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

func (aurepo *authorRepo) DeleteAuthor(author *repo.Author) (int64, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	n, err := author.Delete(ctx, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return n, nil

}

func (aurepo *authorRepo) UpdateAuthor(author *repo.Author) (*repo.Author, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	n, err := author.Update(ctx, tx, boil.Infer())
	if err != nil || n == 0 {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return author, nil

}

func (aurepo *authorRepo) AllAuthors() (repo.AuthorSlice, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	authors, err := repo.Authors().All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return authors, nil

}

func (aurepo *authorRepo) AuthorByID(id int) (*repo.Author, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	author, err := repo.FindAuthor(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return author, nil

}
