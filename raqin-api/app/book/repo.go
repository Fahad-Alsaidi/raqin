package book

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type BookRepo interface {
	NewBook(*repo.Book, NewBookRequest) (*repo.Book, error)
	UpdateBook(book *repo.Book) (*repo.Book, error)
	DeleteBook(book *repo.Book) (int64, error)
	AllBooks() (repo.BookSlice, error)
	BookByID(id int) (*repo.Book, error)
	BookAuthors(id int) (*repo.AuthorSlice, error)
	BookCategories(id int) (*repo.CategorySlice, error)
	BookInitiators(id int) (*repo.UserSlice, error)
}

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepo {
	return &bookRepo{db}
}

func (br *bookRepo) NewBook(book *repo.Book, in NewBookRequest) (*repo.Book, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = book.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
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
			return nil, err
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
			return nil, err
		}
	}

	// Add Book Initiators relationship to db
	for _, initiator := range in.Initiators {
		bookInitiator := &repo.BookInitiater{
			BookID: book.ID,
			UserID: initiator,
		}
		err := book.AddBookInitiaters(ctx, tx, true, bookInitiator)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()

	return book, nil
}

func (br *bookRepo) UpdateBook(book *repo.Book) (*repo.Book, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	n, err := book.Update(ctx, tx, boil.Whitelist("name", "note"))
	if err != nil || n == 0 {
		tx.Rollback()
		return nil, err
	}

	return book, nil
}

func (br *bookRepo) AllBooks() (repo.BookSlice, error) { panic("") }

func (br *bookRepo) BookAuthors(id int) (*repo.AuthorSlice, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	authors, err := repo.BookAuthors(
		qm.Where("book_id = ?", id),
		qm.Load(repo.BookAuthorRels.Author)).
		All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	auths := repo.AuthorSlice{}
	for _, v := range authors {
		auths = append(auths, v.R.Author)
	}

	return &auths, nil
}

func (br *bookRepo) BookCategories(id int) (*repo.CategorySlice, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	bookCategories, err := repo.BookCategories(
		qm.Where("book_id = ?", id),
		qm.Load(repo.BookCategoryRels.Category)).
		All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	cats := repo.CategorySlice{}
	for _, v := range bookCategories {
		cats = append(cats, v.R.Category)
	}

	tx.Commit()

	return &cats, nil
}

func (br *bookRepo) BookInitiators(id int) (*repo.UserSlice, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	bookInitiator, err := repo.BookInitiaters(
		qm.Where("book_id = ?", id),
		qm.Load(repo.BookInitiaterRels.User)).
		All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	inits := repo.UserSlice{}
	for _, v := range bookInitiator {
		inits = append(inits, v.R.User)
	}

	tx.Commit()

	return &inits, nil
}

func (br *bookRepo) BookByID(id int) (*repo.Book, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	book, err := repo.FindBook(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return book, nil
}

func (br *bookRepo) DeleteBook(book *repo.Book) (int64, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	n, err := book.Delete(ctx, tx)
	if err != nil || n == 0 {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return n, nil
}

func (br *bookRepo) AddBookAuthor(bkAuthor *repo.BookAuthor) error {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = bkAuthor.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (br *bookRepo) RemoveBookAuthor(bkAuthor *repo.BookAuthor) (int64, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	n, err := bkAuthor.Delete(ctx, tx)
	if err != nil || n == 0 {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return n, nil
}

func (br *bookRepo) AddBookCategory(bkCategory *repo.BookCategory) error {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = bkCategory.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (br *bookRepo) RemoveBookCategory(bkCategory *repo.BookCategory) (int64, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	n, err := bkCategory.Delete(ctx, tx)
	if err != nil || n == 0 {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return n, nil
}
