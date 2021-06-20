package book

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"
	"raqin-api/utils/irror"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	errCanNotInsert   = irror.New("can not insert book")
	errCanNotDelete   = irror.New("can not delete book")
	errCanNotUpdate   = irror.New("can not update book")
	errCanNotGetBook  = irror.New("can not get book")
	errCanNotGetBooks = irror.New("can not get books")

	errCanNotAddBookAuthor       = irror.New("failed, can not add book author")
	errCanNotRemoveBookAuthor    = irror.New("failed, can not remove book author")
	errCanNotAddBookCategory     = irror.New("failed, can not add book category")
	errCanNotRemoveBookCategory  = irror.New("failed, can not remove book category")
	errCanNotAddBookInitiator    = irror.New("failed, can not add book initiator")
	errCanNotRemoveBookInitiator = irror.New("failed, can not remove book initiator")

	errCanNotFindBookAuthors    = irror.New("can not find book authors")
	errCanNotFindBookCategories = irror.New("can not find book categories")
	errCanNotFindBookInitiators = irror.New("can not find book initiators")

	errCanNotAddBookPages = irror.New("failed, can not add book pages")
)

type BookRepo interface {
	NewBook(*repo.Book, NewBookRequest) (*repo.Book, error)
	UpdateBook(book *repo.Book) (int64, error)
	DeleteBook(book *repo.Book) (int64, error)
	AllBooks() (repo.BookSlice, error)
	BookByID(id int) (*repo.Book, error)
	BookAuthors(id int) (*repo.AuthorSlice, error)
	BookCategories(id int) (*repo.CategorySlice, error)
	BookInitiators(id int) (*repo.UserSlice, error)
	AddBookPages(pages []string, BookID int) (int, error)
	AddBookAuthor(bkAuthor *repo.BookAuthor) error
	RemoveBookAuthor(bkAuthor *repo.BookAuthor) (int64, error)
	AddBookCategory(bkCategory *repo.BookCategory) error
	RemoveBookCategory(bkCategory *repo.BookCategory) (int64, error)
	AddBookInitiator(bkInitiater *repo.BookInitiater) error
	RemoveBookInitiator(bkInitiater *repo.BookInitiater) (int64, error)
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
		return nil, errCanNotInsert.Wrap(err)
	}

	err = book.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, errCanNotInsert.Wrap(err)
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
			return nil, errCanNotAddBookAuthor.Wrap(err)
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
			return nil, errCanNotAddBookCategory.Wrap(err)
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
			return nil, errCanNotAddBookInitiator.Wrap(err)
		}
	}

	tx.Commit()
	return book, nil
}

func (br *bookRepo) UpdateBook(book *repo.Book) (int64, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotUpdate.Wrap(err)
	}

	n, err := book.Update(ctx, tx, boil.Whitelist("name", "note", "updated_at"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotUpdate.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (br *bookRepo) AllBooks() (repo.BookSlice, error) {
	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetBooks.Wrap(err)
	}

	books, err := repo.Books(qm.Where("deleted_at = '0000-00-00 00:00:00'")).All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotGetBooks.Wrap(err)
	}

	tx.Commit()
	return books, nil
}

func (br *bookRepo) BookAuthors(id int) (*repo.AuthorSlice, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotFindBookAuthors.Wrap(err)
	}

	authors, err := repo.BookAuthors(
		qm.Where("book_id = ?", id),
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Load(repo.BookAuthorRels.Author)).
		All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotFindBookAuthors.Wrap(err)
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
		return nil, errCanNotFindBookCategories.Wrap(err)
	}

	bookCategories, err := repo.BookCategories(
		qm.Where("book_id = ?", id),
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Load(repo.BookCategoryRels.Category)).
		All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotFindBookCategories.Wrap(err)
	}

	cats := repo.CategorySlice{}
	for _, v := range bookCategories {
		cats = append(cats, v.R.Category)
	}

	return &cats, nil
}

func (br *bookRepo) BookInitiators(id int) (*repo.UserSlice, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotFindBookInitiators.Wrap(err)
	}

	bookInitiator, err := repo.BookInitiaters(
		qm.Where("book_id = ?", id),
		qm.Where("deleted_at = '0000-00-00 00:00:00'"),
		qm.Load(repo.BookInitiaterRels.User)).
		All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotFindBookInitiators.Wrap(err)
	}

	inits := repo.UserSlice{}
	for _, v := range bookInitiator {
		inits = append(inits, v.R.User)
	}

	return &inits, nil
}

func (br *bookRepo) BookByID(id int) (*repo.Book, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetBook.Wrap(err)
	}

	book, err := repo.FindBook(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotGetBook.Wrap(err)
	}

	return book, nil
}

func (br *bookRepo) DeleteBook(book *repo.Book) (int64, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotDelete.Wrap(err)
	}

	n, err := book.Update(ctx, tx, boil.Whitelist("deleted_at"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotDelete.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (br *bookRepo) AddBookAuthor(bkAuthor *repo.BookAuthor) error {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return errCanNotAddBookAuthor.Wrap(err)
	}

	err = bkAuthor.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return errCanNotAddBookAuthor.Wrap(err)
	}

	tx.Commit()
	return nil
}

func (br *bookRepo) RemoveBookAuthor(bkAuthor *repo.BookAuthor) (int64, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotRemoveBookAuthor.Wrap(err)
	}

	n, err := repo.BookAuthors(
		qm.Where("book_id = ?", bkAuthor.BookID),
		qm.Where("author_id = ?", bkAuthor.AuthorID),
	).DeleteAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotRemoveBookAuthor.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (br *bookRepo) AddBookCategory(bkCategory *repo.BookCategory) error {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return errCanNotAddBookCategory.Wrap(err)
	}

	err = bkCategory.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return errCanNotAddBookCategory.Wrap(err)
	}

	tx.Commit()
	return nil
}

func (br *bookRepo) RemoveBookCategory(bkCategory *repo.BookCategory) (int64, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotRemoveBookCategory.Wrap(err)
	}

	n, err := repo.BookCategories(
		qm.Where("book_id = ?", bkCategory.BookID),
		qm.Where("category_id = ?", bkCategory.CategoryID),
	).DeleteAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotRemoveBookCategory.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (br *bookRepo) AddBookInitiator(bkInitiater *repo.BookInitiater) error {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return errCanNotAddBookInitiator.Wrap(err)
	}

	err = bkInitiater.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return errCanNotAddBookInitiator.Wrap(err)
	}

	tx.Commit()
	return nil
}

func (br *bookRepo) RemoveBookInitiator(bkInitiater *repo.BookInitiater) (int64, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotRemoveBookInitiator.Wrap(err)
	}

	n, err := repo.BookInitiaters(
		qm.Where("book_id = ?", bkInitiater.BookID),
		qm.Where("user_id = ?", bkInitiater.UserID),
	).DeleteAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotRemoveBookInitiator.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (br *bookRepo) AddBookPages(pages []string, BookID int) (int, error) {

	ctx := context.Background()
	tx, err := br.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotAddBookPages.Wrap(err)
	}

	for i, v := range pages {
		bkPage := &repo.Page{
			BookID: BookID,
			Path:   v,
			Number: i + 1,
		}
		err = bkPage.Insert(ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return 0, errCanNotAddBookPages.Wrap(err)
		}
	}

	tx.Commit()
	return len(pages), nil
}
