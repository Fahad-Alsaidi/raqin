package book

import (
	"fmt"
	"io/ioutil"
	"raqin-api/app"
	"raqin-api/app/author"
	"raqin-api/app/category"
	"raqin-api/app/user"
	"raqin-api/storage/repo"
	"time"

	"github.com/volatiletech/null/v8"
)

type BookService interface {
	NewBook(in NewBookRequest) (NewBookResponse, error)
	UpdateBook(in UpdateBookRequest) (res NewBookResponse, err error)
	DeleteBook(in BookIDRequest) error
	AllBooks() ([]NewBookResponse, error)
	BookByID(in BookIDRequest) (res NewBookResponse, err error)
	BookToPages(in BookIDRequest) (res NewBookResponse, err error)
}

type bookService struct {
	bookRepo BookRepo
}

func NewBookService(bookRepo BookRepo) *bookService {
	return &bookService{bookRepo}
}

// NewBook will register book in db and file system and returns NewBookResponse
func (bkSrvc *bookService) NewBook(in NewBookRequest) (res NewBookResponse, err error) {

	defer in.File.Close()
	fileBytes, err := ioutil.ReadAll(in.File)
	if err != nil {
		return res, err
	}

	fileName := fmt.Sprintf("%d.pdf", time.Now().Unix())

	book := &repo.Book{
		Name: in.Name,
		Note: null.StringFrom(in.Notes),
		Path: fileName,
	}

	// save book in db
	b, err := bkSrvc.bookRepo.NewBook(book, in)
	if err != nil {
		return res, err
	}

	path := fmt.Sprintf("%d", b.ID)
	dir, err := app.CreateOrGetBookDir(path)
	if err != nil {
		return res, err
	}

	filePath := fmt.Sprintf("%s/%s", dir, fileName)

	// save the pdf if book is registered in db
	err = ioutil.WriteFile(filePath, fileBytes, 0700)
	if err != nil {
		return res, err
	}

	res, err = bkSrvc.BookRelations(b.ID)
	if err != nil {
		return res, err
	}

	res.Name = b.Name
	res.Notes = b.Note.String

	return res, nil
}

func (bkSrvc *bookService) UpdateBook(in UpdateBookRequest) (res NewBookResponse, err error) {

	book := &repo.Book{
		ID:   in.ID,
		Name: in.Name,
		Note: null.StringFrom(in.Notes),
	}

	bk, err := bkSrvc.bookRepo.UpdateBook(book)
	if err != nil {
		return res, err
	}

	res = NewBookResponse{
		Name:  bk.Name,
		Notes: bk.Note.String,
	}

	return res, nil
}

func (bkSrvc *bookService) DeleteBook(in BookIDRequest) error { panic("") }

func (bkSrvc *bookService) AllBooks() ([]NewBookResponse, error) { panic("") }

func (bkSrvc *bookService) BookByID(in BookIDRequest) (res NewBookResponse, err error) {

	bk, err := bkSrvc.bookRepo.BookByID(in.ID)
	if err != nil {
		return res, err
	}

	rels, err := bkSrvc.BookRelations(bk.ID)
	if err != nil {
		return res, err
	}

	res = NewBookResponse{
		Name:     bk.Name,
		Notes:    bk.Note.String,
		Authors:  rels.Authors,
		Category: rels.Category,
		Users:    rels.Users,
	}

	return res, nil
}

func (bkSrvc *bookService) BookRelations(id int) (res NewBookResponse, err error) {

	// get the authors of the book to return as response
	authors := []author.AuthorResponse{}
	auths, err := bkSrvc.bookRepo.BookAuthors(id)
	if err != nil {
		return res, err
	}
	for _, v := range *auths {
		a := author.AuthorResponse{
			ID:        int64(v.ID),
			FirstName: v.FirstName,
			LastName:  v.LastName,
		}
		authors = append(authors, a)
	}

	// get the categories of the book to return as response
	categories := []category.CategoryResponse{}
	cats, err := bkSrvc.bookRepo.BookCategories(id)
	if err != nil {
		return res, err
	}
	for _, v := range *cats {
		c := category.CategoryResponse{
			ID:   int64(v.ID),
			Name: v.Name,
		}
		categories = append(categories, c)
	}

	// get the initiators of the book to return as response
	users := []user.UserResponse{}
	inits, err := bkSrvc.bookRepo.BookInitiators(id)
	if err != nil {
		return res, err
	}
	for _, v := range *inits {
		u := user.UserResponse{
			ID:        int64(v.ID),
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			Gender:    v.Gender.String,
			Role:      v.Role,
		}
		users = append(users, u)
	}

	res = NewBookResponse{
		Authors:  authors,
		Category: categories,
		Users:    users,
	}

	return res, nil
}

func (bkSrvc *bookService) BookToPages(in BookIDRequest) (res NewBookResponse, err error) {

	bk, err := bkSrvc.bookRepo.BookByID(in.ID)
	if err != nil {
		return res, err
	}

	path := fmt.Sprintf("%d", bk.ID)
	dir, err := app.CreateOrGetBookDir(path)
	if err != nil {
		return res, err
	}

	pagesDir, err := app.CreateOrGetPagesDir(dir)
	if err != nil {
		return res, err
	}

	srcFile := fmt.Sprintf("%s/%s", dir, bk.Path)
	_, err = app.BookPagesToImages(srcFile, pagesDir)
	if err != nil {
		return res, err
	}

	return NewBookResponse{
		Name:  bk.Name,
		Notes: bk.Note.String,
	}, err
}
