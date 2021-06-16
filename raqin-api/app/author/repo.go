package author

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"

	"raqin-api/utils/irror"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	errCanNotInsert     = irror.New("can not insert author")
	errCanNotDelete     = irror.New("can not delete author")
	errCanNotUpdate     = irror.New("can not update author")
	errCanNotGetAuthor  = irror.New("can not get author")
	errCanNotGetAuthors = irror.New("can not get authors")
)

type AuthorRepo interface {
	NewAuthor(in *repo.Author) (*repo.Author, error)
	DeleteAuthor(author *repo.Author) (int64, error)
	UpdateAuthor(author *repo.Author) (int64, error)
	AllAuthors() (*repo.AuthorSlice, error)
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
		return nil, errCanNotInsert.Wrap(err)
	}

	err = author.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, errCanNotInsert.Wrap(err)
	}

	tx.Commit()

	return author, nil
}

func (aurepo *authorRepo) DeleteAuthor(author *repo.Author) (int64, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotDelete.Wrap(err)
	}

	n, err := author.Update(ctx, tx, boil.Whitelist("deleted_at"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotDelete.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (aurepo *authorRepo) UpdateAuthor(author *repo.Author) (int64, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotUpdate.Wrap(err)
	}

	n, err := author.Update(ctx, tx, boil.Whitelist("first_name", "last_name", "updated_at"))
	if err != nil {
		tx.Rollback()
		return n, errCanNotUpdate.Wrap(err)
	}

	tx.Commit()
	return 0, nil
}

func (aurepo *authorRepo) AllAuthors() (*repo.AuthorSlice, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetAuthors.Wrap(err)
	}

	authors, err := repo.Authors().All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotGetAuthors.Wrap(err)
	}

	tx.Commit()
	return &authors, nil
}

func (aurepo *authorRepo) AuthorByID(id int) (*repo.Author, error) {

	ctx := context.Background()
	tx, err := aurepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetAuthor.Wrap(err)
	}

	author, err := repo.FindAuthor(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotGetAuthor.Wrap(err)
	}

	tx.Commit()

	return author, nil
}
