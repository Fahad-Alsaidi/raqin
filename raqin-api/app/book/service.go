package book

import (
	"context"
	"database/sql"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"raqin-api/storage/repo"
	"time"

	"github.com/gen2brain/go-fitz"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type BookService interface {
	NewBookProject(in NewBookRequest)
}

type bookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *bookService {
	return &bookService{db}
}

func (bs *bookService) NewBookProject(in NewBookRequest) {

	defer in.File.Close()
	fileBytes, err := ioutil.ReadAll(in.File)
	if err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%d.pdf", time.Now().Unix())
	err = ioutil.WriteFile(fileName, fileBytes, 0755)
	if err != nil {
		panic(err)
	}

	book := &repo.Book{
		Name: "some book",
		Note: null.StringFrom("i dont know"),
		Path: fileName,
	}

	book.Insert(context.Background(), bs.db, boil.Infer())
	if err != nil {
		panic(err)
	}

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
