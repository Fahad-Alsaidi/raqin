package book

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"raqin-api/storage/repo"
	"time"

	"github.com/gen2brain/go-fitz"
	"github.com/volatiletech/null/v8"
)

type BookService interface {
	NewBook(in NewBookRequest) (NewBookResponse, error)
}

type bookService struct {
	bookRepo BookRepo
}

func NewBookService(bookRepo BookRepo) *bookService {
	return &bookService{bookRepo}
}

func (bkSrvc *bookService) NewBook(in NewBookRequest) (res NewBookResponse, err error) {

	defer in.File.Close()
	fileBytes, err := ioutil.ReadAll(in.File)
	if err != nil {
		return res, err
	}

	fileName := fmt.Sprintf("%d.pdf", time.Now().Unix())
	err = ioutil.WriteFile(fileName, fileBytes, 0755)
	if err != nil {
		return res, err
	}

	book := &repo.Book{
		Name: in.Name,
		Note: null.StringFrom(in.Notes),
		Path: fileName,
	}

	b, err := bkSrvc.bookRepo.NewBook(book)
	if err != nil {
		return res, err
	}

	res = NewBookResponse{
		Name:  b.Name,
		Notes: b.Note.String,
	}

	return res, nil

}

func BookPagesToImages(filename, destDir string) (int, error) {

	doc, err := fitz.New(filename)
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(destDir, fmt.Sprintf("page%d.jpg", n)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()
	}

	return doc.NumPage(), nil
}
