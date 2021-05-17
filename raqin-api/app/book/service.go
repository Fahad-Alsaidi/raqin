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
	NewBook(NewBookRequest) (BookResponse, error)
	UpdateBook(UpdateBookRequest) (BookResponse, error)
	DeleteBook(BookIDRequest) error
	AllBooks() ([]BookResponse, error)
	BookByID(BookIDRequest) (BookResponse, error)
	BookToPages(BookIDRequest) (int, error)
	AddBookAuthor(bkAuthor AddBookRel) error
	RemoveBookAuthor(bkAuthor RemoveBookRel) error
	AddBookCategory(bkCategory AddBookRel) error
	RemoveBookCategory(bkCategory RemoveBookRel) error
	AddBookInitiator(bkInitiater AddBookRel) error
	RemoveBookInitiator(bkInitiater RemoveBookRel) error
}

type bookService struct {
	bookRepo BookRepo
}

func NewBookService(bookRepo BookRepo) *bookService {
	return &bookService{bookRepo}
}

// NewBook will register book in db and file system and returns BookResponse
func (bkSrvc *bookService) NewBook(in NewBookRequest) (res BookResponse, err error) {

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

	res.ID = b.ID
	res.Name = b.Name
	res.Notes = b.Note.String

	return res, nil
}

func (bkSrvc *bookService) UpdateBook(in UpdateBookRequest) (res BookResponse, err error) {

	book := &repo.Book{
		ID:        in.ID,
		Name:      in.Name,
		UpdatedAt: time.Now(),
		Note:      null.StringFrom(in.Notes),
	}

	bk, err := bkSrvc.bookRepo.UpdateBook(book)
	if err != nil {
		return res, err
	}

	res, err = bkSrvc.BookRelations(bk.ID)
	if err != nil {
		return res, err
	}

	res.ID = bk.ID
	res.Name = bk.Name
	res.Notes = bk.Note.String

	return res, nil
}

func (bkSrvc *bookService) DeleteBook(in BookIDRequest) error {

	book := &repo.Book{
		ID:        in.ID,
		DeletedAt: time.Now(),
	}

	n, err := bkSrvc.bookRepo.DeleteBook(book)
	if err != nil || n == 0 {
		return err
	}

	return nil
}

func (bkSrvc *bookService) AllBooks() ([]BookResponse, error) {
	books, err := bkSrvc.bookRepo.AllBooks()
	if err != nil {
		return nil, err
	}

	bookResponse := []BookResponse{}
	for _, bk := range books {
		res, err := bkSrvc.BookRelations(bk.ID)
		if err != nil {
			return nil, err
		}
		res.ID = bk.ID
		res.Name = bk.Name
		res.Notes = bk.Note.String

		bookResponse = append(bookResponse, res)
	}

	return bookResponse, nil
}

func (bkSrvc *bookService) BookByID(in BookIDRequest) (res BookResponse, err error) {

	bk, err := bkSrvc.bookRepo.BookByID(in.ID)
	if err != nil {
		return res, err
	}

	rels, err := bkSrvc.BookRelations(bk.ID)
	if err != nil {
		return res, err
	}

	res = BookResponse{
		ID:       bk.ID,
		Name:     bk.Name,
		Notes:    bk.Note.String,
		Authors:  rels.Authors,
		Category: rels.Category,
		Users:    rels.Users,
	}

	return res, nil
}

func (bkSrvc *bookService) BookRelations(id int) (res BookResponse, err error) {

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

	res = BookResponse{
		Authors:  authors,
		Category: categories,
		Users:    users,
	}

	return res, nil
}

func (bkSrvc *bookService) BookToPages(in BookIDRequest) (res int, err error) {

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
	n, pages, err := app.BookPagesToImages(srcFile, pagesDir)
	if err != nil || n != len(pages) {
		return res, err
	}

	n, err = bkSrvc.bookRepo.AddBookPages(pages, bk.ID)
	if err != nil || n != len(pages) {
		return res, err
	}

	return n, err
}

func (bkSrvc *bookService) AddBookAuthor(bkAuthor AddBookRel) error {

	ba := &repo.BookAuthor{
		BookID:   bkAuthor.BookID,
		AuthorID: bkAuthor.ID,
	}

	err := bkSrvc.bookRepo.AddBookAuthor(ba)
	if err != nil {
		return err
	}

	return nil
}

func (bkSrvc *bookService) RemoveBookAuthor(bkAuthor RemoveBookRel) error {
	ba := &repo.BookAuthor{
		ID: bkAuthor.ID,
	}

	_, err := bkSrvc.bookRepo.RemoveBookAuthor(ba)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

func (bkSrvc *bookService) AddBookCategory(bkCategory AddBookRel) error {

	bc := &repo.BookCategory{
		BookID:     bkCategory.BookID,
		CategoryID: bkCategory.ID,
	}

	err := bkSrvc.bookRepo.AddBookCategory(bc)
	if err != nil {
		return err
	}

	return nil
}

func (bkSrvc *bookService) RemoveBookCategory(bkCategory RemoveBookRel) error {

	bc := &repo.BookCategory{
		ID: bkCategory.ID,
	}

	_, err := bkSrvc.bookRepo.RemoveBookCategory(bc)
	if err != nil {
		return err
	}

	return nil
}

func (bkSrvc *bookService) AddBookInitiator(bkInitiater AddBookRel) error {

	bi := &repo.BookInitiater{
		BookID: bkInitiater.BookID,
		UserID: bkInitiater.ID,
	}

	err := bkSrvc.bookRepo.AddBookInitiator(bi)
	if err != nil {
		return err
	}

	return nil
}

func (bkSrvc *bookService) RemoveBookInitiator(bkInitiater RemoveBookRel) error {

	bi := &repo.BookInitiater{
		ID: bkInitiater.ID,
	}

	_, err := bkSrvc.bookRepo.RemoveBookInitiator(bi)
	if err != nil {
		return err
	}

	return nil
}
