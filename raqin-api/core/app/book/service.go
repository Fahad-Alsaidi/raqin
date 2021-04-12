package book

import (
	"fmt"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gen2brain/go-fitz"
)

// book interface the will hold all book related operations
type IBook interface {
	UploadNewBook(in NewBookRequest) // registers new book
}

// book struct that will implement IBook interface
type bookService struct{}

// NewBookController return bookController struct with IBook service
// this method is the entry to this controller
func NewBookService() bookService {
	return bookService{}
}

func (bs bookService) UploadNewBook(in NewBookRequest) {

	// Destination
	dst, err := os.Create(time.Now().String())
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	defer in.File.Close()

	// Copy
	if _, err = io.Copy(dst, in.File); err != nil {
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
